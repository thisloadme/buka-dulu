import 'package:flutter/material.dart';
import 'package:go_router/go_router.dart';

class OnboardingPage extends StatefulWidget {
  const OnboardingPage({super.key});

  @override
  State<OnboardingPage> createState() => _OnboardingPageState();
}

class _OnboardingPageState extends State<OnboardingPage> {
  final _pageController = PageController();
  int _currentPage = 0;

  static const _slides = [
    _SlideData(
      icon: Icons.compress_outlined,
      title: 'Compress Ide Mentahmu',
      desc:
          'Tulis ide bisnis F&B kamu dalam bahasa sehari-hari. AI akan menyusunnya '
          'menjadi konsep bisnis yang tajam, target customer, dan asumsi utama.',
    ),
    _SlideData(
      icon: Icons.track_changes_outlined,
      title: 'Fokus ke 1-3 SKU',
      desc:
          'Pilih maksimal 3 menu yang paling realistis untuk diuji. '
          'Sistem merekomendasikan hero product yang punya peluang terbaik.',
    ),
    _SlideData(
      icon: Icons.calculate_outlined,
      title: 'Hitung Realitas Bisnis',
      desc:
          'Input bahan, harga, dan kemasan. Lihat HPP, margin, dan modal awal '
          'yang kamu butuhkan — bukan perasaan, tapi angka nyata.',
    ),
    _SlideData(
      icon: Icons.flag_outlined,
      title: 'Jalankan Misi, Kumpulkan Bukti',
      desc:
          'Dapatkan tugas harian: pre-order, sampling, atau polling. '
          'Upload bukti sebagai progress. Sistem akan menilai dan memberi '
          'keputusan tegas: Lanjut, Pivot, atau Stop.',
    ),
  ];

  @override
  void dispose() {
    _pageController.dispose();
    super.dispose();
  }

  void _goToLogin() => context.go('/login');
  void _goToRegister() => context.go('/register');

  @override
  Widget build(BuildContext context) {
    final theme = Theme.of(context);

    return Scaffold(
      body: SafeArea(
        child: Column(
          children: [
            // Skip button
            Align(
              alignment: Alignment.topRight,
              child: TextButton(
                onPressed: _goToLogin,
                child: Text(
                  _currentPage < _slides.length - 1 ? 'Skip' : '',
                  style: TextStyle(color: Colors.grey[500]),
                ),
              ),
            ),

            // PageView
            Expanded(
              child: PageView.builder(
                controller: _pageController,
                onPageChanged: (i) => setState(() => _currentPage = i),
                itemCount: _slides.length,
                itemBuilder: (_, i) {
                  final slide = _slides[i];
                  return Padding(
                    padding: const EdgeInsets.symmetric(horizontal: 32),
                    child: Column(
                      mainAxisAlignment: MainAxisAlignment.center,
                      children: [
                        // Icon
                        Container(
                          width: 100,
                          height: 100,
                          decoration: BoxDecoration(
                            color: const Color(0xFFfff7ed),
                            borderRadius: BorderRadius.circular(28),
                          ),
                          child: Icon(
                            slide.icon,
                            size: 48,
                            color: const Color(0xFFea580c),
                          ),
                        ),
                        const SizedBox(height: 40),

                        // Title
                        Text(
                          slide.title,
                          style: theme.textTheme.headlineSmall?.copyWith(
                            fontWeight: FontWeight.w400,
                            letterSpacing: -0.3,
                          ),
                          textAlign: TextAlign.center,
                        ),
                        const SizedBox(height: 16),

                        // Description
                        Text(
                          slide.desc,
                          style: theme.textTheme.bodyLarge?.copyWith(
                            color: Colors.grey[600],
                            fontWeight: FontWeight.w300,
                            height: 1.5,
                          ),
                          textAlign: TextAlign.center,
                        ),
                      ],
                    ),
                  );
                },
              ),
            ),

            // Dots indicator
            Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: List.generate(
                _slides.length,
                (i) => AnimatedContainer(
                  duration: const Duration(milliseconds: 300),
                  margin: const EdgeInsets.symmetric(horizontal: 4),
                  width: _currentPage == i ? 24 : 8,
                  height: 8,
                  decoration: BoxDecoration(
                    borderRadius: BorderRadius.circular(4),
                    color: _currentPage == i
                        ? const Color(0xFFea580c)
                        : Colors.grey[300],
                  ),
                ),
              ),
            ),
            const SizedBox(height: 32),

            // Bottom buttons
            Padding(
              padding: const EdgeInsets.fromLTRB(24, 0, 24, 32),
              child: _currentPage == _slides.length - 1
                  ? Column(
                      children: [
                        SizedBox(
                          width: double.infinity,
                          height: 50,
                          child: ElevatedButton(
                            onPressed: _goToRegister,
                            style: ElevatedButton.styleFrom(
                              backgroundColor: const Color(0xFFea580c),
                              foregroundColor: Colors.white,
                              shape: RoundedRectangleBorder(
                                borderRadius: BorderRadius.circular(10),
                              ),
                              elevation: 0,
                            ),
                            child: const Text(
                              'Mulai Sekarang',
                              style: TextStyle(
                                fontSize: 16,
                                fontWeight: FontWeight.w500,
                              ),
                            ),
                          ),
                        ),
                        const SizedBox(height: 12),
                        TextButton(
                          onPressed: _goToLogin,
                          child: const Text(
                            'Sudah punya akun? Masuk',
                            style: TextStyle(color: Color(0xFFea580c)),
                          ),
                        ),
                      ],
                    )
                  : SizedBox(
                      width: double.infinity,
                      height: 50,
                      child: ElevatedButton(
                        onPressed: () {
                          _pageController.nextPage(
                            duration: const Duration(milliseconds: 400),
                            curve: Curves.easeInOut,
                          );
                        },
                        style: ElevatedButton.styleFrom(
                          backgroundColor: const Color(0xFFea580c),
                          foregroundColor: Colors.white,
                          shape: RoundedRectangleBorder(
                            borderRadius: BorderRadius.circular(10),
                          ),
                          elevation: 0,
                        ),
                        child: const Text(
                          'Lanjut',
                          style: TextStyle(
                            fontSize: 16,
                            fontWeight: FontWeight.w500,
                          ),
                        ),
                      ),
                    ),
            ),
          ],
        ),
      ),
    );
  }
}

class _SlideData {
  final IconData icon;
  final String title;
  final String desc;

  const _SlideData({
    required this.icon,
    required this.title,
    required this.desc,
  });
}
