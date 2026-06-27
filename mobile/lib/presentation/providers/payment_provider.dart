import 'dart:async';
import 'package:flutter/foundation.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:in_app_purchase/in_app_purchase.dart';
import 'package:in_app_purchase_android/in_app_purchase_android.dart';

/// Product IDs for subscription plans (defined in Google Play Console).
/// Free plan doesn't need a product ID.
const kProductSprint = 'bukadulu_sprint_monthly';
const kProductPro = 'bukadulu_pro_monthly';

/// Provider for Google Play Billing IAP service.
final playBillingProvider = Provider<PlayBillingService>((ref) {
  return PlayBillingService();
});

class PlayBillingService {
  final InAppPurchase _iap = InAppPurchase.instance;
  StreamSubscription<List<PurchaseDetails>>? _subscription;

  /// Whether Play Billing is available.
  /// Disabled in debug/dev — app must be signed & on Google Play.
  bool get isAvailable => kReleaseMode;

  /// Listen for purchase updates (must be called early in app lifecycle).
  void startListening() {
    if (!isAvailable) return;
    _subscription = _iap.purchaseStream.listen(_handlePurchaseUpdate);
  }

  void stopListening() {
    _subscription?.cancel();
    _subscription = null;
  }

  void _handlePurchaseUpdate(List<PurchaseDetails> purchases) {
    for (final purchase in purchases) {
      if (purchase.status == PurchaseStatus.purchased ||
          purchase.status == PurchaseStatus.restored) {
        // Verify purchase server-side & grant entitlement
        _verifyAndGrant(purchase);
      } else if (purchase.status == PurchaseStatus.error) {
        debugPrint('Play Billing error: ${purchase.error}');
      }
    }
  }

  /// Fetch subscription product details from Google Play.
  Future<List<ProductDetails>> getProducts() async {
    if (!isAvailable) return [];

    final ids = {kProductSprint, kProductPro};
    final resp = await _iap.queryProductDetails(ids);

    if (resp.error != null) {
      debugPrint('Product query error: ${resp.error}');
      return [];
    }
    return resp.productDetails;
  }

  /// Launch billing flow for a given product.
  Future<bool> purchaseProduct(ProductDetails product) async {
    if (!isAvailable) return false;

    final purchaseParam = PurchaseParam(
      productDetails: product,
    );

    try {
      await _iap.buyNonConsumable(purchaseParam: purchaseParam);
      return true; // Stream will handle the result
    } catch (e) {
      debugPrint('Purchase failed: $e');
      return false;
    }
  }

  /// Verify purchase server-side (placeholder — call backend API).
  Future<void> _verifyAndGrant(PurchaseDetails purchase) async {
    // TODO: Send purchase.verificationData to backend for verification
    // POST /api/v1/subscriptions/verify with purchase token
    // On success: completePurchase(purchase)
    if (purchase.pendingCompletePurchase) {
      await _iap.completePurchase(purchase);
    }
  }

  /// Get active subscription product IDs (from past verified purchases).
  Future<List<String>> getActiveSubscriptions() async {
    if (!isAvailable) return [];

    final purchases = await _iap.queryPastPurchases();
    return purchases
        .where((p) => p.status == PurchaseStatus.purchased)
        .map((p) => p.productID)
        .toList();
  }

  void dispose() {
    stopListening();
  }
}
