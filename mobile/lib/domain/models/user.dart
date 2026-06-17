class User {
  final String id;
  final String role;
  final String fullName;
  final String? email;
  final String? phone;
  final String status;
  final String? lastLoginAt;
  final String createdAt;
  final String updatedAt;

  User({
    required this.id,
    required this.role,
    required this.fullName,
    this.email,
    this.phone,
    required this.status,
    this.lastLoginAt,
    required this.createdAt,
    required this.updatedAt,
  });

  factory User.fromJson(Map<String, dynamic> json) => User(
        id: json['id'] as String,
        role: json['role'] as String,
        fullName: json['full_name'] as String,
        email: json['email'] as String?,
        phone: json['phone'] as String?,
        status: json['status'] as String,
        lastLoginAt: json['last_login_at'] as String?,
        createdAt: json['created_at'] as String,
        updatedAt: json['updated_at'] as String,
      );
}
