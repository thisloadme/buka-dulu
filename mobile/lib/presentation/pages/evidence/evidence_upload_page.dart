import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:bukadulu/data/datasources/api.dart';
import 'package:bukadulu/presentation/widgets/common/brand_icons.dart';

class EvidenceUploadPage extends ConsumerStatefulWidget {
  final String ventureId;
  final String missionId;
  const EvidenceUploadPage({super.key, required this.ventureId, required this.missionId});
  @override
  ConsumerState<EvidenceUploadPage> createState() => _EvidenceUploadPageState();
}

class _EvidenceUploadPageState extends ConsumerState<EvidenceUploadPage> {
  final _textCtl = TextEditingController();
  final _linkCtl = TextEditingController();
  String _type = 'text';
  Map<String, dynamic>? _review;
  bool _loading = false;

  @override
  void dispose() {
    _textCtl.dispose();
    _linkCtl.dispose();
    super.dispose();
  }

  Future<void> _submit() async {
    if (_type == 'text' && _textCtl.text.trim().isEmpty) return;
    if (_type == 'link' && _linkCtl.text.trim().isEmpty) return;
    setState(() => _loading = true);
    try {
      final api = ref.read(authApiProvider);
      final ev = await api.uploadEvidence(widget.ventureId,
        missionId: widget.missionId,
        evidenceType: _type,
        textContent: _type == 'text' ? _textCtl.text.trim() : _linkCtl.text.trim(),
      );
      // Auto-review
      final r = await api.reviewEvidence(widget.ventureId, ev['id']);
      setState(() => _review = r);
    } catch (e) {
      if (mounted) ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text('$e')));
    }
    setState(() => _loading = false);
  }

  Color _verdictColor(String v) {
    switch (v) {
      case 'valid': return Colors.green;
      case 'weak': return Colors.orange;
      case 'invalid': return Colors.red;
      default: return Colors.grey;
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Upload Evidence')),
      body: SingleChildScrollView(
        padding: const EdgeInsets.all(24),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            const Text('Pilih tipe bukti:', style: TextStyle(fontWeight: FontWeight.w600)),
            const SizedBox(height: 12),
            SegmentedButton<String>(
              segments: const [
                ButtonSegment(value: 'text', label: Text('Catatan'), icon: Icon(Icons.text_fields)),
                ButtonSegment(value: 'link', label: Text('Link'), icon: Icon(Icons.link)),
                ButtonSegment(value: 'image', label: Text('Foto'), icon: Icon(Icons.image)),
              ],
              selected: {_type},
              onSelectionChanged: (v) => setState(() => _type = v.first),
            ),
            const SizedBox(height: 24),
            if (_type == 'text')
              TextFormField(
                controller: _textCtl,
                maxLines: 6,
                decoration: InputDecoration(
                  labelText: 'Tulis buktimu',
                  hintText: 'Contoh: Saya sudah tanya 10 orang...',
                  border: OutlineInputBorder(borderRadius: BorderRadius.circular(12)),
                ),
              ),
            if (_type == 'link')
              TextFormField(
                controller: _linkCtl,
                decoration: InputDecoration(
                  labelText: 'Link URL',
                  hintText: 'https://...',
                  border: OutlineInputBorder(borderRadius: BorderRadius.circular(12)),
                ),
              ),
            if (_type == 'image')
              Container(
                height: 120,
                decoration: BoxDecoration(
                  border: Border.all(color: Color(0xFFe7e5e4)),
                  borderRadius: BorderRadius.circular(12),
                ),
                child: Center(
                  child: Column(
                    mainAxisSize: MainAxisSize.min,
                    children: [
                      BrandIcons.camera(size: 40, color: Color(0xFFa8a29e)),
                      const SizedBox(height: 8),
                      const Text('Ambil foto dari galeri atau kamera', style: TextStyle(color: Color(0xFF57534e))),
                    ],
                  ),
                ),
              ),
            const SizedBox(height: 24),
            SizedBox(
              width: double.infinity, height: 48,
              child: ElevatedButton(
                onPressed: _loading ? null : _submit,
                child: _loading
                    ? const SizedBox(width: 20, height: 20, child: CircularProgressIndicator(strokeWidth: 2))
                    : const Text('Kirim & Review'),
              ),
            ),
            if (_review != null) ...[
              const SizedBox(height: 24),
              const Divider(),
              const SizedBox(height: 16),
              Row(children: [
                _review!['verdict'] == 'valid'
                    ? BrandIcons.checkCircle(color: _verdictColor(_review!['verdict']))
                    : _review!['verdict'] == 'weak'
                        ? BrandIcons.alertTriangle(color: _verdictColor(_review!['verdict']))
                        : BrandIcons.xCircle(color: _verdictColor(_review!['verdict'])),
                const SizedBox(width: 8),
                const Text('Verdict: ', style: TextStyle(color: Color(0xFF57534e))),
                Text(_review!['verdict'] ?? '', style: TextStyle(
                  fontWeight: FontWeight.w600,
                  color: _verdictColor(_review!['verdict']),
                )),
              ]),
              const SizedBox(height: 8),
              Text(_review!['rationale'] ?? ''),
              if (_review!['next_action'] != null) ...[
                const SizedBox(height: 8),
                Text('Next: ${_review!['next_action']}', style: TextStyle(
                  color: _review!['next_action'] == 'continue' ? Colors.green : Colors.orange,
                  fontWeight: FontWeight.w600,
                )),
              ],
              const SizedBox(height: 24),
              SizedBox(
                width: double.infinity,
                child: OutlinedButton(
                  onPressed: () => context.go('/venture/${widget.ventureId}/missions'),
                  child: const Text('Kembali ke Misi'),
                ),
              ),
            ],
          ],
        ),
      ),
    );
  }
}
