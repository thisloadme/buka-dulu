import 'dart:math' as math;
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:bukadulu/data/datasources/api.dart';
import 'package:bukadulu/presentation/widgets/common/brand_icons.dart';

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
                  // Big score circle with animated counter
                  TweenAnimationBuilder<double>(
                    tween: Tween<double>(
                      begin: 0,
                      end: (_score!['total_score'] as num).toDouble(),
                    ),
                    duration: const Duration(milliseconds: 1200),
                    curve: Curves.easeOutCubic,
                    builder: (context, animatedScore, _) {
                      final intScore = animatedScore.toInt();
                      final showCelebration = intScore >= 80;
                      return Stack(
                        alignment: Alignment.center,
                        children: [
                          // Celebration dots when score >= 80
                          if (showCelebration) ...List.generate(8, (i) {
                            final angle = (i * 3.14159 * 2) / 8;
                            return TweenAnimationBuilder<double>(
                              tween: Tween<double>(begin: 0, end: 1),
                              duration: Duration(milliseconds: 600 + (i * 60)),
                              curve: Curves.easeOutBack,
                              builder: (_, t, __) {
                                final radius = 100.0 * t;
                                return Positioned(
                                  left: 80 + radius * math.cos(angle) - 4,
                                  top: 80 + radius * math.sin(angle) - 4,
                                  child: Container(
                                    width: 8,
                                    height: 8,
                                    decoration: BoxDecoration(
                                      color: BrandColors.brandOrange
                                          .withValues(alpha: 1 - t * 0.3),
                                      shape: BoxShape.circle,
                                    ),
                                  ),
                                );
                              },
                            );
                          }),
                          Container(
                            width: 160, height: 160,
                            decoration: BoxDecoration(
                              shape: BoxShape.circle,
                              color: BrandColors.brandOrange.withValues(alpha: 0.1),
                              border: Border.all(color: BrandColors.brandOrange, width: 4),
                            ),
                            child: Center(
                              child: Column(
                                mainAxisSize: MainAxisSize.min,
                                children: [
                                  Text('$intScore', style: TextStyle(fontSize: 48, fontWeight: FontWeight.w600, color: BrandColors.brandOrange)),
                                  Text('/100', style: TextStyle(color: BrandColors.body)),
                                ],
                              ),
                            ),
                          ),
                        ],
                      );
                    },
                  ),
                  const SizedBox(height: 32),

                  // Score breakdown — all bars use brand orange
                  _scoreBar('Clarity', _score!['clarity_score']),
                  _scoreBar('Focus', _score!['focus_score']),
                  _scoreBar('Economics', _score!['economics_score']),
                  _scoreBar('Execution', _score!['execution_score']),
                  _scoreBar('Evidence', _score!['evidence_score']),
                  _scoreBar('Market Response', _score!['market_response_score']),
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
                      _decisionIcon(_decision!['decision'], size: 48),
                      const SizedBox(height: 16),
                      Text(
                        _decision!['decision'] == 'continue' ? 'LANJUTKAN' :
                        _decision!['decision'] == 'repeat' ? 'ULANGI' :
                        _decision!['decision'] == 'pivot' ? 'PIVOT' : 'STOP',
                        style: theme.textTheme.headlineSmall?.copyWith(
                          fontWeight: FontWeight.w600, color: _decisionColor(_decision!['decision']),
                        ),
                      ),
                      const SizedBox(height: 12),
                          Text(_decision!['rationale'] ?? '', textAlign: TextAlign.center, style: TextStyle(color: BrandColors.body)),
                      const SizedBox(height: 16),
                      OutlinedButton(
                        onPressed: () => context.go('/venture/${widget.ventureId}/decision'),
                        child: const Text('Lihat Detail'),
                      ),
                      const SizedBox(height: 8),
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

  Widget _scoreBar(String label, dynamic scoreVal) {
    final score = (scoreVal as num).toDouble();
    return Padding(
      padding: const EdgeInsets.only(bottom: 12),
      child: Column(crossAxisAlignment: CrossAxisAlignment.start, children: [
        Row(mainAxisAlignment: MainAxisAlignment.spaceBetween, children: [
          Text(label, style: TextStyle(color: BrandColors.body, fontSize: 13)),
          // Animated score number (count up to value)
          TweenAnimationBuilder<double>(
            tween: Tween<double>(begin: 0, end: score),
            duration: const Duration(milliseconds: 1000),
            curve: Curves.easeOutCubic,
            builder: (_, v, __) => Text(
              '${v.toInt()}',
              style: const TextStyle(fontWeight: FontWeight.w600),
            ),
          ),
        ]),
        const SizedBox(height: 4),
        // Animated progress bar
        TweenAnimationBuilder<double>(
          tween: Tween<double>(begin: 0, end: score / 100),
          duration: const Duration(milliseconds: 1000),
          curve: Curves.easeOutCubic,
          builder: (_, v, __) => ClipRRect(
            borderRadius: BorderRadius.circular(4),
            child: LinearProgressIndicator(
              value: v,
              backgroundColor: BrandColors.borderLight,
              color: BrandColors.brandOrange,
              minHeight: 8,
            ),
          ),
        ),
      ]),
    );
  }

  Widget _decisionIcon(String decision, {double size = 48}) {
    switch (decision) {
      case 'continue':
        return BrandIcons.celebration(size: size, color: _decisionColor(decision));
      case 'repeat':
        return BrandIcons.refresh(size: size, color: _decisionColor(decision));
      case 'pivot':
        return BrandIcons.flag(size: size, color: _decisionColor(decision));
      case 'stop':
        return BrandIcons.xCircle(size: size, color: _decisionColor(decision));
      default:
        return BrandIcons.celebration(size: size, color: _decisionColor(decision));
    }
  }

  Color _decisionColor(String d) {
    switch (d) {
      case 'continue': return BrandColors.success;
      case 'repeat': return BrandColors.warning;
      case 'pivot': return BrandColors.brandAmber;
      case 'stop': return BrandColors.danger;
      default: return BrandColors.disabled;
    }
  }
}
