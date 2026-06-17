import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:bukadulu/config/theme.dart';
import 'package:bukadulu/routing/router.dart';
import 'package:bukadulu/presentation/providers/token_provider.dart';

class BukaDuluApp extends ConsumerStatefulWidget {
  const BukaDuluApp({super.key});

  @override
  ConsumerState<BukaDuluApp> createState() => _BukaDuluAppState();
}

class _BukaDuluAppState extends ConsumerState<BukaDuluApp> {
  bool _ready = false;

  @override
  void initState() {
    super.initState();
    _init();
  }

  Future<void> _init() async {
    // Load persisted token before first frame
    await ref.read(tokenProvider.notifier).loadToken();
    if (mounted) setState(() => _ready = true);
  }

  @override
  Widget build(BuildContext context) {
    if (!_ready) {
      // Brief loading while token loads (splash will show next frame)
      return MaterialApp(
        title: 'BukaDulu',
        theme: AppTheme.light,
        debugShowCheckedModeBanner: false,
        home: const Scaffold(
          body: Center(child: CircularProgressIndicator()),
        ),
      );
    }

    return MaterialApp.router(
      title: 'BukaDulu',
      theme: AppTheme.light,
      routerConfig: router,
      debugShowCheckedModeBanner: false,
    );
  }
}
