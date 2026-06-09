import 'package:flutter/material.dart';
import 'package:bukadulu/config/theme.dart';
import 'package:bukadulu/routing/router.dart';

class BukaDuluApp extends StatelessWidget {
  const BukaDuluApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp.router(
      title: 'BukaDulu',
      theme: AppTheme.light,
      routerConfig: router,
      debugShowCheckedModeBanner: false,
    );
  }
}
