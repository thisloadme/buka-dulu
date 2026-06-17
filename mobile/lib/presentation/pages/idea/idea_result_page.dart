import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:bukadulu/data/datasources/api.dart';
import 'package:bukadulu/presentation/providers/idea_provider.dart';

class IdeaResultPage extends ConsumerStatefulWidget {
  final String ventureId;
  const IdeaResultPage({super.key, required this.ventureId});
  @override
  ConsumerState<IdeaResultPage> createState() => _IdeaResultPageState();
}

class _IdeaResultPageState extends ConsumerState<IdeaResultPage> {
  bool _processing = false;
  bool _confirming = false;

  Future<void> _process() async {
    setState(() => _processing = true);
    try {
      final api = ref.read(authApiProvider);
      await api.processIdea(widget.ventureId);
      ref.invalidate(ideaProvider(widget.ventureId));
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text('Gagal proses: $e')));
      }
    } finally {
      if (mounted) setState(() => _processing = false);
    }
  }

  Future<void> _confirm() async {
    setState(() => _confirming = true);
    try {
      final api = ref.read(authApiProvider);
      await api.confirmIdea(widget.ventureId);
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('Ide berhasil dikunci!')),
        );
        context.go('/dashboard');
      }
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text('Gagal: $e')));
      }
    } finally {
      if (mounted) setState(() => _confirming = false);
    }
  }

  @override
  Widget build(BuildContext context) {
    final ideaAsync = ref.watch(ideaProvider(widget.ventureId));
    final theme = Theme.of(context);

    return Scaffold(
      appBar: AppBar(title: const Text('Konsep Ide')),
      body: ideaAsync.when(
        loading: () => const Center(child: CircularProgressIndicator()),
        error: (err, _) => Center(
          child: Column(
            mainAxisSize: MainAxisSize.min,
            children: [
              const Icon(Icons.error_outline, size: 48, color: Colors.red),
              const SizedBox(height: 16),
              Text('Gagal memuat ide', style: theme.textTheme.titleMedium),
              const SizedBox(height: 8),
              ElevatedButton(onPressed: () => ref.invalidate(ideaProvider(widget.ventureId)), child: const Text('Coba Lagi')),
            ],
          ),
        ),
        data: (idea) {
          if (idea == null) {
            return const Center(child: Text('Belum ada ide'));
          }

          if (idea.status == 'pending' || idea.status == 'processing') {
            return Center(
              child: Column(
                mainAxisSize: MainAxisSize.min,
                children: [
                  const CircularProgressIndicator(),
                  const SizedBox(height: 16),
                  Text('AI sedang menyusun konsep...', style: theme.textTheme.titleMedium),
                  const SizedBox(height: 24),
                  ElevatedButton(
                    onPressed: _processing ? null : _process,
                    child: _processing ? const Text('Memproses...') : const Text('Proses Sekarang'),
                  ),
                ],
              ),
            );
          }

          return SingleChildScrollView(
            padding: const EdgeInsets.all(16),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                if (idea.status == 'done') ...[
                  const Row(
                    children: [
                      Icon(Icons.check_circle, color: Colors.green),
                      SizedBox(width: 8),
                      Text('Ide berhasil diproses!', style: TextStyle(color: Colors.green, fontWeight: FontWeight.w600)),
                    ],
                  ),
                  const SizedBox(height: 24),
                  _ConceptCard(title: 'Konsep 1 Kalimat', content: idea.oneLineConcept ?? '-'),
                  _ConceptCard(title: 'Target Customer', content: idea.targetCustomer ?? '-'),
                  _ConceptCard(title: 'Value Proposition', content: idea.valueProposition ?? '-'),
                  if (idea.keyAssumptions != null)
                    _ConceptCard(title: 'Key Assumptions', content: _parseList(idea.keyAssumptions!)),
                  if (idea.earlyRisks != null)
                    _ConceptCard(title: 'Early Risks', content: _parseList(idea.earlyRisks!)),
                  const SizedBox(height: 24),
                  SizedBox(
                    width: double.infinity,
                    height: 48,
                    child: ElevatedButton(
                      onPressed: _confirming ? null : _confirm,
                      child: _confirming
                          ? const SizedBox(width: 20, height: 20, child: CircularProgressIndicator(strokeWidth: 2))
                          : const Text('Konfirmasi Ide'),
                    ),
                  ),
                  if (idea.status == 'failed') ...[
                    const Row(
                      children: [
                        Icon(Icons.error, color: Colors.red),
                        SizedBox(width: 8),
                        Text('Gagal memproses ide', style: TextStyle(color: Colors.red)),
                      ],
                    ),
                    const SizedBox(height: 16),
                    ElevatedButton(
                      onPressed: _processing ? null : _process,
                      child: const Text('Coba Lagi'),
                    ),
                  ],
                ],
              ],
            ),
          );
        },
      ),
    );
  }

  String _parseList(String jsonStr) {
    try {
      final list = jsonDecode(jsonStr) as List;
      return list.map((e) => '• $e').join('\n');
    } catch (_) {
      return jsonStr;
  }
}
}

class _ConceptCard extends StatelessWidget {
  final String title;
  final String content;
  const _ConceptCard({required this.title, required this.content});

  @override
  Widget build(BuildContext context) {
    return Card(
      margin: const EdgeInsets.only(bottom: 12),
      child: Padding(
        padding: const EdgeInsets.all(16),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(title, style: TextStyle(color: Color(0xFF57534e), fontSize: 12, fontWeight: FontWeight.w600)),
            const SizedBox(height: 8),
            Text(content, style: Theme.of(context).textTheme.bodyMedium),
          ],
        ),
      ),
    );
  }
}
