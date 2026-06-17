import 'package:flutter/material.dart';

// =============================================================================
// StageColors — static color mapping from DESIGN.md
// =============================================================================

/// Color values for every venture stage, drawn from DESIGN.md.
///
/// Usage:
/// ```dart
/// final color = StageColors.forStage('mission_active'); // Color(0xFF3b82f6)
/// ```
class StageColors {
  StageColors._();

  // Progression stages (9-stage pipeline)
  static const Color draft = Color(0xFF78716c);
  static const Color ideaDefined = Color(0xFFea580c);
  static const Color customerDefined = Color(0xFFea580c);
  static const Color skuFocused = Color(0xFFea580c);
  static const Color costEvaluated = Color(0xFFea580c);
  static const Color missionActive = Color(0xFF3b82f6);
  static const Color evidenceSubmitted = Color(0xFFf59e0b);
  static const Color evidenceReviewed = Color(0xFF22c55e);
  static const Color readyToDecide = Color(0xFF8b5cf6);

  // Outcome stages (end states)
  static const Color continue_ = Color(0xFF22c55e);
  static const Color repeat = Color(0xFFf59e0b);
  static const Color pivot = Color(0xFFf59e0b);
  static const Color stop = Color(0xFFef4444);

  /// Look up the colour for a [stage] key.
  ///
  /// Returns [draft] (gray) for unknown stages so the UI never breaks.
  static Color forStage(String? stage) {
    if (stage == null) return draft;
    switch (stage) {
      case 'draft':
        return draft;
      case 'idea_defined':
        return ideaDefined;
      case 'customer_defined':
        return customerDefined;
      case 'sku_focused':
        return skuFocused;
      case 'cost_evaluated':
        return costEvaluated;
      case 'mission_active':
        return missionActive;
      case 'evidence_submitted':
        return evidenceSubmitted;
      case 'evidence_reviewed':
        return evidenceReviewed;
      case 'ready_to_decide':
        return readyToDecide;
      case 'continue':
        return continue_;
      case 'repeat':
        return repeat;
      case 'pivot':
        return pivot;
      case 'stop':
        return stop;
      default:
        return draft;
    }
  }
}

// =============================================================================
// StageLabels — human-readable (Indonesian) labels
// =============================================================================

/// Indonesian labels for every venture stage.
///
/// Usage:
/// ```dart
/// final label = StageLabels.forStage('mission_active'); // "Misi"
/// ```
class StageLabels {
  StageLabels._();

  static const String draft = 'Draft';
  static const String ideaDefined = 'Ide';
  static const String customerDefined = 'Target';
  static const String skuFocused = 'Menu';
  static const String costEvaluated = 'Biaya';
  static const String missionActive = 'Misi';
  static const String evidenceSubmitted = 'Bukti';
  static const String evidenceReviewed = 'Review';
  static const String readyToDecide = 'Siap';
  static const String continue_ = 'Lanjut';
  static const String repeat = 'Ulang';
  static const String pivot = 'Pivot';
  static const String stop = 'Stop';

  /// Look up the label for a [stage] key.
  ///
  /// Returns `'Draft'` for unknown stages.
  static String forStage(String? stage) {
    if (stage == null) return draft;
    switch (stage) {
      case 'draft':
        return draft;
      case 'idea_defined':
        return ideaDefined;
      case 'customer_defined':
        return customerDefined;
      case 'sku_focused':
        return skuFocused;
      case 'cost_evaluated':
        return costEvaluated;
      case 'mission_active':
        return missionActive;
      case 'evidence_submitted':
        return evidenceSubmitted;
      case 'evidence_reviewed':
        return evidenceReviewed;
      case 'ready_to_decide':
        return readyToDecide;
      case 'continue':
        return continue_;
      case 'repeat':
        return repeat;
      case 'pivot':
        return pivot;
      case 'stop':
        return stop;
      default:
        return draft;
    }
  }
}

// =============================================================================
// Stage Progress — ordered pipeline of 9 stages
// =============================================================================

/// Ordered list of progression stages (excludes outcome stages).
const List<String> kStageProgression = [
  'draft',
  'idea_defined',
  'customer_defined',
  'sku_focused',
  'cost_evaluated',
  'mission_active',
  'evidence_submitted',
  'evidence_reviewed',
  'ready_to_decide',
];

/// Returns the 0-based index of [stage] within [kStageProgression],
/// or -1 if the stage is an outcome or unknown.
int stageIndex(String? stage) {
  if (stage == null) return -1;
  return kStageProgression.indexOf(stage);
}

/// Returns the normalised progress (0.0–1.0) for a progression stage.
double stageProgress(String? stage) {
  final idx = stageIndex(stage);
  if (idx < 0) return 0.0;
  return (idx + 1) / kStageProgression.length;
}

// =============================================================================
// StageBadge
// =============================================================================

/// A small coloured pill that displays a venture stage label.
///
/// Rendered with:
/// - Colored background at 0.1 alpha
/// - Colored text at full opacity
/// - Border radius of 12px
/// - Small font (11–12px)
class StageBadge extends StatelessWidget {
  const StageBadge({super.key, required this.stage});

  /// The stage key (e.g. `'mission_active'`).
  final String? stage;

  @override
  Widget build(BuildContext context) {
    final color = StageColors.forStage(stage);
    final label = StageLabels.forStage(stage);

    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 10, vertical: 3),
      decoration: BoxDecoration(
        color: color.withValues(alpha: 0.1),
        borderRadius: BorderRadius.circular(12),
      ),
      child: Text(
        label,
        style: TextStyle(
          color: color,
          fontSize: 11.5,
          fontWeight: FontWeight.w600,
          letterSpacing: 0.3,
        ),
      ),
    );
  }
}

// =============================================================================
// StageProgressBar
// =============================================================================

/// A thin horizontal progress bar that reflects how far a venture has
/// progressed through the 9-stage pipeline.
///
/// The bar is coloured according to the current stage.  Outcomes (continue,
/// repeat, pivot, stop) render a full bar.
class StageProgressBar extends StatelessWidget {
  const StageProgressBar({
    super.key,
    required this.stage,
    this.height = 3,
  });

  /// The current stage key.
  final String? stage;

  /// Height of the bar.  Defaults to 3px.
  final double height;

  @override
  Widget build(BuildContext context) {
    final color = StageColors.forStage(stage);
    final progress = stageProgress(stage);
    // Clamp to 0.0–1.0; outcome stages get full bar via stageProgress.
    final clampedProgress = progress.clamp(0.0, 1.0);

    return SizedBox(
      height: height,
      child: Row(
        children: [
          Expanded(
            child: ClipRRect(
              borderRadius: BorderRadius.circular(height / 2),
              child: LinearProgressIndicator(
                value: clampedProgress,
                backgroundColor: const Color(0xFFe7e5e4),
                color: color,
                minHeight: height,
              ),
            ),
          ),
        ],
      ),
    );
  }
}

// =============================================================================
// AppShell
// =============================================================================

/// A Scaffold-like wrapper for venture flow screens.
///
/// Provides:
/// - Optional [AppBar] with a back button (when [showBack] is true) and [title]
/// - Optional [StageBadge] pill rendered beside the title or in the app bar
/// - Optional [StageProgressBar] at the top of the body area
/// - Optional sticky bottom CTA button (when [bottomCta] is provided)
/// - [body] content in the centre
///
/// Usage:
/// ```dart
/// AppShell(
///   title: 'Biaya & Margin',
///   stage: 'cost_evaluated',
///   bottomCta: 'Konfirmasi Biaya',
///   onBottomCta: () => context.read<VentureBloc>().add(ConfirmCost()),
///   body: CostForm(),
/// )
/// ```
class AppShell extends StatelessWidget {
  const AppShell({
    super.key,
    this.title,
    this.stage,
    this.showBack = true,
    this.bottomCta,
    this.onBottomCta,
    this.bottomDisabled = false,
    this.body = const SizedBox.shrink(),
  });

  /// Optional title displayed in the app bar.
  final String? title;

  /// Optional stage key for the badge pill and progress bar.
  ///
  /// When provided, both [StageBadge] and [StageProgressBar] are shown.
  final String? stage;

  /// Whether to show a back arrow in the app bar.  Defaults to `true`.
  final bool showBack;

  /// Optional label for a sticky bottom CTA button.
  ///
  /// When `null`, no bottom button is rendered.
  final String? bottomCta;

  /// Callback invoked when the bottom CTA is pressed.
  ///
  /// Ignored when [bottomCta] is `null`.
  final VoidCallback? onBottomCta;

  /// Whether the bottom CTA button should appear disabled.  Defaults to `false`.
  final bool bottomDisabled;

  /// The primary content widget placed in the centre of the shell.
  final Widget body;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: _buildAppBar(context),
      body: _buildBody(context),
      bottomNavigationBar: _buildBottomCta(context),
    );
  }

  PreferredSizeWidget? _buildAppBar(BuildContext context) {
    // If there's no title and no stage, omit the app bar entirely.
    if (title == null && stage == null) return null;

    return AppBar(
      leading: showBack
          ? IconButton(
              icon: const Icon(Icons.arrow_back_ios_rounded, size: 20),
              onPressed: () => Navigator.of(context).maybePop(),
            )
          : null,
      title: title != null
          ? Text(
              title!,
              style: const TextStyle(
                fontWeight: FontWeight.w400,
                fontSize: 16,
              ),
            )
          : null,
      actions: [
        if (stage != null)
          Padding(
            padding: const EdgeInsets.only(right: 16),
            child: StageBadge(stage: stage),
          ),
      ],
    );
  }

  Widget _buildBody(BuildContext context) {
    return Column(
      children: [
        // Progress bar at top of body (beneath app bar)
        if (stage != null)
          Padding(
            padding: const EdgeInsets.symmetric(horizontal: 16),
            child: StageProgressBar(stage: stage),
          ),

        // Main content — expanded to fill remaining space
        Expanded(child: body),
      ],
    );
  }

  Widget? _buildBottomCta(BuildContext context) {
    if (bottomCta == null) return null;

    return SafeArea(
      child: Padding(
        padding: const EdgeInsets.fromLTRB(16, 8, 16, 8),
        child: SizedBox(
          width: double.infinity,
          child: ElevatedButton(
            onPressed: bottomDisabled ? null : onBottomCta,
            style: ElevatedButton.styleFrom(
              backgroundColor: const Color(0xFFea580c),
              foregroundColor: Colors.white,
              disabledBackgroundColor: const Color(0xFFe7e5e4),
              disabledForegroundColor: const Color(0xFFa8a29e),
              elevation: 0,
              shape: RoundedRectangleBorder(
                borderRadius: BorderRadius.circular(6),
              ),
              padding: const EdgeInsets.symmetric(vertical: 14),
              textStyle: const TextStyle(
                fontSize: 16,
                fontWeight: FontWeight.w600,
              ),
            ),
            child: Text(bottomCta!),
          ),
        ),
      ),
    );
  }
}
