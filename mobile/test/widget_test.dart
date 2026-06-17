import 'package:flutter_test/flutter_test.dart';
import 'package:bukadulu/app.dart';

void main() {
  testWidgets('App loads', (WidgetTester tester) async {
    await tester.pumpWidget(const BukaDuluApp());
    expect(find.text('BukaDulu'), findsOneWidget);
  });
}
