import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:in_app_purchase/in_app_purchase.dart';
import 'package:bukadulu/domain/models/subscription_plan.dart';
import 'package:bukadulu/presentation/providers/payment_provider.dart';

class SubscriptionPage extends ConsumerStatefulWidget {
  const SubscriptionPage({super.key});

  @override
  ConsumerState<SubscriptionPage> createState() => _SubscriptionPageState();
}

class _SubscriptionPageState extends ConsumerState<SubscriptionPage> {
  List<ProductDetails>? _products;
  bool _loading = true;

  @override
  void initState() {
    super.initState();
    _loadProducts();
  }

  Future<void> _loadProducts() async {
    final billing = ref.read(playBillingProvider);
    if (!billing.isAvailable) {
      setState(() => _loading = false);
      return;
    }
    final products = await billing.getProducts();
    if (mounted) setState(() { _products = products; _loading = false; });
  }

  ProductDetails? _findProduct(SubscriptionPlan plan) {
    if (_products == null) return null;
    final id = plan.id == 'plan_sprint'
        ? kProductSprint
        : plan.id == 'plan_pro'
            ? kProductPro
            : null;
    if (id == null) return null;
    try {
      return _products!.firstWhere((p) => p.id == id);
    } catch (_) {
      return null;
    }
  }

  Future<void> _purchase(SubscriptionPlan plan) async {
    if (plan.priceCents == 0) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('Paket Free aktif!')),
        );
      }
      return;
    }

    // Dev mode: show warning banner
    if (kDebugMode || !kReleaseMode) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('Pembayaran nonaktif — Google Play Billing hanya aktif di release build')),
        );
      }
      return;
    }

    final product = _findProduct(plan);
    if (product == null) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('Produk tidak ditemukan di Google Play')),
        );
      }
      return;
    }

    final billing = ref.read(playBillingProvider);
    final ok = await billing.purchaseProduct(product);
    if (ok && mounted) {
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('Pembelian diproses...')),
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    final billing = ref.watch(playBillingProvider);
    final canPurchase = billing.isAvailable;

    return Scaffold(
      appBar: AppBar(title: const Text('Langganan')),
      body: SafeArea(
        child: ListView(
          padding: const EdgeInsets.all(16),
          children: [
            Text('Pilih Paket', style: Theme.of(context).textTheme.headlineSmall),
            const SizedBox(height: 8),
            Text(
              'Akses fitur lengkap untuk validasi bisnis F&B kamu',
              style: Theme.of(context).textTheme.bodyMedium?.copyWith(color: const Color(0xFF57534e)),
            ),
            const SizedBox(height: 24),

            // Dev mode notice
            if (kDebugMode || !kReleaseMode)
              Container(
                padding: const EdgeInsets.all(12),
                margin: const EdgeInsets.only(bottom: 16),
                decoration: BoxDecoration(
                  color: const Color(0xFFfef3c7),
                  borderRadius: BorderRadius.circular(8),
                  border: Border.all(color: const Color(0xFFf59e0b)),
                ),
                child: const Row(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Icon(Icons.build, color: Color(0xFF92400e), size: 20),
                    SizedBox(width: 8),
                    Expanded(
                      child: Text(
                        'Mode Development — Google Play Billing dinonaktifkan. '
                        'Hanya aktif di release build dengan aplikasi terdaftar di Google Play Console.',
                        style: TextStyle(color: Color(0xFF92400e), fontSize: 13),
                      ),
                    ),
                  ],
                ),
              ),

            if (_loading)
              const Center(child: CircularProgressIndicator())
            else
              ...SubscriptionPlan.availablePlans.map((plan) => _PlanCard(
                    plan: plan,
                    canPurchase: canPurchase && plan.priceCents > 0,
                    priceFromPlay: _findProduct(plan)?.price,
                    onSubscribe: () => _purchase(plan),
                  )),
          ],
        ),
      ),
    );
  }
}

class _PlanCard extends StatelessWidget {
  final SubscriptionPlan plan;
  final bool canPurchase;
  final String? priceFromPlay;
  final VoidCallback onSubscribe;

  const _PlanCard({
    required this.plan,
    required this.canPurchase,
    this.priceFromPlay,
    required this.onSubscribe,
  });

  @override
  Widget build(BuildContext context) {
    final isFree = plan.priceCents == 0;

    return Container(
      margin: const EdgeInsets.only(bottom: 12),
      decoration: BoxDecoration(
        color: plan.isPopular ? const Color(0xFFfff7ed) : Colors.white,
        borderRadius: BorderRadius.circular(12),
        border: Border.all(
          color: plan.isPopular ? const Color(0xFFea580c) : const Color(0xFFe7e5e4),
          width: plan.isPopular ? 2 : 1,
        ),
      ),
      child: Stack(
        children: [
          if (plan.isPopular)
            Positioned(
              top: 8,
              right: 8,
              child: Container(
                padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 4),
                decoration: BoxDecoration(
                  color: const Color(0xFFea580c),
                  borderRadius: BorderRadius.circular(4),
                ),
                child: const Text(
                  'POPULER',
                  style: TextStyle(color: Colors.white, fontSize: 10, fontWeight: FontWeight.w500),
                ),
              ),
            ),
          Padding(
            padding: const EdgeInsets.all(16),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Text(plan.name, style: Theme.of(context).textTheme.titleLarge),
                const SizedBox(height: 4),
                Text(plan.description,
                    style: const TextStyle(color: Color(0xFF57534e), fontSize: 13)),
                const SizedBox(height: 12),
                Row(
                  crossAxisAlignment: CrossAxisAlignment.end,
                  children: [
                    Text(
                      priceFromPlay ?? (isFree ? 'Gratis' : plan.priceFormatted),
                      style: Theme.of(context).textTheme.headlineMedium?.copyWith(
                            fontWeight: FontWeight.w300,
                            color: const Color(0xFFea580c),
                          ),
                    ),
                    if (!isFree)
                      const Padding(
                        padding: EdgeInsets.only(bottom: 6, left: 4),
                        child: Text('/bulan', style: TextStyle(color: Color(0xFFa8a29e), fontSize: 13)),
                      ),
                  ],
                ),
                const SizedBox(height: 12),
                ...plan.features.map((f) => Padding(
                      padding: const EdgeInsets.only(bottom: 6),
                      child: Row(
                        children: [
                          const Icon(Icons.check, size: 16, color: Color(0xFF22c55e)),
                          const SizedBox(width: 8),
                          Text(f, style: const TextStyle(fontSize: 14)),
                        ],
                      ),
                    )),
                const SizedBox(height: 16),
                SizedBox(
                  width: double.infinity,
                  child: ElevatedButton(
                    onPressed: onSubscribe,
                    style: plan.isPopular
                        ? ElevatedButton.styleFrom(
                            backgroundColor: const Color(0xFFea580c),
                            foregroundColor: Colors.white,
                          )
                        : null,
                    child: Text(isFree ? 'Mulai Gratis' : 'Langganan'),
                  ),
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}
