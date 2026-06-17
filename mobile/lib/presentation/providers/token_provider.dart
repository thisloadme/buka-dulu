import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

const _tokenKey = 'bukadulu_auth_token';

final tokenProvider = StateNotifierProvider<TokenNotifier, String?>((ref) {
  return TokenNotifier();
});

class TokenNotifier extends StateNotifier<String?> {
  TokenNotifier() : super(null);

  final _storage = const FlutterSecureStorage();

  /// Load token from secure storage (call at app start).
  Future<void> loadToken() async {
    final token = await _storage.read(key: _tokenKey);
    if (token != null && token.isNotEmpty) {
      state = token;
    }
  }

  /// Save token to secure storage and update state.
  Future<void> save(String token) async {
    await _storage.write(key: _tokenKey, value: token);
    state = token;
  }

  /// Clear token from secure storage and update state.
  Future<void> clear() async {
    await _storage.delete(key: _tokenKey);
    state = null;
  }
}
