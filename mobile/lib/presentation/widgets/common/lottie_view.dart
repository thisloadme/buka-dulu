import 'package:flutter/material.dart';
import 'package:lottie/lottie.dart';

/// Thin wrapper around [Lottie.asset] with sensible defaults.
///
/// `asset` is the filename inside `assets/lottie/` (no path prefix).
/// Use [loop] (default true) to loop the animation forever.
class LottieView extends StatelessWidget {
  final String asset;
  final double size;
  final bool loop;
  final BoxFit fit;

  const LottieView({
    super.key,
    required this.asset,
    this.size = 120,
    this.loop = true,
    this.fit = BoxFit.contain,
  });

  @override
  Widget build(BuildContext context) {
    return SizedBox(
      width: size,
      height: size,
      child: Lottie.asset(
        'assets/lottie/$asset',
        repeat: loop,
        fit: fit,
      ),
    );
  }
}
