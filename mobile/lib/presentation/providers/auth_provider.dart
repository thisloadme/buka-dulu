import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:bukadulu/data/datasources/api.dart';
import 'package:bukadulu/domain/models/auth_response.dart';
import 'package:bukadulu/presentation/providers/token_provider.dart';

final authProvider = StateNotifierProvider<AuthNotifier, AsyncValue<AuthResponse?>>((ref) {
  return AuthNotifier(ref.read(authApiProvider), ref.read(tokenProvider.notifier));
});

class AuthNotifier extends StateNotifier<AsyncValue<AuthResponse?>> {
  final AuthApi _api;
  final TokenNotifier _tokenNotifier;

  // Track registration email for OTP flow
  String? _pendingEmail;

  AuthNotifier(this._api, this._tokenNotifier) : super(const AsyncValue.data(null));

  String? get pendingEmail => _pendingEmail;

  Future<void> register(String fullName, String email, String password) async {
    state = const AsyncValue.loading();
    state = await AsyncValue.guard(() async {
      final data = await _api.register(fullName: fullName, email: email, password: password);
      final resp = AuthResponse.fromJson(data);
      _pendingEmail = email;
      // Don't save token — user must verify OTP first
      return resp;
    });
  }

  Future<void> verifyOTP(String email, String otp) async {
    state = const AsyncValue.loading();
    state = await AsyncValue.guard(() async {
      final data = await _api.verifyOTP(email: email, otp: otp);
      final resp = AuthResponse.fromJson(data);
      await _tokenNotifier.save(resp.token);
      _pendingEmail = null;
      return resp;
    });
  }

  Future<void> resendOTP(String email) async {
    await _api.resendOTP(email: email);
  }

  Future<void> login(String emailOrPhone, String password) async {
    state = const AsyncValue.loading();
    state = await AsyncValue.guard(() async {
      final data = await _api.login(emailOrPhone: emailOrPhone, password: password);
      final resp = AuthResponse.fromJson(data);
      await _tokenNotifier.save(resp.token);
      return resp;
    });
  }

  Future<void> logout() async {
    await _tokenNotifier.clear();
    _pendingEmail = null;
    state = const AsyncValue.data(null);
  }
}
