import 'package:dio/dio.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:bukadulu/config/api_config.dart';
import 'package:bukadulu/presentation/providers/token_provider.dart';

final authApiProvider = Provider<AuthApi>((ref) {
  final dio = ref.read(dioProvider);
  final token = ref.watch(tokenProvider);
  return AuthApi(dio, token);
});

class AuthApi {
  final Dio _dio;
  final String? _token;

  AuthApi(this._dio, this._token);

  Map<String, dynamic> get _headers => {
        'Content-Type': 'application/json',
        if (_token != null) 'Authorization': 'Bearer $_token',
      };

  // ─── Auth ───
  Future<Map<String, dynamic>> register({required String fullName, required String email, required String password}) async {
    final r = await _dio.post('/auth/register', data: {'full_name': fullName, 'email': email, 'password': password});
    return r.data;
  }

  Future<Map<String, dynamic>> login({required String emailOrPhone, required String password}) async {
    final r = await _dio.post('/auth/login', data: {'email_or_phone': emailOrPhone, 'password': password});
    return r.data;
  }

  // ─── Venture ───
  Future<Map<String, dynamic>> createVenture({required String name, String? category, String? region}) async {
    final r = await _dio.post('/ventures', data: {'name': name, 'category': category, 'region': region}, options: Options(headers: _headers));
    return r.data;
  }

  Future<Map<String, dynamic>> listVentures() async {
    final r = await _dio.get('/ventures', options: Options(headers: _headers));
    return r.data;
  }

  Future<Map<String, dynamic>> getVenture(String id) async {
    final r = await _dio.get('/ventures/$id', options: Options(headers: _headers));
    return r.data;
  }

  // ─── Idea ───
  Future<Map<String, dynamic>> captureIdea(String ventureId, String rawInput) async {
    final r = await _dio.post('/ventures/$ventureId/idea', data: {'raw_input': rawInput}, options: Options(headers: _headers));
    return r.data;
  }

  Future<Map<String, dynamic>> getIdea(String ventureId) async {
    final r = await _dio.get('/ventures/$ventureId/idea', options: Options(headers: _headers));
    return r.data;
  }

  Future<Map<String, dynamic>> processIdea(String ventureId) async {
    final r = await _dio.post('/ventures/$ventureId/idea/process', options: Options(headers: _headers));
    return r.data;
  }

  Future<Map<String, dynamic>> confirmIdea(String ventureId) async {
    final r = await _dio.post('/ventures/$ventureId/idea/confirm', options: Options(headers: _headers));
    return r.data;
  }

  // ─── Customer ───
  Future<Map<String, dynamic>> createCustomer(String ventureId, {String? name, String? ageRange, String? buyContext, String? budgetRange, String? location, String? consumptionMoment}) async {
    final r = await _dio.post('/ventures/$ventureId/customer', data: {
      'name': name, 'age_range': ageRange, 'buy_context': buyContext,
      'budget_range': budgetRange, 'location': location, 'consumption_moment': consumptionMoment,
    }, options: Options(headers: _headers));
    return r.data;
  }

  Future<Map<String, dynamic>> confirmCustomer(String ventureId) async {
    final r = await _dio.post('/ventures/$ventureId/customer/confirm', options: Options(headers: _headers));
    return r.data;
  }

  // ─── Menu ───
  Future<Map<String, dynamic>> listMenus(String ventureId) async {
    final r = await _dio.get('/ventures/$ventureId/menus', options: Options(headers: _headers));
    return r.data;
  }

  Future<Map<String, dynamic>> createMenu(String ventureId, {required String name, String? description}) async {
    final r = await _dio.post('/ventures/$ventureId/menus', data: {'name': name, 'description': description}, options: Options(headers: _headers));
    return r.data;
  }

  Future<Map<String, dynamic>> updateMenu(String ventureId, String menuId, {String? status, bool? isHero}) async {
    final r = await _dio.put('/ventures/$ventureId/menus/$menuId', data: {'status': status, 'is_hero': isHero}, options: Options(headers: _headers));
    return r.data;
  }

  Future<void> deleteMenu(String ventureId, String menuId) async {
    await _dio.delete('/ventures/$ventureId/menus/$menuId', options: Options(headers: _headers));
  }

  Future<void> focusMenus(String ventureId) async {
    await _dio.post('/ventures/$ventureId/menus/focus', options: Options(headers: _headers));
  }

  // ─── Cost / Ingredients ───
  Future<Map<String, dynamic>> addIngredient(String ventureId, {required String menuId, required String name, required String unit, required double quantity, required double unitPrice}) async {
    final r = await _dio.post('/ventures/$ventureId/ingredients', data: {
      'menu_id': menuId, 'name': name, 'unit': unit, 'quantity': quantity, 'unit_price': unitPrice,
    }, options: Options(headers: _headers));
    return r.data;
  }

  Future<Map<String, dynamic>> listIngredients(String ventureId, String menuId) async {
    final r = await _dio.get('/ventures/$ventureId/ingredients?menu_id=$menuId', options: Options(headers: _headers));
    return r.data;
  }

  Future<Map<String, dynamic>> calculateCost(String ventureId, String menuId, {double laborPerUnit = 0, double overheadPerUnit = 0}) async {
    final r = await _dio.post('/ventures/$ventureId/cost/calculate/$menuId', data: {
      'labor_per_unit': laborPerUnit, 'overhead_per_unit': overheadPerUnit,
    }, options: Options(headers: _headers));
    return r.data;
  }

  Future<void> confirmCost(String ventureId) async {
    await _dio.post('/ventures/$ventureId/cost/confirm', options: Options(headers: _headers));
  }

  // ─── Sprint 3: Mission ───
  Future<Map<String, dynamic>> generateMissions(String ventureId) async {
    final r = await _dio.post('/ventures/$ventureId/missions/generate', options: Options(headers: _headers));
    return r.data;
  }

  Future<Map<String, dynamic>> listMissions(String ventureId) async {
    final r = await _dio.get('/ventures/$ventureId/missions', options: Options(headers: _headers));
    return r.data;
  }

  Future<Map<String, dynamic>> acceptMission(String ventureId, String missionId) async {
    final r = await _dio.post('/ventures/$ventureId/missions/$missionId/accept', options: Options(headers: _headers));
    return r.data;
  }

  // ─── Sprint 3: Evidence ───
  Future<Map<String, dynamic>> uploadEvidence(String ventureId, {required String missionId, required String evidenceType, String? textContent}) async {
    final r = await _dio.post('/ventures/$ventureId/evidence', data: {
      'mission_id': missionId, 'evidence_type': evidenceType, 'text_content': textContent,
    }, options: Options(headers: _headers));
    return r.data;
  }

  Future<Map<String, dynamic>> reviewEvidence(String ventureId, String evidenceId) async {
    final r = await _dio.post('/ventures/$ventureId/evidence/$evidenceId/review', options: Options(headers: _headers));
    return r.data;
  }

  // ─── Sprint 4: Scoring ───
  Future<Map<String, dynamic>> calculateScore(String ventureId) async {
    final r = await _dio.post('/ventures/$ventureId/score/calculate', options: Options(headers: _headers));
    return r.data;
  }

  Future<Map<String, dynamic>> generateDecision(String ventureId) async {
    final r = await _dio.post('/ventures/$ventureId/score/decision', options: Options(headers: _headers));
    return r.data;
  }

  Future<Map<String, dynamic>> getDecision(String ventureId) async {
    final r = await _dio.get('/ventures/$ventureId/score/decision', options: Options(headers: _headers));
    return r.data;
  }
}
