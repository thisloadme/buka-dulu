import 'dart:math' as math;

import 'package:flutter/material.dart';

/// Brand colors used throughout BukaDulu.
class BrandColors {
  const BrandColors._();

  static const Color brandOrange = Color(0xFFea580c);
  static const Color brandAmber = Color(0xFFf59e0b);
  static const Color brandCream = Color(0xFFfff7ed);
  static const Color brandDark = Color(0xFF1c1917);
  static const Color body = Color(0xFF57534e);
  static const Color label = Color(0xFF292524);
  static const Color border = Color(0xFFe7e5e4);
  static const Color borderLight = Color(0xFFf5f5f4);
  static const Color bgWarm = Color(0xFFfafaf9);
}

/// Convenience semantic palette for icon status colours.
class _IconPalette {
  const _IconPalette._();
  static const Color valid = Color(0xFF22c55e); // green
  static const Color warning = Color(0xFFf59e0b); // amber
  static const Color invalid = Color(0xFFef4444); // red
  static const Color star = Color(0xFFf59e0b); // amber fill
}

// =============================================================================
// BrandIcons – static getters returning ready-to-use Widgets
// =============================================================================

/// SVG-like brand icons implemented via [CustomPaint] with [Path] shapes.
///
/// Every getter returns a [StatelessWidget] that accepts optional `size`
/// (default 24) and `color` parameters via builder methods.  Use directly in
/// widget trees:
///
/// ```dart
/// BrandIcons.checkCircle(size: 20)
/// BrandIcons.alertTriangle(color: BrandColors.brandAmber)
/// ```
///
/// Default colours follow a semantic palette:
///   • checkCircle  → green   (valid/complete)
///   • alertTriangle → amber  (warning)
///   • xCircle      → red     (invalid/stop)
///   • star         → amber fill
///   • all others   → brandDark
class BrandIcons {
  BrandIcons._();

  static const double kDefaultSize = 24.0;

  // ──────────────────────────────────────────────────────────────────────────
  // Status icons
  // ──────────────────────────────────────────────────────────────────────────

  static Widget checkCircle({double size = kDefaultSize, Color? color}) =>
      _iconWidget<_CheckCirclePainter>(
        size: size,
        color: color ?? _IconPalette.valid,
        label: 'Check circle',
      );

  static Widget alertTriangle({double size = kDefaultSize, Color? color}) =>
      _iconWidget<_AlertTrianglePainter>(
        size: size,
        color: color ?? _IconPalette.warning,
        label: 'Alert triangle',
      );

  static Widget xCircle({double size = kDefaultSize, Color? color}) =>
      _iconWidget<_XCirclePainter>(
        size: size,
        color: color ?? _IconPalette.invalid,
        label: 'X circle',
      );

  static Widget star({double size = kDefaultSize, Color? color}) =>
      _iconWidget<_StarPainter>(
        size: size,
        color: color ?? _IconPalette.star,
        label: 'Star',
      );

  // ──────────────────────────────────────────────────────────────────────────
  // Action / Score icons
  // ──────────────────────────────────────────────────────────────────────────

  static Widget trendingUp({double size = kDefaultSize, Color? color}) =>
      _iconWidget<_TrendingUpPainter>(
        size: size,
        color: color ?? BrandColors.brandDark,
        label: 'Trending up',
      );

  static Widget arrowUp({double size = kDefaultSize, Color? color}) =>
      _iconWidget<_ArrowUpPainter>(
        size: size,
        color: color ?? BrandColors.brandDark,
        label: 'Arrow up',
      );

  // ──────────────────────────────────────────────────────────────────────────
  // Time icons
  // ──────────────────────────────────────────────────────────────────────────

  static Widget clock({double size = kDefaultSize, Color? color}) =>
      _iconWidget<_ClockPainter>(
        size: size,
        color: color ?? BrandColors.brandDark,
        label: 'Clock',
      );

  // ──────────────────────────────────────────────────────────────────────────
  // Mission / Outcome icons
  // ──────────────────────────────────────────────────────────────────────────

  static Widget flag({double size = kDefaultSize, Color? color}) =>
      _iconWidget<_FlagPainter>(
        size: size,
        color: color ?? BrandColors.brandDark,
        label: 'Flag',
      );

  static Widget locationPin({double size = kDefaultSize, Color? color}) =>
      _iconWidget<_LocationPinPainter>(
        size: size,
        color: color ?? BrandColors.brandDark,
        label: 'Location pin',
      );

  static Widget chatBubble({double size = kDefaultSize, Color? color}) =>
      _iconWidget<_ChatBubblePainter>(
        size: size,
        color: color ?? BrandColors.brandDark,
        label: 'Chat bubble',
      );

  // ──────────────────────────────────────────────────────────────────────────
  // Mission-type icons
  // ──────────────────────────────────────────────────────────────────────────

  static Widget barChart({double size = kDefaultSize, Color? color}) =>
      _iconWidget<_BarChartPainter>(
        size: size,
        color: color ?? BrandColors.brandDark,
        label: 'Bar chart',
      );

  static Widget restaurant({double size = kDefaultSize, Color? color}) =>
      _iconWidget<_RestaurantPainter>(
        size: size,
        color: color ?? BrandColors.brandDark,
        label: 'Restaurant',
      );

  static Widget camera({double size = kDefaultSize, Color? color}) =>
      _iconWidget<_CameraPainter>(
        size: size,
        color: color ?? BrandColors.brandDark,
        label: 'Camera',
      );

  // ──────────────────────────────────────────────────────────────────────────
  // Evidence icons
  // ──────────────────────────────────────────────────────────────────────────

  static Widget link({double size = kDefaultSize, Color? color}) =>
      _iconWidget<_LinkPainter>(
        size: size,
        color: color ?? BrandColors.brandDark,
        label: 'Link',
      );

  static Widget textFile({double size = kDefaultSize, Color? color}) =>
      _iconWidget<_TextFilePainter>(
        size: size,
        color: color ?? BrandColors.brandDark,
        label: 'Text file',
      );

  // ──────────────────────────────────────────────────────────────────────────
  // Decision icons
  // ──────────────────────────────────────────────────────────────────────────

  static Widget celebration({double size = kDefaultSize, Color? color}) =>
      _iconWidget<_CelebrationPainter>(
        size: size,
        color: color ?? BrandColors.brandDark,
        label: 'Celebration',
      );

  static Widget refresh({double size = kDefaultSize, Color? color}) =>
      _iconWidget<_RefreshPainter>(
        size: size,
        color: color ?? BrandColors.brandDark,
        label: 'Refresh',
      );

  // ──────────────────────────────────────────────────────────────────────────
  // Internal helpers
  // ──────────────────────────────────────────────────────────────────────────

  static Widget _iconWidget<T extends CustomPainter>({
    required double size,
    required Color color,
    required String label,
  }) =>
      _BrandIconWidget(
        size: size,
        painter: _PainterFactory<T>(color) as CustomPainter,
        label: label,
      );
}

// =============================================================================
// Generic factory that creates a painter of type T via reflection.
//
// We use a trick: store the colour and the Type, then at build time we
// reflectively construct the right painter.  Because we control all the
// painters and they all take a single Color argument, this is safe.
// =============================================================================

class _PainterFactory<T extends CustomPainter> extends CustomPainter {
  _PainterFactory(this.color);

  final Color color;

  @override
  void paint(Canvas canvas, Size size) {
    // Each painter type is instantiated here.  This is the only place that
    // needs updating when a new painter is added.
    final CustomPainter painter;
    final type = T;
    if (type == _CheckCirclePainter) {
      painter = _CheckCirclePainter(color);
    } else if (type == _AlertTrianglePainter) {
      painter = _AlertTrianglePainter(color);
    } else if (type == _XCirclePainter) {
      painter = _XCirclePainter(color);
    } else if (type == _StarPainter) {
      painter = _StarPainter(color);
    } else if (type == _TrendingUpPainter) {
      painter = _TrendingUpPainter(color);
    } else if (type == _ArrowUpPainter) {
      painter = _ArrowUpPainter(color);
    } else if (type == _ClockPainter) {
      painter = _ClockPainter(color);
    } else if (type == _FlagPainter) {
      painter = _FlagPainter(color);
    } else if (type == _LocationPinPainter) {
      painter = _LocationPinPainter(color);
    } else if (type == _ChatBubblePainter) {
      painter = _ChatBubblePainter(color);
    } else if (type == _BarChartPainter) {
      painter = _BarChartPainter(color);
    } else if (type == _RestaurantPainter) {
      painter = _RestaurantPainter(color);
    } else if (type == _CameraPainter) {
      painter = _CameraPainter(color);
    } else if (type == _LinkPainter) {
      painter = _LinkPainter(color);
    } else if (type == _TextFilePainter) {
      painter = _TextFilePainter(color);
    } else if (type == _CelebrationPainter) {
      painter = _CelebrationPainter(color);
    } else if (type == _RefreshPainter) {
      painter = _RefreshPainter(color);
    } else {
      painter = _CheckCirclePainter(color);
    }
    painter.paint(canvas, size);
  }

  @override
  bool shouldRepaint(covariant _PainterFactory<T> old) => old.color != color;
}

// =============================================================================
// Internal widget
// =============================================================================

class _BrandIconWidget extends StatelessWidget {
  final double size;
  final CustomPainter painter;
  final String label;

  const _BrandIconWidget({
    required this.size,
    required this.painter,
    required this.label,
  });

  @override
  Widget build(BuildContext context) {
    return Semantics(
      label: label,
      child: SizedBox(
        width: size,
        height: size,
        child: CustomPaint(
          size: Size.infinite,
          painter: painter,
        ),
      ),
    );
  }
}

// =============================================================================
// Icon painters — each draws a single icon in a 24×24 viewport
// =============================================================================

/// Creates a [Paint] configured for stroke-style drawing.
Paint _strokePaint(Color color) => Paint()
  ..color = color
  ..style = PaintingStyle.stroke
  ..strokeWidth = 2.0
  ..strokeCap = StrokeCap.round
  ..strokeJoin = StrokeJoin.round;

/// Creates a [Paint] configured for fill-style drawing.
Paint _fillPaint(Color color) => Paint()
  ..color = color
  ..style = PaintingStyle.fill;

// ─────────────────────────────────────────────────────────────────────────────
// 1. CheckCircle – green circle with a white checkmark
// ─────────────────────────────────────────────────────────────────────────────

class _CheckCirclePainter extends CustomPainter {
  _CheckCirclePainter(this.color);
  final Color color;

  @override
  void paint(Canvas canvas, Size size) {
    final s = size.width / 24.0;
    canvas.save();
    canvas.scale(s, s);

    final stroke = _strokePaint(color);
    canvas.drawOval(
      Rect.fromCircle(center: const Offset(12, 12), radius: 10),
      stroke,
    );

    final check = Path()
      ..moveTo(7, 12)
      ..lineTo(10.5, 15.5)
      ..lineTo(17, 9);
    canvas.drawPath(check, stroke);

    canvas.restore();
  }

  @override
  bool shouldRepaint(covariant _CheckCirclePainter old) => old.color != color;
}

// ─────────────────────────────────────────────────────────────────────────────
// 2. AlertTriangle – amber outlined triangle with exclamation
// ─────────────────────────────────────────────────────────────────────────────

class _AlertTrianglePainter extends CustomPainter {
  _AlertTrianglePainter(this.color);
  final Color color;

  @override
  void paint(Canvas canvas, Size size) {
    final s = size.width / 24.0;
    canvas.save();
    canvas.scale(s, s);

    final stroke = _strokePaint(color);

    final triangle = Path()
      ..moveTo(12, 3)
      ..lineTo(3, 21)
      ..lineTo(21, 21)
      ..close();
    canvas.drawPath(triangle, stroke);

    final excl = Path()
      ..moveTo(12, 9)
      ..lineTo(12, 15)
      ..moveTo(12, 18)
      ..lineTo(12, 19);
    canvas.drawPath(excl, stroke);

    canvas.restore();
  }

  @override
  bool shouldRepaint(covariant _AlertTrianglePainter old) =>
      old.color != color;
}

// ─────────────────────────────────────────────────────────────────────────────
// 3. XCircle – red circle with an X
// ─────────────────────────────────────────────────────────────────────────────

class _XCirclePainter extends CustomPainter {
  _XCirclePainter(this.color);
  final Color color;

  @override
  void paint(Canvas canvas, Size size) {
    final s = size.width / 24.0;
    canvas.save();
    canvas.scale(s, s);

    final stroke = _strokePaint(color);

    canvas.drawOval(
      Rect.fromCircle(center: const Offset(12, 12), radius: 10),
      stroke,
    );

    final x = Path()
      ..moveTo(8, 8)
      ..lineTo(16, 16)
      ..moveTo(16, 8)
      ..lineTo(8, 16);
    canvas.drawPath(x, stroke);

    canvas.restore();
  }

  @override
  bool shouldRepaint(covariant _XCirclePainter old) => old.color != color;
}

// ─────────────────────────────────────────────────────────────────────────────
// 4. Star – 5-point filled star (amber by default)
// ─────────────────────────────────────────────────────────────────────────────

class _StarPainter extends CustomPainter {
  _StarPainter(this.color);
  final Color color;

  @override
  void paint(Canvas canvas, Size size) {
    final s = size.width / 24.0;
    canvas.save();
    canvas.scale(s, s);

    final fill = _fillPaint(color);
    const cx = 12.0, cy = 12.0;
    const outerR = 10.0, innerR = 4.0;

    final star = Path();
    for (int i = 0; i < 5; i++) {
      final outerAngle = -math.pi / 2 + i * 2 * math.pi / 5;
      final innerAngle = outerAngle + math.pi / 5;
      final ox = cx + outerR * math.cos(outerAngle);
      final oy = cy + outerR * math.sin(outerAngle);
      final ix = cx + innerR * math.cos(innerAngle);
      final iy = cy + innerR * math.sin(innerAngle);
      if (i == 0) {
        star.moveTo(ox, oy);
      } else {
        star.lineTo(ox, oy);
      }
      star.lineTo(ix, iy);
    }
    star.close();
    canvas.drawPath(star, fill);

    canvas.restore();
  }

  @override
  bool shouldRepaint(covariant _StarPainter old) => old.color != color;
}

// ─────────────────────────────────────────────────────────────────────────────
// 5. TrendingUp – line chart trending upward with arrowhead
// ─────────────────────────────────────────────────────────────────────────────

class _TrendingUpPainter extends CustomPainter {
  _TrendingUpPainter(this.color);
  final Color color;

  @override
  void paint(Canvas canvas, Size size) {
    final s = size.width / 24.0;
    canvas.save();
    canvas.scale(s, s);

    final stroke = _strokePaint(color);

    final path = Path()
      ..moveTo(3, 17)
      ..lineTo(8, 12)
      ..lineTo(13, 15)
      ..lineTo(21, 7)
      // arrowhead
      ..moveTo(17, 7)
      ..lineTo(21, 7)
      ..moveTo(21, 7)
      ..lineTo(21, 11);
    canvas.drawPath(path, stroke);

    canvas.restore();
  }

  @override
  bool shouldRepaint(covariant _TrendingUpPainter old) => old.color != color;
}

// ─────────────────────────────────────────────────────────────────────────────
// 6. ArrowUp – simple upward arrow (chevron style)
// ─────────────────────────────────────────────────────────────────────────────

class _ArrowUpPainter extends CustomPainter {
  _ArrowUpPainter(this.color);
  final Color color;

  @override
  void paint(Canvas canvas, Size size) {
    final s = size.width / 24.0;
    canvas.save();
    canvas.scale(s, s);

    final stroke = _strokePaint(color);

    final path = Path()
      ..moveTo(12, 3)
      ..lineTo(5, 11)
      ..moveTo(12, 3)
      ..lineTo(19, 11)
      ..moveTo(12, 3)
      ..lineTo(12, 22);
    canvas.drawPath(path, stroke);

    canvas.restore();
  }

  @override
  bool shouldRepaint(covariant _ArrowUpPainter old) => old.color != color;
}

// ─────────────────────────────────────────────────────────────────────────────
// 7. Clock – circle with hour/minute hands
// ─────────────────────────────────────────────────────────────────────────────

class _ClockPainter extends CustomPainter {
  _ClockPainter(this.color);
  final Color color;

  @override
  void paint(Canvas canvas, Size size) {
    final s = size.width / 24.0;
    canvas.save();
    canvas.scale(s, s);

    final stroke = _strokePaint(color);

    // Outer ring
    canvas.drawOval(
      Rect.fromCircle(center: const Offset(12, 12), radius: 10),
      stroke,
    );
    // Center dot
    canvas.drawCircle(const Offset(12, 12), 1.5, stroke);

    // Hour hand (short, pointing up-left)
    final hourHand = Path()
      ..moveTo(12, 12)
      ..lineTo(9, 7);
    canvas.drawPath(hourHand, stroke);

    // Minute hand (long, pointing right)
    final minHand = Path()
      ..moveTo(12, 12)
      ..lineTo(17, 13);
    canvas.drawPath(minHand, stroke);

    canvas.restore();
  }

  @override
  bool shouldRepaint(covariant _ClockPainter old) => old.color != color;
}

// ─────────────────────────────────────────────────────────────────────────────
// 8. Flag – pole with waving flag
// ─────────────────────────────────────────────────────────────────────────────

class _FlagPainter extends CustomPainter {
  _FlagPainter(this.color);
  final Color color;

  @override
  void paint(Canvas canvas, Size size) {
    final s = size.width / 24.0;
    canvas.save();
    canvas.scale(s, s);

    final stroke = _strokePaint(color);

    // Pole
    canvas.drawLine(const Offset(5, 2), const Offset(5, 22), stroke);

    // Flag
    final flag = Path()
      ..moveTo(5, 4)
      ..lineTo(19, 4)
      ..lineTo(19, 13)
      ..lineTo(5, 13)
      ..close();
    canvas.drawPath(flag, stroke);

    canvas.restore();
  }

  @override
  bool shouldRepaint(covariant _FlagPainter old) => old.color != color;
}

// ─────────────────────────────────────────────────────────────────────────────
// 9. LocationPin – map pin (teardrop + inner circle)
// ─────────────────────────────────────────────────────────────────────────────

class _LocationPinPainter extends CustomPainter {
  _LocationPinPainter(this.color);
  final Color color;

  @override
  void paint(Canvas canvas, Size size) {
    final s = size.width / 24.0;
    canvas.save();
    canvas.scale(s, s);

    final stroke = _strokePaint(color);

    // Teardrop outline
    final pin = Path()
      ..moveTo(12, 2)
      ..quadraticBezierTo(4, 8, 12, 22)
      ..quadraticBezierTo(20, 8, 12, 2)
      ..close();
    canvas.drawPath(pin, stroke);

    // Inner circle
    canvas.drawCircle(const Offset(12, 10), 3.5, stroke);

    canvas.restore();
  }

  @override
  bool shouldRepaint(covariant _LocationPinPainter old) => old.color != color;
}

// ─────────────────────────────────────────────────────────────────────────────
// 10. ChatBubble – rounded rect with tail
// ─────────────────────────────────────────────────────────────────────────────

class _ChatBubblePainter extends CustomPainter {
  _ChatBubblePainter(this.color);
  final Color color;

  @override
  void paint(Canvas canvas, Size size) {
    final s = size.width / 24.0;
    canvas.save();
    canvas.scale(s, s);

    final stroke = _strokePaint(color);

    final bubble = Path()
      ..moveTo(5, 3)
      ..lineTo(19, 3)
      ..quadraticBezierTo(22, 3, 22, 6)
      ..lineTo(22, 15)
      ..quadraticBezierTo(22, 18, 19, 18)
      ..lineTo(10, 18)
      ..lineTo(6, 22)
      ..lineTo(6, 18)
      ..lineTo(5, 18)
      ..quadraticBezierTo(2, 18, 2, 15)
      ..lineTo(2, 6)
      ..quadraticBezierTo(2, 3, 5, 3)
      ..close();
    canvas.drawPath(bubble, stroke);

    canvas.restore();
  }

  @override
  bool shouldRepaint(covariant _ChatBubblePainter old) => old.color != color;
}

// ─────────────────────────────────────────────────────────────────────────────
// 11. BarChart – three bars of varying heights on a baseline
// ─────────────────────────────────────────────────────────────────────────────

class _BarChartPainter extends CustomPainter {
  _BarChartPainter(this.color);
  final Color color;

  @override
  void paint(Canvas canvas, Size size) {
    final s = size.width / 24.0;
    canvas.save();
    canvas.scale(s, s);

    final stroke = _strokePaint(color);

    final path = Path()
      // Baseline
      ..moveTo(3, 20)
      ..lineTo(21, 20)
      // Bar 1 (short)
      ..moveTo(6, 20)
      ..lineTo(6, 12)
      ..lineTo(9, 12)
      ..lineTo(9, 20)
      // Bar 2 (tall)
      ..moveTo(11, 20)
      ..lineTo(11, 4)
      ..lineTo(14, 4)
      ..lineTo(14, 20)
      // Bar 3 (medium)
      ..moveTo(16, 20)
      ..lineTo(16, 8)
      ..lineTo(19, 8)
      ..lineTo(19, 20);
    canvas.drawPath(path, stroke);

    canvas.restore();
  }

  @override
  bool shouldRepaint(covariant _BarChartPainter old) => old.color != color;
}

// ─────────────────────────────────────────────────────────────────────────────
// 12. Restaurant – fork and knife
// ─────────────────────────────────────────────────────────────────────────────

class _RestaurantPainter extends CustomPainter {
  _RestaurantPainter(this.color);
  final Color color;

  @override
  void paint(Canvas canvas, Size size) {
    final s = size.width / 24.0;
    canvas.save();
    canvas.scale(s, s);

    final stroke = _strokePaint(color);

    // Fork (left side)
    final path = Path()
      // Prongs
      ..moveTo(7, 3)
      ..lineTo(7, 12)
      ..moveTo(9, 3)
      ..lineTo(9, 12)
      ..moveTo(11, 3)
      ..lineTo(11, 12)
      // Prong base
      ..moveTo(6, 12)
      ..lineTo(12, 12)
      // Handle
      ..lineTo(12, 16)
      ..lineTo(8, 16)

      // Knife (right side)
      ..moveTo(16, 3)
      ..lineTo(16, 16)
      // Blade curve
      ..moveTo(16, 3)
      ..quadraticBezierTo(21, 3, 21, 8)
      // Handle
      ..lineTo(21, 16)
      ..lineTo(16, 16);
    canvas.drawPath(path, stroke);

    canvas.restore();
  }

  @override
  bool shouldRepaint(covariant _RestaurantPainter old) => old.color != color;
}

// ─────────────────────────────────────────────────────────────────────────────
// 13. Camera – rectangular body with lens and flash
// ─────────────────────────────────────────────────────────────────────────────

class _CameraPainter extends CustomPainter {
  _CameraPainter(this.color);
  final Color color;

  @override
  void paint(Canvas canvas, Size size) {
    final s = size.width / 24.0;
    canvas.save();
    canvas.scale(s, s);

    final stroke = _strokePaint(color);

    // Body
    final body = RRect.fromRectAndRadius(
      Rect.fromLTRB(3, 7, 21, 21),
      const Radius.circular(3),
    );
    canvas.drawRRect(body, stroke);

    // Top hump (flash/viewfinder area)
    final hump = Path()
      ..moveTo(6, 7)
      ..lineTo(6, 5)
      ..lineTo(10, 5)
      ..lineTo(11, 7);
    canvas.drawPath(hump, stroke);

    // Lens outer
    canvas.drawOval(
      Rect.fromCircle(center: const Offset(12, 14), radius: 4),
      stroke,
    );
    // Lens inner
    canvas.drawOval(
      Rect.fromCircle(center: const Offset(12, 14), radius: 2),
      stroke,
    );

    canvas.restore();
  }

  @override
  bool shouldRepaint(covariant _CameraPainter old) => old.color != color;
}

// ─────────────────────────────────────────────────────────────────────────────
// 14. Link – two interlocking chain links
// ─────────────────────────────────────────────────────────────────────────────

class _LinkPainter extends CustomPainter {
  _LinkPainter(this.color);
  final Color color;

  @override
  void paint(Canvas canvas, Size size) {
    final s = size.width / 24.0;
    canvas.save();
    canvas.scale(s, s);

    final stroke = _strokePaint(color);

    // Figure-8 interlocking chain links
    final path = Path()
      // Left link – top-left curve down to bottom-left
      ..moveTo(10, 5)
      ..quadraticBezierTo(4, 5, 4, 12)
      ..quadraticBezierTo(4, 19, 10, 19)
      // Bridge to right link
      ..moveTo(10, 19)
      ..quadraticBezierTo(10, 15, 12, 15)
      ..quadraticBezierTo(14, 15, 14, 19)
      // Right link – bottom-right curve up to top-right
      ..moveTo(14, 19)
      ..quadraticBezierTo(20, 19, 20, 12)
      ..quadraticBezierTo(20, 5, 14, 5)
      // Bridge back to left link
      ..moveTo(14, 5)
      ..quadraticBezierTo(14, 9, 12, 9)
      ..quadraticBezierTo(10, 9, 10, 5);
    canvas.drawPath(path, stroke);

    canvas.restore();
  }

  @override
  bool shouldRepaint(covariant _LinkPainter old) => old.color != color;
}

// ─────────────────────────────────────────────────────────────────────────────
// 15. TextFile – document page with folded corner and text lines
// ─────────────────────────────────────────────────────────────────────────────

class _TextFilePainter extends CustomPainter {
  _TextFilePainter(this.color);
  final Color color;

  @override
  void paint(Canvas canvas, Size size) {
    final s = size.width / 24.0;
    canvas.save();
    canvas.scale(s, s);

    final stroke = _strokePaint(color);

    // Document body
    final doc = Path()
      ..moveTo(5, 2)
      ..lineTo(15, 2)
      ..lineTo(19, 6)
      ..lineTo(19, 22)
      ..lineTo(5, 22)
      ..close();
    canvas.drawPath(doc, stroke);

    // Fold corner
    final fold = Path()
      ..moveTo(15, 2)
      ..lineTo(15, 6)
      ..lineTo(19, 6);
    canvas.drawPath(fold, stroke);

    // Text lines
    final lines = Path()
      ..moveTo(8, 10)
      ..lineTo(16, 10)
      ..moveTo(8, 14)
      ..lineTo(16, 14)
      ..moveTo(8, 18)
      ..lineTo(13, 18);
    canvas.drawPath(lines, stroke);

    canvas.restore();
  }

  @override
  bool shouldRepaint(covariant _TextFilePainter old) => old.color != color;
}

// ─────────────────────────────────────────────────────────────────────────────
// 16. Celebration – radiating sparkle / burst lines
// ─────────────────────────────────────────────────────────────────────────────

class _CelebrationPainter extends CustomPainter {
  _CelebrationPainter(this.color);
  final Color color;

  @override
  void paint(Canvas canvas, Size size) {
    final s = size.width / 24.0;
    canvas.save();
    canvas.scale(s, s);

    final stroke = _strokePaint(color);

    final path = Path()
      // Center cross
      ..moveTo(12, 2)
      ..lineTo(12, 6)
      ..moveTo(12, 18)
      ..lineTo(12, 22)
      ..moveTo(2, 12)
      ..lineTo(6, 12)
      ..moveTo(18, 12)
      ..lineTo(22, 12)
      // Diagonals
      ..moveTo(5, 5)
      ..lineTo(8, 8)
      ..moveTo(16, 16)
      ..lineTo(19, 19)
      ..moveTo(19, 5)
      ..lineTo(16, 8)
      ..moveTo(5, 19)
      ..lineTo(8, 16)
      // Small accent dashes
      ..moveTo(10, 4)
      ..lineTo(10, 5)
      ..moveTo(4, 10)
      ..lineTo(5, 10)
      ..moveTo(20, 14)
      ..lineTo(21, 14)
      ..moveTo(14, 20)
      ..lineTo(14, 21);
    canvas.drawPath(path, stroke);

    canvas.restore();
  }

  @override
  bool shouldRepaint(covariant _CelebrationPainter old) => old.color != color;
}

// ─────────────────────────────────────────────────────────────────────────────
// 17. Refresh – two curved arrows forming a circle (clockwise/anti)
// ─────────────────────────────────────────────────────────────────────────────

class _RefreshPainter extends CustomPainter {
  _RefreshPainter(this.color);
  final Color color;

  @override
  void paint(Canvas canvas, Size size) {
    final s = size.width / 24.0;
    canvas.save();
    canvas.scale(s, s);

    final stroke = _strokePaint(color);

    // Top arrow: left-to-right curving over the top
    final path = Path()
      ..moveTo(6, 8)
      ..quadraticBezierTo(12, 3, 18, 8)
      // Arrowhead
      ..moveTo(18, 8)
      ..lineTo(15, 5)
      ..moveTo(18, 8)
      ..lineTo(15, 11)

      // Bottom arrow: right-to-left curving under the bottom
      ..moveTo(18, 16)
      ..quadraticBezierTo(12, 21, 6, 16)
      // Arrowhead
      ..moveTo(6, 16)
      ..lineTo(9, 13)
      ..moveTo(6, 16)
      ..lineTo(9, 19);
    canvas.drawPath(path, stroke);

    canvas.restore();
  }

  @override
  bool shouldRepaint(covariant _RefreshPainter old) => old.color != color;
}
