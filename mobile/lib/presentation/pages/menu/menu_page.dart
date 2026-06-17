import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:bukadulu/data/datasources/api.dart';

class MenuPage extends ConsumerStatefulWidget {
  final String ventureId;
  const MenuPage({super.key, required this.ventureId});
  @override
  ConsumerState<MenuPage> createState() => _MenuPageState();
}

class _MenuPageState extends ConsumerState<MenuPage> {
  final _nameCtl = TextEditingController();
  final _descCtl = TextEditingController();
  List<Map<String, dynamic>> _menus = [];
  bool _loading = true;

  @override
  void initState() {
    super.initState();
    _load();
  }

  @override
  void dispose() {
    _nameCtl.dispose();
    _descCtl.dispose();
    super.dispose();
  }

  Future<void> _load() async {
    setState(() => _loading = true);
    try {
      final api = ref.read(authApiProvider);
      final data = await api.listMenus(widget.ventureId);
      setState(() => _menus = List<Map<String, dynamic>>.from(data['data']));
    } catch (_) {}
    setState(() => _loading = false);
  }

  Future<void> _add() async {
    if (_nameCtl.text.trim().isEmpty) return;
    try {
      final api = ref.read(authApiProvider);
      await api.createMenu(widget.ventureId, name: _nameCtl.text.trim(), description: _descCtl.text.trim());
      _nameCtl.clear();
      _descCtl.clear();
      _load();
    } catch (e) {
      if (mounted) ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text('$e')));
    }
  }

  Future<void> _toggleActive(String id, bool active) async {
    try {
      final api = ref.read(authApiProvider);
      await api.updateMenu(widget.ventureId, id, status: active ? 'active' : 'candidate');
      _load();
    } catch (e) {
      if (mounted) ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text('$e')));
    }
  }

  Future<void> _delete(String id) async {
    try {
      final api = ref.read(authApiProvider);
      await api.deleteMenu(widget.ventureId, id);
      _load();
    } catch (_) {}
  }

  Future<void> _focus() async {
    try {
      final api = ref.read(authApiProvider);
      await api.focusMenus(widget.ventureId);
      if (mounted) context.go('/venture/${widget.ventureId}/cost');
    } catch (e) {
      if (mounted) ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text('$e')));
    }
  }

  @override
  Widget build(BuildContext context) {
    final activeCount = _menus.where((m) => m['status'] == 'active').length;

    return Scaffold(
      appBar: AppBar(title: const Text('Pilih Menu')),
      body: _loading
          ? const Center(child: CircularProgressIndicator())
          : Column(
              children: [
                Padding(
                  padding: const EdgeInsets.all(16),
                  child: Row(
                    children: [
                      Expanded(
                        child: TextFormField(
                          controller: _nameCtl,
                          decoration: const InputDecoration(labelText: 'Nama Menu', hintText: 'Nasi Goreng', isDense: true),
                        ),
                      ),
                      const SizedBox(width: 8),
                      Expanded(
                        child: TextFormField(
                          controller: _descCtl,
                          decoration: const InputDecoration(labelText: 'Deskripsi', isDense: true),
                        ),
                      ),
                      const SizedBox(width: 8),
                      IconButton.filled(onPressed: _add, icon: const Icon(Icons.add)),
                    ],
                  ),
                ),
                Padding(
                  padding: const EdgeInsets.symmetric(horizontal: 16),
                  child: Text('Aktif: $activeCount/3 SKU', style: Theme.of(context).textTheme.bodySmall),
                ),
                Expanded(
                  child: _menus.isEmpty
                      ? const Center(child: Text('Belum ada menu'))
                      : ListView.builder(
                          padding: const EdgeInsets.symmetric(horizontal: 16),
                          itemCount: _menus.length,
                          itemBuilder: (_, i) {
                            final m = _menus[i];
                            final isActive = m['status'] == 'active';
                            final isHero = m['is_hero'] == true;
                            return Card(
                              margin: const EdgeInsets.only(bottom: 8),
                              child: ListTile(
                                leading: isHero ? const Icon(Icons.star, color: Color(0xFFf59e0b)) : null,
                                title: Text(m['name'] ?? ''),
                                subtitle: Text('${m['status']} ${isHero ? '⭐ Hero' : ''}'),
                                trailing: Row(
                                  mainAxisSize: MainAxisSize.min,
                                  children: [
                                    if (!isActive && activeCount < 3)
                                      IconButton(
                                        icon: const Icon(Icons.check_circle_outline, color: Color(0xFF22c55e)),
                                        onPressed: () => _toggleActive(m['id'], true),
                                      ),
                                    if (isActive)
                                      IconButton(
                                        icon: const Icon(Icons.undo, color: Color(0xFFea580c)),
                                        onPressed: () => _toggleActive(m['id'], false),
                                      ),
                                    IconButton(
                                      icon: const Icon(Icons.delete_outline, color: Color(0xFFef4444)),
                                      onPressed: () => _delete(m['id']),
                                    ),
                                  ],
                                ),
                              ),
                            );
                          },
                        ),
                ),
                Padding(
                  padding: const EdgeInsets.all(16),
                  child: SizedBox(
                    width: double.infinity, height: 48,
                    child: ElevatedButton(
                      onPressed: activeCount >= 1 && activeCount <= 3 ? _focus : null,
                      child: const Text('Kunci Pilihan Menu'),
                    ),
                  ),
                ),
              ],
            ),
    );
  }
}
