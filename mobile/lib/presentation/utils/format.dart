import 'package:intl/intl.dart';

/// Format a DateTime as Indonesian relative time:
/// "Baru saja" (<1m), "5 menit lalu", "2 jam lalu", "Kemarin", "3 hari lalu",
/// "2 minggu lalu", or fallback to "d MMM yyyy".
String relativeTime(DateTime dt) {
  final now = DateTime.now();
  final diff = now.difference(dt);

  if (diff.inSeconds < 60) return 'Baru saja';
  if (diff.inMinutes < 60) return '${diff.inMinutes} menit lalu';
  if (diff.inHours < 24) return '${diff.inHours} jam lalu';
  if (diff.inDays == 1) return 'Kemarin';
  if (diff.inDays < 7) return '${diff.inDays} hari lalu';
  if (diff.inDays < 30) {
    final weeks = (diff.inDays / 7).floor();
    return '$weeks minggu lalu';
  }
  // Fallback to absolute date in id_ID locale
  return DateFormat('d MMM yyyy', 'id_ID').format(dt);
}

/// Format as "2 Jul 2026" using id_ID locale.
String formatDateShort(DateTime dt) {
  return DateFormat('d MMM yyyy', 'id_ID').format(dt);
}
