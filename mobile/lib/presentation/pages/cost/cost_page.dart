import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:bukadulu/data/datasources/api.dart';
import 'package:bukadulu/presentation/widgets/common/brand_icons.dart';

class CostPage extends ConsumerStatefulWidget {
  final String ventureId;
  const CostPage({super.key, required this.ventureId});
  @override
  ConsumerState<CostPage> createState() => _CostPageState();
}

class _CostPageState extends ConsumerState<CostPage> {
  final _nameCtl = TextEditingController();
  final _unitCtl = TextEditingController();
  final _qtyCtl = TextEditingController();
  final _priceCtl = TextEditingController();
  final _laborCtl = TextEditingController();
  final _overheadCtl = TextEditingController();

  List<Map<String, dynamic>> _menus = [];
  String? _selectedMenuId;
  List<Map<String, dynamic>> _ingredients = [];
  Map<String, dynamic>? _summary;
  bool _loading = true;

  @override
  void initState() {
    super.initState();
    _load();
  }

  @override
  void dispose() {
    _nameCtl.dispose();
    _unitCtl.dispose();
    _qtyCtl.dispose();
    _priceCtl.dispose();
    _laborCtl.dispose();
    _overheadCtl.dispose();
    super.dispose();
  }

  Future<void> _load() async {
    setState(() => _loading = true);
    try {
      final api = ref.read(authApiProvider);
      final menuData = await api.listMenus(widget.ventureId);
      setState(() => _menus = List<Map<String, dynamic>>.from(menuData['data']).where((m) => m['status'] == 'active').toList());
      if (_menus.isNotEmpty && _selectedMenuId == null) {
        _selectedMenuId = _menus.first['id'];
      }
      if (_selectedMenuId != null) {
        await _loadIngredients();
      }
    } catch (_) {}
    setState(() => _loading = false);
  }

  Future<void> _loadIngredients() async {
    if (_selectedMenuId == null) return;
    try {
      final api = ref.read(authApiProvider);
      final ingData = await api.listIngredients(widget.ventureId, _selectedMenuId!);
      setState(() => _ingredients = List<Map<String, dynamic>>.from(ingData['data']));
    } catch (_) {
      setState(() => _ingredients = []);
    }
  }

  Future<void> _addIngredient() async {
    if (_selectedMenuId == null || _nameCtl.text.isEmpty) return;
    try {
      final api = ref.read(authApiProvider);
      await api.addIngredient(widget.ventureId,
        menuId: _selectedMenuId!,
        name: _nameCtl.text.trim(),
        unit: _unitCtl.text.trim(),
        quantity: double.parse(_qtyCtl.text),
        unitPrice: double.parse(_priceCtl.text),
      );
      _nameCtl.clear(); _unitCtl.clear(); _qtyCtl.clear(); _priceCtl.clear();
      await _loadIngredients();
    } catch (e) {
      if (mounted) ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text('$e')));
    }
  }

  Future<void> _calculate() async {
    if (_selectedMenuId == null) return;
    setState(() => _summary = null);
    try {
      final api = ref.read(authApiProvider);
      final labor = double.tryParse(_laborCtl.text) ?? 0;
      final overhead = double.tryParse(_overheadCtl.text) ?? 0;
      final s = await api.calculateCost(widget.ventureId, _selectedMenuId!, laborPerUnit: labor, overheadPerUnit: overhead);
      setState(() => _summary = s);
    } catch (e) {
      if (mounted) ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text('$e')));
    }
  }

  Future<void> _confirm() async {
    try {
      final api = ref.read(authApiProvider);
      await api.confirmCost(widget.ventureId);
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(const SnackBar(content: Text('Biaya dikunci!')));
        context.go('/dashboard');
      }
    } catch (e) {
      if (mounted) ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text('$e')));
    }
  }

  Color _marginColor(String status) {
    switch (status) {
      case 'sehat': return Colors.green;
      case 'tipis': return Colors.orange;
      default: return Colors.red;
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Biaya & Margin')),
      body: _loading
          ? const Center(child: CircularProgressIndicator())
          : SingleChildScrollView(
              padding: const EdgeInsets.all(16),
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  // Menu selector
                  DropdownButtonFormField<String>(
                    value: _selectedMenuId,
                    items: _menus.map<DropdownMenuItem<String>>((m) => DropdownMenuItem<String>(value: m['id'] as String, child: Text(m['name'] as String))).toList(),
                    onChanged: (v) async {
                      setState(() => _selectedMenuId = v);
                      await _loadIngredients();
                    },
                    decoration: const InputDecoration(labelText: 'Pilih Menu'),
                  ),
                  const SizedBox(height: 16),

                  // Ingredient input
                  const Text('Tambah Bahan', style: TextStyle(fontWeight: FontWeight.w600)),
                  const SizedBox(height: 8),
                  Row(children: [
                    Expanded(child: TextFormField(controller: _nameCtl, decoration: const InputDecoration(labelText: 'Bahan', isDense: true))),
                    const SizedBox(width: 4),
                    SizedBox(width: 60, child: TextFormField(controller: _unitCtl, decoration: const InputDecoration(labelText: 'Unit', isDense: true))),
                    const SizedBox(width: 4),
                    SizedBox(width: 70, child: TextFormField(controller: _qtyCtl, decoration: const InputDecoration(labelText: 'Qty', isDense: true), keyboardType: TextInputType.number)),
                    const SizedBox(width: 4),
                    SizedBox(width: 80, child: TextFormField(controller: _priceCtl, decoration: const InputDecoration(labelText: 'Harga', isDense: true), keyboardType: TextInputType.number)),
                    const SizedBox(width: 4),
                    IconButton.filled(onPressed: _addIngredient, icon: const Icon(Icons.add, size: 20)),
                  ]),
                  const SizedBox(height: 8),

                  // Ingredient list
                  ..._ingredients.map((ing) => ListTile(
                    dense: true,
                    title: Text('${ing['name']} — ${ing['quantity']} ${ing['unit']} @ ${ing['unit_price']}'),
                    trailing: Text('${(ing['quantity'] * ing['unit_price']).toStringAsFixed(0)}', style: const TextStyle(fontWeight: FontWeight.w600)),
                  )),
                  const Divider(),

                  // Labor & overhead
                  Row(children: [
                    Expanded(child: TextFormField(controller: _laborCtl, decoration: const InputDecoration(labelText: 'Tenaga/porsi', isDense: true), keyboardType: TextInputType.number)),
                    const SizedBox(width: 12),
                    Expanded(child: TextFormField(controller: _overheadCtl, decoration: const InputDecoration(labelText: 'Overhead/porsi', isDense: true), keyboardType: TextInputType.number)),
                  ]),
                  const SizedBox(height: 16),
                  SizedBox(width: double.infinity, child: ElevatedButton(onPressed: _calculate, child: const Text('Hitung HPP & Margin'))),

                  if (_summary != null) ...[
                    const Divider(height: 32),
                    Card(
                      child: Padding(
                        padding: const EdgeInsets.all(16),
                        child: Column(
                          children: [
                            _summaryRow('HPP per Porsi', 'Rp ${(_summary!['hpp_per_porsi'] as num).toStringAsFixed(0)}'),
                            _summaryRow('Harga Jual Saran', 'Rp ${(_summary!['suggested_price'] as num).toStringAsFixed(0)}'),
                            _summaryRow('Margin Kotor', '${(_summary!['gross_margin'] as num).toStringAsFixed(1)}%'),
                            Row(
                              mainAxisAlignment: MainAxisAlignment.spaceBetween,
                              children: [
                                const Text('Status Margin'),
                                Container(
                                  padding: const EdgeInsets.symmetric(horizontal: 12, vertical: 4),
                                  decoration: BoxDecoration(
                                    color: _marginColor(_summary!['margin_status']).withValues(alpha: 0.1),
                                    borderRadius: BorderRadius.circular(12),
                                  ),
                                  child: Row(
                                    mainAxisSize: MainAxisSize.min,
                                    children: [
                                      _summary!['margin_status'] == 'sehat'
                                          ? BrandIcons.checkCircle(size: 16, color: const Color(0xFF22c55e))
                                          : _summary!['margin_status'] == 'tipis'
                                              ? BrandIcons.alertTriangle(size: 16, color: const Color(0xFFf59e0b))
                                              : BrandIcons.xCircle(size: 16, color: const Color(0xFFef4444)),
                                      const SizedBox(width: 4),
                                      Text(
                                        _summary!['margin_status'] == 'sehat' ? 'Sehat' : _summary!['margin_status'] == 'tipis' ? 'Tipis' : 'Berbahaya',
                                        style: TextStyle(color: _marginColor(_summary!['margin_status']), fontWeight: FontWeight.w600),
                                      ),
                                    ],
                                  ),
                                ),
                              ],
                            ),
                            _summaryRow('Break-even', '${_summary!['break_even_unit']} unit/bulan'),
                          ],
                        ),
                      ),
                    ),
                    const SizedBox(height: 16),
                    SizedBox(
                      width: double.infinity, height: 48,
                      child: ElevatedButton(
                        onPressed: _confirm,
                        child: const Text('Konfirmasi & Selesai'),
                      ),
                    ),
                  ],
                ],
              ),
            ),
    );
  }

  Widget _summaryRow(String label, String value) {
    return Padding(
      padding: const EdgeInsets.symmetric(vertical: 4),
      child: Row(mainAxisAlignment: MainAxisAlignment.spaceBetween, children: [
        Text(label, style: const TextStyle(color: Color(0xFF57534e))),
        Text(value, style: const TextStyle(fontWeight: FontWeight.w600)),
      ]),
    );
  }
}
