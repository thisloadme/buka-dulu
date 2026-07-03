import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:bukadulu/data/datasources/api.dart';

const _kCategories = <String>[
  'makanan_berat',
  'minuman',
  'snack',
  'frozen_food',
  'catering',
  'kue_dan_roti',
];

class VentureCreatePage extends ConsumerStatefulWidget {
  const VentureCreatePage({super.key});
  @override
  ConsumerState<VentureCreatePage> createState() => _VentureCreatePageState();
}

class _VentureCreatePageState extends ConsumerState<VentureCreatePage> {
  final _formKey = GlobalKey<FormState>();
  final _nameController = TextEditingController();
  final _regionController = TextEditingController();
  String? _selectedCategory;
  bool _loading = false;

  @override
  void dispose() {
    _nameController.dispose();
    _regionController.dispose();
    super.dispose();
  }

  Future<void> _submit() async {
    if (!_formKey.currentState!.validate()) return;
    setState(() => _loading = true);
    try {
      final api = ref.read(authApiProvider);
      final data = await api.createVenture(
        name: _nameController.text.trim(),
        category: _selectedCategory ?? '',
        region: _regionController.text.trim(),
      );
      if (mounted) context.go('/venture/${data['id']}/idea');
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text('Gagal: $e')));
      }
    } finally {
      if (mounted) setState(() => _loading = false);
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Ide Baru')),
      body: SingleChildScrollView(
        padding: const EdgeInsets.all(24),
        child: Form(
          key: _formKey,
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text('Ceritakan ide bisnismu', style: Theme.of(context).textTheme.titleMedium?.copyWith(fontWeight: FontWeight.w600)),
              const SizedBox(height: 24),
              TextFormField(
                controller: _nameController,
                decoration: const InputDecoration(labelText: 'Nama Usaha *', hintText: 'Contoh: Warung Nasi Ayam Geprek'),
                validator: (v) => v == null || v.trim().isEmpty ? 'Wajib diisi' : null,
              ),
              const SizedBox(height: 16),
              DropdownButtonFormField<String>(
                value: _selectedCategory,
                items: _kCategories
                    .map((c) => DropdownMenuItem<String>(
                          value: c,
                          child: Text(c.replaceAll('_', ' ')),
                        ))
                    .toList(),
                onChanged: (v) => setState(() => _selectedCategory = v),
                decoration: const InputDecoration(labelText: 'Kategori'),
              ),
              const SizedBox(height: 16),
              TextFormField(
                controller: _regionController,
                decoration: const InputDecoration(labelText: 'Lokasi', hintText: 'Contoh: Jakarta Pusat'),
              ),
              const SizedBox(height: 32),
              SizedBox(
                width: double.infinity,
                height: 48,
                child: ElevatedButton(
                  onPressed: _loading ? null : _submit,
                  child: _loading
                      ? const SizedBox(width: 20, height: 20, child: CircularProgressIndicator(strokeWidth: 2))
                      : const Text('Lanjut ke Ide'),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
