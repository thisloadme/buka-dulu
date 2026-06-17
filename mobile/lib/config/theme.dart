import 'package:flutter/material.dart';
import 'package:google_fonts/google_fonts.dart';

class AppTheme {
  static ThemeData get light {
    final textTheme = GoogleFonts.interTextTheme();
    const orange = Color(0xFFea580c);
    const label = Color(0xFF292524);
    const border = Color(0xFFe7e5e4);
    const white = Colors.white;

    return ThemeData(
      useMaterial3: true,
      colorScheme: ColorScheme.fromSeed(
        seedColor: orange,
        brightness: Brightness.light,
      ),
      textTheme: textTheme,
      appBarTheme: AppBarTheme(
        centerTitle: true,
        elevation: 0,
        titleTextStyle: textTheme.titleLarge?.copyWith(
          fontWeight: FontWeight.w400,
        ),
      ),
      elevatedButtonTheme: ElevatedButtonThemeData(
        style: ElevatedButton.styleFrom(
          backgroundColor: orange,
          foregroundColor: white,
          elevation: 0,
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(6),
          ),
          disabledBackgroundColor: Colors.grey,
          disabledForegroundColor: white,
          padding: const EdgeInsets.symmetric(horizontal: 22, vertical: 10),
          textStyle: const TextStyle(
            fontSize: 16,
            fontWeight: FontWeight.w600,
          ),
        ),
      ),
      outlinedButtonTheme: OutlinedButtonThemeData(
        style: OutlinedButton.styleFrom(
          backgroundColor: Colors.transparent,
          foregroundColor: orange,
          side: const BorderSide(color: Color(0xFFfed7aa)),
          shape: RoundedRectangleBorder(
            borderRadius: BorderRadius.circular(6),
          ),
          padding: const EdgeInsets.symmetric(horizontal: 22, vertical: 10),
          textStyle: const TextStyle(
            fontSize: 16,
            fontWeight: FontWeight.w600,
          ),
        ),
      ),
      textButtonTheme: TextButtonThemeData(
        style: TextButton.styleFrom(
          foregroundColor: orange,
        ),
      ),
      inputDecorationTheme: InputDecorationTheme(
        border: OutlineInputBorder(
          borderRadius: BorderRadius.circular(8),
          borderSide: const BorderSide(color: border),
        ),
        enabledBorder: OutlineInputBorder(
          borderRadius: BorderRadius.circular(8),
          borderSide: const BorderSide(color: border),
        ),
        focusedBorder: OutlineInputBorder(
          borderRadius: BorderRadius.circular(8),
          borderSide: const BorderSide(color: orange, width: 2),
        ),
        errorBorder: OutlineInputBorder(
          borderRadius: BorderRadius.circular(8),
          borderSide: const BorderSide(color: Color(0xFFfecaca)),
        ),
        focusedErrorBorder: OutlineInputBorder(
          borderRadius: BorderRadius.circular(8),
          borderSide: const BorderSide(color: Color(0xFFdc2626), width: 2),
        ),
        contentPadding: const EdgeInsets.symmetric(horizontal: 16, vertical: 12),
        hintStyle: const TextStyle(
          color: Color(0xFFa8a29e),
          fontWeight: FontWeight.w300,
        ),
        labelStyle: const TextStyle(
          fontSize: 14, // 0.875rem
          fontWeight: FontWeight.w500,
          color: label,
        ),
        floatingLabelBehavior: FloatingLabelBehavior.auto,
      ),
      cardTheme: CardThemeData(
        elevation: 0,
        color: white,
        surfaceTintColor: Colors.transparent,
        shape: RoundedRectangleBorder(
          borderRadius: BorderRadius.circular(8),
          side: const BorderSide(color: border),
        ),
        shadowColor: Colors.transparent,
      ),
    );
  }

  /// Custom card shadow from DESIGN.md:
  ///   rgba(80,50,20,0.12) 0px 50px 80px -40px
  ///   rgba(0,0,0,0.06) 0px 20px 40px -20px
  static List<BoxShadow> get cardShadow => const [
        BoxShadow(
          color: Color(0x1F503214), // rgba(80,50,20,0.12)
          offset: Offset(0, 50),
          blurRadius: 80,
          spreadRadius: -40,
        ),
        BoxShadow(
          color: Color(0x0F000000), // rgba(0,0,0,0.06)
          offset: Offset(0, 20),
          blurRadius: 40,
          spreadRadius: -20,
        ),
      ];
}
