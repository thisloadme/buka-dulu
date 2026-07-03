import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:dio/dio.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:bukadulu/presentation/providers/token_provider.dart';

final dioProvider = Provider<Dio>((ref) {
  final dio = Dio(BaseOptions(
    baseUrl: dotenv.env['API_BASE_URL'] ?? 'https://api-buka-dulu.riyantobudi.biz.id/api/v1',
    connectTimeout: const Duration(seconds: 10),
    receiveTimeout: const Duration(seconds: 30),
    headers: {'Content-Type': 'application/json'},
  ));

  // Auth interceptor: read token dynamically at request time.
  // This avoids race conditions where AuthApi captures token by value
  // and FutureProviders see a stale token right after verifyOTP.
  dio.interceptors.add(InterceptorsWrapper(
    onRequest: (options, handler) {
      // Skip auth header for public endpoints
      final path = options.path;
      final isPublic = path.contains('/auth/');
      if (!isPublic) {
        final token = ref.read(tokenProvider);
        if (token != null && token.isNotEmpty) {
          options.headers['Authorization'] = 'Bearer $token';
        }
      }
      handler.next(options);
    },
    onError: (err, handler) {
      if (kDebugMode) {
        debugPrint('[Dio] ${err.requestOptions.method} ${err.requestOptions.path} → ${err.response?.statusCode}: ${err.message}');
      }
      handler.next(err);
    },
  ));

  return dio;
});
