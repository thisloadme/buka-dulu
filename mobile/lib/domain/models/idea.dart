class Idea {
  final String id;
  final String ventureId;
  final String rawInput;
  final String? oneLineConcept;
  final String? targetCustomer;
  final String? valueProposition;
  final String? keyAssumptions;
  final String? earlyRisks;
  final int version;
  final bool isLocked;
  final String status;
  final String createdAt;
  final String updatedAt;

  Idea({
    required this.id,
    required this.ventureId,
    required this.rawInput,
    this.oneLineConcept,
    this.targetCustomer,
    this.valueProposition,
    this.keyAssumptions,
    this.earlyRisks,
    required this.version,
    required this.isLocked,
    required this.status,
    required this.createdAt,
    required this.updatedAt,
  });

  factory Idea.fromJson(Map<String, dynamic> json) => Idea(
        id: json['id'] as String,
        ventureId: json['venture_id'] as String,
        rawInput: json['raw_input'] as String,
        oneLineConcept: json['one_line_concept'] as String?,
        targetCustomer: json['target_customer'] as String?,
        valueProposition: json['value_proposition'] as String?,
        keyAssumptions: json['key_assumptions'] as String?,
        earlyRisks: json['early_risks'] as String?,
        version: json['version'] as int,
        isLocked: json['is_locked'] as bool,
        status: json['status'] as String,
        createdAt: json['created_at'] as String,
        updatedAt: json['updated_at'] as String,
      );
}
