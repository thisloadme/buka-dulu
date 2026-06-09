import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:bukadulu/data/datasources/api.dart';
import 'package:bukadulu/domain/models/auth_response.dart';
import 'package:bukadulu/presentation/providers/token_provider.dart';

final authProvider = StateNotifierProvider<AuthNotifier, AsyncValue<AuthResponse?>>((ref) {
  return AuthNotifier(ref.read(authApiProvider), ref.read(tokenProvider.notifier));
});

class AuthNotifier extends StateNotifier<AsyncValue<AuthResponse?>> {
  final AuthApi _api;
  final StateController<String?> _tokenNotifier;

  AuthNotifier(this._api, this._tokenNotifier) : super(const AsyncValue.data(null));

  Future<void> register(String fullName, String email, String password) async {
    state = const AsyncValue.loading();
    state = await AsyncValue.guard(() async {
      final data = await _api.register(fullName: fullName, email: email, password: password);
      final resp = AuthResponse.fromJson(data);
      _tokenNotifier.state = resp.token;
      return resp;
    });
  }

  Future<void> login(String emailOrPhone, String password) async {
    state = const AsyncValue.loading();
    state = await AsyncValue.guard(() async {
      final data = await _api.login(emailOrPhone: emailOrPhone, password: password);
      final resp = AuthResponse.fromJson(data);
      _tokenNotifier.state = resp.token;
      return resp;
    });
  }

  void logout() {
    _tokenNotifier.state = null;
    state = const AsyncValue.data(null);
  }
}
