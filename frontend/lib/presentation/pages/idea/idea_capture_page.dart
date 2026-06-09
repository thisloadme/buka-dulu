import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:bukadulu/data/datasources/api.dart';
class IdeaCapturePage extends ConsumerStatefulWidget {
  final String ventureId;
  const IdeaCapturePage({super.key, required this.ventureId});
  @override
  ConsumerState<IdeaCapturePage> createState() => _IdeaCapturePageState();
}

class _IdeaCapturePageState extends ConsumerState<IdeaCapturePage> {
  final _controller = TextEditingController();
  bool _loading = false;

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }

  Future<void> _submit() async {
    final text = _controller.text.trim();
    if (text.length < 20) {
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('Minimal 20 karakter')),
      );
      return;
    }
    setState(() => _loading = true);
    try {
      final api = ref.read(authApiProvider);
      await api.captureIdea(widget.ventureId, text);
      if (mounted) context.go('/venture/${widget.ventureId}/idea/result');
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
      appBar: AppBar(title: const Text('Ceritakan Idemu')),
      body: SingleChildScrollView(
        padding: const EdgeInsets.all(24),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(
              'Tulis ide bisnismu dengan detail',
              style: Theme.of(context).textTheme.titleMedium?.copyWith(fontWeight: FontWeight.w600),
            ),
            const SizedBox(height: 8),
            Text(
              'Sebutkan produk, target pelanggan, harga, dan lokasi. Semakin detail semakin baik.',
              style: TextStyle(color: Colors.grey[600]),
            ),
            const SizedBox(height: 24),
            TextFormField(
              controller: _controller,
              maxLines: 8,
              decoration: InputDecoration(
                hintText: 'Contoh: Saya ingin jual nasi goreng homemade dengan topping ayam geprek...',
                alignLabelWithHint: true,
                border: OutlineInputBorder(borderRadius: BorderRadius.circular(12)),
              ),
            ),
            const SizedBox(height: 8),
            Text('Minimal 20 karakter', style: TextStyle(fontSize: 12, color: Colors.grey[500])),
            const SizedBox(height: 24),
            SizedBox(
              width: double.infinity,
              height: 48,
              child: ElevatedButton(
                onPressed: _loading ? null : _submit,
                child: _loading
                    ? const SizedBox(width: 20, height: 20, child: CircularProgressIndicator(strokeWidth: 2))
                    : const Text('Proses Ide'),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
