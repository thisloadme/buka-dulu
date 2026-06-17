import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:bukadulu/data/datasources/api.dart';
import 'package:bukadulu/presentation/widgets/common/brand_icons.dart';

class MissionBoardPage extends ConsumerStatefulWidget {
  final String ventureId;
  const MissionBoardPage({super.key, required this.ventureId});
  @override
  ConsumerState<MissionBoardPage> createState() => _MissionBoardPageState();
}

class _MissionBoardPageState extends ConsumerState<MissionBoardPage> {
  List<Map<String, dynamic>> _missions = [];
  bool _loading = true;

  @override
  void initState() {
    super.initState();
    _load();
  }

  Future<void> _load() async {
    setState(() => _loading = true);
    try {
      final api = ref.read(authApiProvider);
      await api.generateMissions(widget.ventureId);
      final r = await api.listMissions(widget.ventureId);
      setState(() => _missions = List<Map<String, dynamic>>.from(r['data']));
    } catch (_) {}
    setState(() => _loading = false);
  }

  Future<void> _accept(String id) async {
    try {
      final api = ref.read(authApiProvider);
      await api.acceptMission(widget.ventureId, id);
      _load();
    } catch (e) {
      if (mounted) ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text('$e')));
    }
  }

  Color _priorityColor(String p) {
    switch (p) {
      case 'high': return Colors.red;
      case 'medium': return Colors.orange;
      default: return Colors.grey;
    }
  }

  Widget _typeIcon(String t, {Color color = Colors.grey, double size = 20}) {
    switch (t) {
      case 'polling': return BrandIcons.barChart(color: color, size: size);
      case 'interview': return BrandIcons.chatBubble(color: color, size: size);
      case 'sampling': return BrandIcons.restaurant(color: color, size: size);
      case 'observation': return BrandIcons.camera(color: color, size: size);
      default: return BrandIcons.barChart(color: color, size: size);
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Misi Harian')),
      body: _loading
          ? const Center(child: CircularProgressIndicator())
          : RefreshIndicator(
              onRefresh: _load,
              child: ListView.builder(
                padding: const EdgeInsets.all(16),
                itemCount: _missions.length + 1, // +1 for header
                itemBuilder: (_, i) {
                  if (i == 0) {
                    final done = _missions.where((m) => m['status'] == 'completed').length;
                    final total = _missions.length;
                    return Padding(
                      padding: const EdgeInsets.only(bottom: 16),
                      child: Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          Text('Hari ke-$done dari 14', style: const TextStyle(color: Color(0xFF57534e), fontSize: 13, fontWeight: FontWeight.w300)),
                          const SizedBox(height: 8),
                          LinearProgressIndicator(
                            value: total > 0 ? done / total : 0,
                            backgroundColor: const Color(0xFFe7e5e4),
                            color: const Color(0xFFea580c),
                            minHeight: 8,
                            borderRadius: BorderRadius.circular(4),
                          ),
                        ],
                      ),
                    );
                  }
                  final m = _missions[i - 1];
                  final status = m['status'] as String;
                  final priority = m['priority'] as String? ?? 'medium';
                  final type = m['mission_type'] as String? ?? 'task';

                  return Card(
                    margin: const EdgeInsets.only(bottom: 12),
                    child: Padding(
                      padding: const EdgeInsets.all(16),
                      child: Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          Row(children: [
                            Container(
                              padding: const EdgeInsets.all(6),
                              decoration: BoxDecoration(
                                color: _priorityColor(priority).withValues(alpha: 0.1),
                                borderRadius: BorderRadius.circular(8),
                              ),
                              child: _typeIcon(type, color: _priorityColor(priority), size: 20),
                            ),
                            const SizedBox(width: 8),
                            Expanded(child: Text(m['title'] ?? '', style: const TextStyle(fontWeight: FontWeight.w600))),
                            Container(
                              padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 2),
                              decoration: BoxDecoration(
                                color: status == 'completed' ? Colors.green.withValues(alpha: 0.1) :
                                       status == 'accepted' ? Colors.blue.withValues(alpha: 0.1) :
                                       status == 'in_progress' ? Colors.orange.withValues(alpha: 0.1) :
                                       Colors.grey.withValues(alpha: 0.1),
                                borderRadius: BorderRadius.circular(12),
                              ),
                              child: Text(status, style: TextStyle(
                                fontSize: 11,
                                color: status == 'completed' ? Colors.green :
                                       status == 'accepted' ? Colors.blue :
                                       status == 'in_progress' ? Colors.orange : Colors.grey,
                                fontWeight: FontWeight.w600,
                              )),
                            ),
                          ]),
                          const SizedBox(height: 8),
                          Text(m['description'] ?? '', style: const TextStyle(color: Color(0xFF57534e), fontSize: 13)),
                          const SizedBox(height: 8),
                          Row(children: [
                            if (m['estimated_minutes'] != null)
                              Row(children: [
                                BrandIcons.clock(color: const Color(0xFF57534e), size: 14),
                                const SizedBox(width: 4),
                                Text('${m['estimated_minutes']} menit', style: const TextStyle(fontSize: 12, color: Color(0xFF57534e))),
                              ]),
                            const Spacer(),
                            if (status == 'pending')
                              TextButton(onPressed: () => _accept(m['id']), child: const Text('Terima Misi')),
                            if (status == 'accepted' || status == 'evidence_submitted')
                              TextButton(
                                onPressed: () => context.go('/venture/${widget.ventureId}/mission/${m['id']}/evidence'),
                                child: const Text('Upload Bukti'),
                              ),
                            if (status == 'completed')
                              const Icon(Icons.check_circle, color: Colors.green, size: 20),
                          ]),
                        ],
                      ),
                    ),
                  );
                },
              ),
            ),
    );
  }
}
