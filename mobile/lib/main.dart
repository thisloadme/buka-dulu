import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'app.dart';
import 'presentation/providers/payment_provider.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await dotenv.load(fileName: '.env');

  // Start listening for Play Billing purchases (no-op in debug/dev)
  final billing = PlayBillingService();
  billing.startListening();

  runApp(const ProviderScope(child: BukaDuluApp()));
}
