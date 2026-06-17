import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:bukadulu/data/datasources/api.dart';

class CustomerPage extends ConsumerStatefulWidget {
  final String ventureId;
  const CustomerPage({super.key, required this.ventureId});
  @override
  ConsumerState<CustomerPage> createState() => _CustomerPageState();
}

class _CustomerPageState extends ConsumerState<CustomerPage> {
  final _nameCtl = TextEditingController();
  final _ageCtl = TextEditingController();
  final _contextCtl = TextEditingController();
  final _budgetCtl = TextEditingController();
  final _locCtl = TextEditingController();
  final _momentCtl = TextEditingController();
  bool _loading = false;

  @override
  void dispose() {
    _nameCtl.dispose();
    _ageCtl.dispose();
    _contextCtl.dispose();
    _budgetCtl.dispose();
    _locCtl.dispose();
    _momentCtl.dispose();
    super.dispose();
  }

  Future<void> _submit() async {
    setState(() => _loading = true);
    try {
      final api = ref.read(authApiProvider);
      await api.createCustomer(widget.ventureId,
        name: _nameCtl.text.trim(),
        ageRange: _ageCtl.text.trim(),
        buyContext: _contextCtl.text.trim(),
        budgetRange: _budgetCtl.text.trim(),
        location: _locCtl.text.trim(),
        consumptionMoment: _momentCtl.text.trim(),
      );
      await api.confirmCustomer(widget.ventureId);
      if (mounted) context.go('/venture/${widget.ventureId}/menu');
    } catch (e) {
      if (mounted) ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text('$e')));
    } finally {
      if (mounted) setState(() => _loading = false);
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Target Pelanggan')),
      body: SingleChildScrollView(
        padding: const EdgeInsets.all(24),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text('Siapa target pelangganmu?', style: Theme.of(context).textTheme.titleMedium?.copyWith(fontWeight: FontWeight.w600)),
            const SizedBox(height: 24),
            TextFormField(controller: _nameCtl, decoration: const InputDecoration(labelText: 'Nama Segmen *', hintText: 'Karyawan kantoran 25-40th')),
            const SizedBox(height: 16),
            TextFormField(controller: _ageCtl, decoration: const InputDecoration(labelText: 'Rentang Usia', hintText: '25-40')),
            const SizedBox(height: 16),
            TextFormField(controller: _contextCtl, decoration: const InputDecoration(labelText: 'Konteks Beli', hintText: 'Makan siang di kantor')),
            const SizedBox(height: 16),
            TextFormField(controller: _budgetCtl, decoration: const InputDecoration(labelText: 'Budget', hintText: 'Rp 15.000 - 20.000')),
            const SizedBox(height: 16),
            TextFormField(controller: _locCtl, decoration: const InputDecoration(labelText: 'Lokasi', hintText: 'Area perkantoran Jakarta Pusat')),
            const SizedBox(height: 16),
            TextFormField(controller: _momentCtl, decoration: const InputDecoration(labelText: 'Momen Konsumsi', hintText: 'Senin-Jumat jam 11.30-13.00')),
            const SizedBox(height: 32),
            SizedBox(
              width: double.infinity, height: 48,
              child: ElevatedButton(
                onPressed: _loading ? null : _submit,
                child: _loading
                    ? const SizedBox(width: 20, height: 20, child: CircularProgressIndicator(strokeWidth: 2))
                    : const Text('Konfirmasi & Lanjut'),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
