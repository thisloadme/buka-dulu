import 'package:bukadulu/domain/models/user.dart';

class AuthResponse {
  final User user;
  final String token;
  final String expiresAt;

  AuthResponse({
    required this.user,
    required this.token,
    required this.expiresAt,
  });

  factory AuthResponse.fromJson(Map<String, dynamic> json) => AuthResponse(
        user: User.fromJson(json['user']),
        token: json['token'] as String,
        expiresAt: json['expires_at'] as String,
      );
}
