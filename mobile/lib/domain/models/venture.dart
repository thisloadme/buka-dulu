class Venture {
  final String id;
  final String ownerUserId;
  final String name;
  final String? category;
  final String? region;
  final String stage;
  final int currentVersion;
  final double? score;
  final String createdAt;
  final String updatedAt;

  Venture({
    required this.id,
    required this.ownerUserId,
    required this.name,
    this.category,
    this.region,
    required this.stage,
    required this.currentVersion,
    this.score,
    required this.createdAt,
    required this.updatedAt,
  });

  factory Venture.fromJson(Map<String, dynamic> json) => Venture(
        id: json['id'] as String,
        ownerUserId: json['owner_user_id'] as String,
        name: json['name'] as String,
        category: json['category'] as String?,
        region: json['region'] as String?,
        stage: json['stage'] as String,
        currentVersion: json['current_version'] as int,
        score: (json['score'] as num?)?.toDouble(),
        createdAt: json['created_at'] as String,
        updatedAt: json['updated_at'] as String,
      );
}
