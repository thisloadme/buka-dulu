import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:bukadulu/data/datasources/api.dart';
import 'package:bukadulu/domain/models/idea.dart';

final ideaProvider = FutureProvider.family<Idea?, String>((ref, ventureId) async {
  final api = ref.read(authApiProvider);
  try {
    final data = await api.getIdea(ventureId);
    return Idea.fromJson(data);
  } catch (e) {
    return null;
  }
});
