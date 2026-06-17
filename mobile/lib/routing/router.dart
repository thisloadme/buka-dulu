import 'package:go_router/go_router.dart';
import 'package:bukadulu/presentation/pages/splash_screen.dart';
import 'package:bukadulu/presentation/pages/onboarding_page.dart';
import 'package:bukadulu/presentation/pages/auth/login_page.dart';
import 'package:bukadulu/presentation/pages/auth/register_page.dart';
import 'package:bukadulu/presentation/pages/dashboard/dashboard_page.dart';
import 'package:bukadulu/presentation/pages/venture/venture_create_page.dart';
import 'package:bukadulu/presentation/pages/idea/idea_capture_page.dart';
import 'package:bukadulu/presentation/pages/idea/idea_result_page.dart';
import 'package:bukadulu/presentation/pages/customer/customer_page.dart';
import 'package:bukadulu/presentation/pages/menu/menu_page.dart';
import 'package:bukadulu/presentation/pages/cost/cost_page.dart';
import 'package:bukadulu/presentation/pages/mission/mission_board_page.dart';
import 'package:bukadulu/presentation/pages/evidence/evidence_upload_page.dart';
import 'package:bukadulu/presentation/pages/score/score_page.dart';

final router = GoRouter(
  initialLocation: '/splash',
  routes: [
    GoRoute(path: '/splash', builder: (_, __) => const SplashScreen()),
    GoRoute(path: '/onboarding', builder: (_, __) => const OnboardingPage()),
    GoRoute(path: '/login', builder: (_, __) => const LoginPage()),
    GoRoute(path: '/register', builder: (_, __) => const RegisterPage()),
    GoRoute(path: '/dashboard', builder: (_, __) => const DashboardPage()),
    GoRoute(path: '/venture/new', builder: (_, __) => const VentureCreatePage()),
    GoRoute(path: '/venture/:id/idea', builder: (_, state) => IdeaCapturePage(ventureId: state.pathParameters['id']!)),
    GoRoute(path: '/venture/:id/idea/result', builder: (_, state) => IdeaResultPage(ventureId: state.pathParameters['id']!)),
    GoRoute(path: '/venture/:id/customer', builder: (_, state) => CustomerPage(ventureId: state.pathParameters['id']!)),
    GoRoute(path: '/venture/:id/menu', builder: (_, state) => MenuPage(ventureId: state.pathParameters['id']!)),
    GoRoute(path: '/venture/:id/cost', builder: (_, state) => CostPage(ventureId: state.pathParameters['id']!)),
    GoRoute(path: '/venture/:id/missions', builder: (_, state) => MissionBoardPage(ventureId: state.pathParameters['id']!)),
    GoRoute(path: '/venture/:id/mission/:missionId/evidence', builder: (_, state) => EvidenceUploadPage(
      ventureId: state.pathParameters['id']!,
      missionId: state.pathParameters['missionId']!,
    )),
    GoRoute(path: '/venture/:id/score', builder: (_, state) => ScorePage(ventureId: state.pathParameters['id']!)),
  ],
);
