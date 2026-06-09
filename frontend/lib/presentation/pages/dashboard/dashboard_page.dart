import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:bukadulu/presentation/providers/auth_provider.dart';
import 'package:bukadulu/presentation/providers/venture_provider.dart';
import 'package:bukadulu/domain/models/venture.dart';

class DashboardPage extends ConsumerWidget {
  const DashboardPage({super.key});

  static const _stageColors = {
    'draft': Colors.grey,
    'idea_defined': Colors.blue,
    'customer_defined': Colors.green,
    'sku_focused': Colors.orange,
    'cost_evaluated': Colors.amber,
    'mission_active': Colors.purple,
    'evidence_submitted': Colors.pink,
    'evidence_reviewed': Colors.teal,
    'ready_to_decide': Colors.indigo,
    'continue': Colors.green,
    'repeat': Colors.orange,
    'pivot': Colors.amber,
    'stop': Colors.red,
  };

  static const _stageLabels = {
    'draft': 'Draft',
    'idea_defined': 'Ide',
    'customer_defined': 'Pelanggan',
    'sku_focused': 'Menu',
    'cost_evaluated': 'Biaya',
    'mission_active': 'Misi',
    'evidence_submitted': 'Bukti',
    'evidence_reviewed': 'Review',
    'ready_to_decide': 'Siap',
    'continue': 'Lanjut',
    'repeat': 'Ulang',
    'pivot': 'Pivot',
    'stop': 'Stop',
  };

  @override
  Widget build(BuildContext context, WidgetRef ref) {
    final authState = ref.watch(authProvider);
    final venturesAsync = ref.watch(ventureListProvider);

    return Scaffold(
      appBar: AppBar(
        title: const Text('BukaDulu'),
        actions: [
          if (authState.valueOrNull != null)
            IconButton(
              icon: const Icon(Icons.logout),
              onPressed: () {
                ref.read(authProvider.notifier).logout();
                context.go('/login');
              },
            ),
        ],
      ),
      floatingActionButton: FloatingActionButton.extended(
        onPressed: () => context.go('/venture/new'),
        icon: const Icon(Icons.add),
        label: const Text('Ide Baru'),
      ),
      body: venturesAsync.when(
        data: (ventures) {
          if (ventures.isEmpty) {
            return Center(
              child: Padding(
                padding: const EdgeInsets.all(32),
                child: Column(
                  mainAxisSize: MainAxisSize.min,
                  children: [
                    Icon(Icons.lightbulb_outline, size: 80, color: Colors.grey[300]),
                    const SizedBox(height: 24),
                    Text('Belum ada ide yang divalidasi', style: Theme.of(context).textTheme.titleMedium?.copyWith(color: Colors.grey[600])),
                    const SizedBox(height: 8),
                    Text('Mulai dengan menekan tombol "Ide Baru"', style: TextStyle(color: Colors.grey[500])),
                    const SizedBox(height: 24),
                    ElevatedButton.icon(
                      onPressed: () => context.go('/venture/new'),
                      icon: const Icon(Icons.add),
                      label: const Text('Mulai Validasi'),
                    ),
                  ],
                ),
              ),
            );
          }
          return RefreshIndicator(
            onRefresh: () => ref.refresh(ventureListProvider.future),
            child: ListView.builder(
              padding: const EdgeInsets.fromLTRB(16, 16, 16, 96),
              itemCount: ventures.length,
              itemBuilder: (context, index) => _VentureCard(venture: ventures[index]),
            ),
          );
        },
        loading: () => const Center(child: CircularProgressIndicator()),
        error: (err, _) => Center(child: Text('Error: $err')),
      ),
    );
  }
}

class _VentureCard extends StatelessWidget {
  final Venture venture;
  const _VentureCard({required this.venture});

  @override
  Widget build(BuildContext context) {
    final stageColor = DashboardPage._stageColors[venture.stage] ?? Colors.grey;
    final stageLabel = DashboardPage._stageLabels[venture.stage] ?? venture.stage;

    return Card(
      margin: const EdgeInsets.only(bottom: 12),
      child: InkWell(
        onTap: () => context.go('/venture/${venture.id}/idea'),
        borderRadius: BorderRadius.circular(12),
        child: Padding(
          padding: const EdgeInsets.all(16),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Row(
                children: [
                  Container(
                    padding: const EdgeInsets.symmetric(horizontal: 8, vertical: 4),
                    decoration: BoxDecoration(
                      color: stageColor.withValues(alpha: 0.1),
                      borderRadius: BorderRadius.circular(12),
                    ),
                    child: Text(stageLabel, style: TextStyle(color: stageColor, fontSize: 12, fontWeight: FontWeight.w600)),
                  ),
                  const Spacer(),
                  Text(venture.createdAt.substring(0, 10), style: TextStyle(color: Colors.grey[500], fontSize: 12)),
                ],
              ),
              const SizedBox(height: 12),
              Text(venture.name, style: Theme.of(context).textTheme.titleMedium?.copyWith(fontWeight: FontWeight.w600)),
              if (venture.category != null && venture.category!.isNotEmpty) ...[
                const SizedBox(height: 4),
                Text(venture.category!, style: TextStyle(color: Colors.grey[600])),
              ],
              const SizedBox(height: 12),
              LinearProgressIndicator(
                value: _stageProgress(venture.stage),
                backgroundColor: Colors.grey[200],
                color: stageColor,
              ),
            ],
          ),
        ),
      ),
    );
  }

  double _stageProgress(String stage) {
    const stages = ['draft', 'idea_defined', 'customer_defined', 'sku_focused', 'cost_evaluated', 'mission_active', 'evidence_submitted', 'evidence_reviewed', 'ready_to_decide'];
    final idx = stages.indexOf(stage);
    return idx < 0 ? 0.0 : (idx + 1) / stages.length;
  }
}
