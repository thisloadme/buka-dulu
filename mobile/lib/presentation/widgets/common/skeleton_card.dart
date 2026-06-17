import 'package:flutter/material.dart';

/// A skeleton loading card with a pulse animation.
///
/// Displays a grey rounded rectangle that animates opacity between 0.3 and 1.0
/// to indicate content is loading. Mimics the shape of a card in the UI.
class SkeletonCard extends StatefulWidget {
  const SkeletonCard({
    super.key,
    this.width = double.infinity,
    this.height = 100,
    this.borderRadius = 8,
  });

  /// Width of the skeleton card. Defaults to [double.infinity].
  final double width;

  /// Height of the skeleton card. Defaults to 100.
  final double height;

  /// Border radius of the skeleton card. Defaults to 8.
  final double borderRadius;

  @override
  State<SkeletonCard> createState() => _SkeletonCardState();
}

class _SkeletonCardState extends State<SkeletonCard>
    with SingleTickerProviderStateMixin {
  late final AnimationController _controller;
  late final Animation<double> _opacityAnimation;

  @override
  void initState() {
    super.initState();

    _controller = AnimationController(
      vsync: this,
      duration: const Duration(milliseconds: 1200),
    )..repeat(reverse: true);

    _opacityAnimation = Tween<double>(begin: 0.3, end: 1.0).animate(
      CurvedAnimation(parent: _controller, curve: Curves.easeInOut),
    );
  }

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return AnimatedBuilder(
      animation: _opacityAnimation,
      builder: (context, child) {
        return Opacity(
          opacity: _opacityAnimation.value,
          child: child,
        );
      },
      child: Container(
        width: widget.width,
        height: widget.height,
        decoration: BoxDecoration(
          color: const Color(0xFFe7e5e4),
          borderRadius: BorderRadius.circular(widget.borderRadius),
        ),
      ),
    );
  }
}

/// A page-level skeleton layout consisting of a column of skeleton cards.
///
/// Mimics a typical content page with a header area and several body cards,
/// providing visual feedback while the actual page content is loading.
class SkeletonPageBody extends StatelessWidget {
  const SkeletonPageBody({super.key});

  @override
  Widget build(BuildContext context) {
    return const Padding(
      padding: EdgeInsets.all(16),
      child: Column(
        children: [
          // Header skeleton
          SkeletonCard(
            height: 180,
            borderRadius: 8,
          ),
          SizedBox(height: 20),
          // Body card 1
          SkeletonCard(
            height: 80,
            borderRadius: 8,
          ),
          SizedBox(height: 12),
          // Body card 2
          SkeletonCard(
            height: 80,
            borderRadius: 8,
          ),
          SizedBox(height: 12),
          // Body card 3
          SkeletonCard(
            height: 80,
            borderRadius: 8,
          ),
        ],
      ),
    );
  }
}
