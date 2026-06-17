import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:bukadulu/data/datasources/api.dart';
import 'package:bukadulu/presentation/widgets/common/brand_icons.dart';

class DecisionPage extends ConsumerStatefulWidget {
  final String ventureId;
  const DecisionPage({super.key, required this.ventureId});

  @override
  ConsumerState<DecisionPage> createState() => _DecisionPageState();
}

class _DecisionPageState extends ConsumerState<DecisionPage> {
  Map<String, dynamic>? _decision;
  Map<String, dynamic>? _score;
  bool _loading = true;
  String? _error;

  @override
  void initState() {
    super.initState();
    _load();
  }

  Future<void> _load() async {
    setState(() {
      _loading = true;
      _error = null;
    });
    try {
      final api = ref.read(authApiProvider);
      final d = await api.getDecision(widget.ventureId);
      final s = await api.calculateScore(widget.ventureId);
      setState(() {
        _decision = d['decision'] as Map<String, dynamic>?;
        _score = s;
      });
    } catch (e) {
      setState(() => _error = 'Gagal memuat keputusan');
    }
    setState(() => _loading = false);
  }

  String get _decisionKey => _decision?['decision'] as String? ?? '';

  String get _decisionLabel {
    switch (_decisionKey) {
      case 'continue':
        return 'LANJUTKAN';
      case 'repeat':
        return 'ULANGI';
      case 'pivot':
        return 'PIVOT';
      case 'stop':
        return 'STOP';
      default:
        return 'KEPUTUSAN';
    }
  }

  Color get _decisionColor {
    switch (_decisionKey) {
      case 'continue':
        return const Color(0xFF22c55e);
      case 'repeat':
        return const Color(0xFFf59e0b);
      case 'pivot':
        return const Color(0xFFea580c);
      case 'stop':
        return const Color(0xFFef4444);
      default:
        return const Color(0xFF57534e);
    }
  }

  Widget _decisionIcon({double size = 80}) {
    switch (_decisionKey) {
      case 'continue':
        return BrandIcons.celebration(size: size, color: _decisionColor);
      case 'repeat':
        return BrandIcons.refresh(size: size, color: _decisionColor);
      case 'pivot':
        return BrandIcons.flag(size: size, color: _decisionColor);
      case 'stop':
        return BrandIcons.xCircle(size: size, color: _decisionColor);
      default:
        return BrandIcons.celebration(size: size, color: _decisionColor);
    }
  }

  List<String> get _nextSteps {
    final raw = _decision?['next_steps'];
    if (raw is List) {
      return raw.map((e) => e.toString()).toList();
    }
    // Fallback default steps per decision type
    switch (_decisionKey) {
      case 'continue':
        return [
          'Finalisasi rencana bisnis',
          'Siapkan dokumen legalitas',
          'Cari investor awal',
        ];
      case 'repeat':
        return [
          'Review kembali area yang lemah',
          'Kumpulkan data tambahan',
          'Perbaiki skor sebelum lanjut',
        ];
      case 'pivot':
        return [
          'Identifikasi arah baru yang lebih menjanjikan',
          'Validasi asumsi baru dengan target pasar',
          'Sesuaikan rencana bisnis',
        ];
      case 'stop':
        return [
          'Evaluasi pembelajaran dari venture ini',
          'Simpan insight untuk ide berikutnya',
          'Pertimbangkan venture baru',
        ];
      default:
        return [];
    }
  }

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);

    return Scaffold(
      appBar: AppBar(title: const Text('Keputusan Akhir')),
      body: _loading
          ? const Center(child: CircularProgressIndicator())
          : _error != null
              ? Center(
                  child: Column(
                    mainAxisSize: MainAxisSize.min,
                    children: [
                      BrandIcons.alertTriangle(size: 48, color: const Color(0xFFef4444)),
                      const SizedBox(height: 16),
                      Text(_error!, style: const TextStyle(color: Color(0xFF57534e))),
                      const SizedBox(height: 16),
                      OutlinedButton(
                        onPressed: _load,
                        child: const Text('Coba Lagi'),
                      ),
                    ],
                  ),
                )
              : SingleChildScrollView(
                  padding: const EdgeInsets.all(24),
                  child: Column(
                    children: [
                      const SizedBox(height: 16),

                      // Large decision icon
                      _decisionIcon(size: 80),
                      const SizedBox(height: 20),

                      // Decision heading — weight 400 as specified
                      Text(
                        _decisionLabel,
                        style: theme.textTheme.headlineMedium?.copyWith(
                          fontWeight: FontWeight.w400,
                          color: _decisionColor,
                        ),
                      ),
                      const SizedBox(height: 24),

                      // Rationale
                      if (_decision?['rationale'] != null) ...[
                        Text(
                          _decision!['rationale'] as String,
                          textAlign: TextAlign.center,
                          style: const TextStyle(
                            fontSize: 15,
                            color: Color(0xFF57534e),
                            height: 1.5,
                          ),
                        ),
                        const SizedBox(height: 32),
                      ],

                      // Next steps section
                      if (_nextSteps.isNotEmpty) ...[
                        Align(
                          alignment: Alignment.centerLeft,
                          child: Text(
                            'Langkah Selanjutnya',
                            style: theme.textTheme.titleMedium?.copyWith(
                              fontWeight: FontWeight.w600,
                              color: const Color(0xFF292524),
                            ),
                          ),
                        ),
                        const SizedBox(height: 12),
                        ..._nextSteps.asMap().entries.map(
                          (entry) => Padding(
                            padding: const EdgeInsets.only(bottom: 10),
                            child: Row(
                              crossAxisAlignment: CrossAxisAlignment.start,
                              children: [
                                Container(
                                  width: 24,
                                  height: 24,
                                  alignment: Alignment.center,
                                  decoration: BoxDecoration(
                                    color: _decisionColor.withValues(alpha: 0.15),
                                    borderRadius: BorderRadius.circular(12),
                                  ),
                                  child: Text(
                                    '${entry.key + 1}',
                                    style: TextStyle(
                                      fontSize: 12,
                                      fontWeight: FontWeight.w600,
                                      color: _decisionColor,
                                    ),
                                  ),
                                ),
                                const SizedBox(width: 12),
                                Expanded(
                                  child: Text(
                                    entry.value,
                                    style: const TextStyle(
                                      fontSize: 14,
                                      color: Color(0xFF57534e),
                                      height: 1.4,
                                    ),
                                  ),
                                ),
                              ],
                            ),
                          ),
                        ),
                        const SizedBox(height: 32),
                      ],

                      // Score detail
                      if (_score != null) ...[
                        Container(
                          width: double.infinity,
                          padding: const EdgeInsets.all(20),
                          decoration: BoxDecoration(
                            color: const Color(0xFFf5f5f4),
                            borderRadius: BorderRadius.circular(12),
                          ),
                          child: Row(
                            children: [
                              BrandIcons.trendingUp(size: 32, color: const Color(0xFFea580c)),
                              const SizedBox(width: 16),
                              Expanded(
                                child: Column(
                                  crossAxisAlignment: CrossAxisAlignment.start,
                                  children: [
                                    const Text(
                                      'Skor Kesiapan',
                                      style: TextStyle(
                                        fontSize: 13,
                                        color: Color(0xFF57534e),
                                      ),
                                    ),
                                    const SizedBox(height: 4),
                                    Text(
                                      '${(_score!['total_score'] as num).toInt()} / 100',
                                      style: const TextStyle(
                                        fontSize: 22,
                                        fontWeight: FontWeight.w600,
                                        color: Color(0xFF292524),
                                      ),
                                    ),
                                  ],
                                ),
                              ),
                            ],
                          ),
                        ),
                        const SizedBox(height: 24),
                      ],

                      // CTA: Lihat Detail Skor
                      SizedBox(
                        width: double.infinity,
                        height: 48,
                        child: ElevatedButton(
                          onPressed: () => context.go('/venture/${widget.ventureId}/score'),
                          style: ElevatedButton.styleFrom(
                            backgroundColor: const Color(0xFFea580c),
                            foregroundColor: Colors.white,
                            shape: RoundedRectangleBorder(
                              borderRadius: BorderRadius.circular(12),
                            ),
                          ),
                          child: const Text(
                            'Lihat Detail Skor',
                            style: TextStyle(fontWeight: FontWeight.w600),
                          ),
                        ),
                      ),
                      const SizedBox(height: 12),

                      // CTA: Mulai Venture Baru
                      SizedBox(
                        width: double.infinity,
                        height: 48,
                        child: OutlinedButton(
                          onPressed: () => context.go('/venture/new'),
                          style: OutlinedButton.styleFrom(
                            foregroundColor: const Color(0xFFea580c),
                            side: const BorderSide(color: Color(0xFFea580c)),
                            shape: RoundedRectangleBorder(
                              borderRadius: BorderRadius.circular(12),
                            ),
                          ),
                          child: const Text(
                            'Mulai Venture Baru',
                            style: TextStyle(fontWeight: FontWeight.w600),
                          ),
                        ),
                      ),
                      const SizedBox(height: 12),

                      // Cari Mentor — text link
                      TextButton(
                        onPressed: () {
                          // Informational — mentor feature coming soon
                          if (mounted) {
                            ScaffoldMessenger.of(context).showSnackBar(
                              const SnackBar(content: Text('Fitur Cari Mentor akan segera hadir')),
                            );
                          }
                        },
                        child: const Text(
                          'Cari Mentor',
                          style: TextStyle(
                            color: Color(0xFF292524),
                            decoration: TextDecoration.underline,
                          ),
                        ),
                      ),
                      const SizedBox(height: 24),
                    ],
                  ),
                ),
    );
  }
}
