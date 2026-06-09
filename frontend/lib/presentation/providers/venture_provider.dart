import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:bukadulu/data/datasources/api.dart';
import 'package:bukadulu/domain/models/venture.dart';

final ventureListProvider = FutureProvider<List<Venture>>((ref) async {
  final api = ref.read(authApiProvider);
  final data = await api.listVentures();
  final list = (data['data'] as List).map((e) => Venture.fromJson(e)).toList();
  return list;
});

final ventureDetailProvider = FutureProvider.family<Venture, String>((ref, id) async {
  final api = ref.read(authApiProvider);
  final data = await api.getVenture(id);
  return Venture.fromJson(data);
});
