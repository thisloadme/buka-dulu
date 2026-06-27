class SubscriptionPlan {
  final String id;
  final String name;
  final String description;
  final int priceCents; // in cents (IDR)
  final String priceLabel;
  final List<String> features;
  final bool isPopular;
  final String duration; // "monthly" | "quarterly" | "yearly"

  const SubscriptionPlan({
    required this.id,
    required this.name,
    required this.description,
    required this.priceCents,
    required this.priceLabel,
    required this.features,
    this.isPopular = false,
    required this.duration,
  });

  String get priceFormatted => 'Rp ${(priceCents / 100).toStringAsFixed(0)}';

  static const List<SubscriptionPlan> availablePlans = [
    SubscriptionPlan(
      id: 'plan_free',
      name: 'Free',
      description: 'Coba fitur dasar',
      priceCents: 0,
      priceLabel: 'Gratis',
      features: [
        '1 venture',
        'Idea compression',
        '3 misi pertama',
      ],
      duration: 'monthly',
    ),
    SubscriptionPlan(
      id: 'plan_sprint',
      name: 'Sprint',
      description: '14-day validation program',
      priceCents: 99000,
      priceLabel: 'Rp 99.000',
      isPopular: true,
      features: [
        'Unlimited ventures',
        'Full AI features',
        'Unlimited missions',
        'Evidence review',
        'Score & decision',
        'Mentor access',
      ],
      duration: 'monthly',
    ),
    SubscriptionPlan(
      id: 'plan_pro',
      name: 'Pro',
      description: 'Untuk bisnis serius',
      priceCents: 249000,
      priceLabel: 'Rp 249.000',
      features: [
        'Semua fitur Sprint',
        'Priority AI review',
        'Multi-founder collab',
        'Export data',
        'Priority support',
      ],
      duration: 'monthly',
    ),
  ];
}
