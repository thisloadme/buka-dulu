import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:bukadulu/data/datasources/api.dart';

class ScorePage extends ConsumerStatefulWidget {
  final String ventureId;
  const ScorePage({super.key, required this.ventureId});
  @override
  ConsumerState<ScorePage> createState() => _ScorePageState();
}

class _ScorePageState extends ConsumerState<ScorePage> {
  Map<String, dynamic>? _score;
  Map<String, dynamic>? _decision;
  bool _loading = true;

  @override
  void initState() {
    super.initState();
    _load();
  }

  Future<void> _load() async {
    setState(() => _loading = true);
    try {
      final api = ref.read(authApiProvider);
      final s = await api.calculateScore(widget.ventureId);
      setState(() => _score = s);
    } catch (_) {}
    setState(() => _loading = false);
  }

  Future<void> generateDecision() async {
    try {
      final api = ref.read(authApiProvider);
      final d = await api.generateDecision(widget.ventureId);
      setState(() {
        _decision = d['decision'];
        _score = d['score'];
      });
    } catch (e) {
      if (mounted) ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text('$e')));
    }
  }

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);

    return Scaffold(
      appBar: AppBar(title: const Text('Skor Kesiapan')),
      body: _loading
          ? const Center(child: CircularProgressIndicator())
          : SingleChildScrollView(
              padding: const EdgeInsets.all(24),
              child: Column(children: [
                if (_score != null) ...[
                  // Big score circle
                  Container(
                    width: 160, height: 160,
                    decoration: BoxDecoration(
                      shape: BoxShape.circle,
                      color: _scoreColor(_score!['total_score']).withValues(alpha: 0.1),
                      border: Border.all(color: _scoreColor(_score!['total_score']), width: 4),
                    ),
                    child: Center(
                      child: Column(
                        mainAxisSize: MainAxisSize.min,
                        children: [
                          Text('${(_score!['total_score'] as num).toInt()}', style: TextStyle(fontSize: 48, fontWeight: FontWeight.bold, color: _scoreColor(_score!['total_score']))),
                          Text('/100', style: TextStyle(color: Colors.grey[500])),
                        ],
                      ),
                    ),
                  ),
                  const SizedBox(height: 32),

                  // Score breakdown
                  _scoreBar('Clarity', _score!['clarity_score'], Colors.blue),
                  _scoreBar('Focus', _score!['focus_score'], Colors.orange),
                  _scoreBar('Economics', _score!['economics_score'], Colors.green),
                  _scoreBar('Execution', _score!['execution_score'], Colors.purple),
                  _scoreBar('Evidence', _score!['evidence_score'], Colors.teal),
                  _scoreBar('Market Response', _score!['market_response_score'], Colors.pink),
                  const SizedBox(height: 24),

                  if (_decision == null)
                    SizedBox(
                      width: double.infinity, height: 48,
                      child: ElevatedButton(
                        onPressed: generateDecision,
                        child: const Text('Hasilkan Keputusan Akhir'),
                      ),
                    ),
                ] else ...[
                  const Center(child: Text('Gagal memuat skor')),
                ],

                if (_decision != null) ...[
                  const Divider(height: 32),
                  Container(
                    width: double.infinity,
                    padding: const EdgeInsets.all(24),
                    decoration: BoxDecoration(
                      color: _decisionColor(_decision!['decision']).withValues(alpha: 0.1),
                      borderRadius: BorderRadius.circular(16),
                      border: Border.all(color: _decisionColor(_decision!['decision'])),
                    ),
                    child: Column(children: [
                      Icon(
                        _decision!['decision'] == 'continue' ? Icons.celebration :
                        _decision!['decision'] == 'stop' ? Icons.cancel : Icons.warning,
                        size: 48, color: _decisionColor(_decision!['decision']),
                      ),
                      const SizedBox(height: 16),
                      Text(
                        _decision!['decision'] == 'continue' ? 'LANJUTKAN! 🚀' :
                        _decision!['decision'] == 'repeat' ? 'ULANGI 🔄' :
                        _decision!['decision'] == 'pivot' ? 'PIVOT 🔀' : 'STOP 🛑',
                        style: theme.textTheme.headlineSmall?.copyWith(
                          fontWeight: FontWeight.bold, color: _decisionColor(_decision!['decision']),
                        ),
                      ),
                      const SizedBox(height: 12),
                      Text(_decision!['rationale'] ?? '', textAlign: TextAlign.center, style: TextStyle(color: Colors.grey[700])),
                      const SizedBox(height: 24),
                      OutlinedButton(
                        onPressed: () => context.go('/dashboard'),
                        child: const Text('Kembali ke Dashboard'),
                      ),
                    ]),
                  ),
                ],
              ]),
            ),
    );
  }

  Widget _scoreBar(String label, dynamic scoreVal, Color color) {
    final score = (scoreVal as num).toDouble();
    return Padding(
      padding: const EdgeInsets.only(bottom: 12),
      child: Column(crossAxisAlignment: CrossAxisAlignment.start, children: [
        Row(mainAxisAlignment: MainAxisAlignment.spaceBetween, children: [
          Text(label, style: TextStyle(color: Colors.grey[600], fontSize: 13)),
          Text('${score.toInt()}', style: const TextStyle(fontWeight: FontWeight.w600)),
        ]),
        const SizedBox(height: 4),
        ClipRRect(
          borderRadius: BorderRadius.circular(4),
          child: LinearProgressIndicator(
            value: score / 100,
            backgroundColor: Colors.grey[200],
            color: color,
            minHeight: 8,
          ),
        ),
      ]),
    );
  }

  Color _scoreColor(num s) {
    if (s >= 70) return Colors.green;
    if (s >= 40) return Colors.orange;
    return Colors.red;
  }

  Color _decisionColor(String d) {
    switch (d) {
      case 'continue': return Colors.green;
      case 'repeat': return Colors.orange;
      case 'pivot': return Colors.amber;
      case 'stop': return Colors.red;
      default: return Colors.grey;
    }
  }
}
