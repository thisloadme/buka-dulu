import 'dart:async';
import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:go_router/go_router.dart';
import 'package:bukadulu/presentation/providers/auth_provider.dart';

class OTPVerifyPage extends ConsumerStatefulWidget {
  final String email;
  const OTPVerifyPage({super.key, required this.email});

  @override
  ConsumerState<OTPVerifyPage> createState() => _OTPVerifyPageState();
}

class _OTPVerifyPageState extends ConsumerState<OTPVerifyPage> {
  final _otpController = TextEditingController();
  int _countdown = 60;
  Timer? _timer;
  bool _canResend = false;

  @override
  void initState() {
    super.initState();
    _startCountdown();
  }

  @override
  void dispose() {
    _otpController.dispose();
    _timer?.cancel();
    super.dispose();
  }

  void _startCountdown() {
    _canResend = false;
    _countdown = 60;
    _timer?.cancel();
    _timer = Timer.periodic(const Duration(seconds: 1), (timer) {
      setState(() {
        _countdown--;
        if (_countdown <= 0) {
          _canResend = true;
          timer.cancel();
        }
      });
    });
  }

  Future<void> _verify() async {
    final otp = _otpController.text.trim();
    if (otp.length != 6) return;

    await ref.read(authProvider.notifier).verifyOTP(widget.email, otp);
    final authState = ref.read(authProvider);
    authState.whenOrNull(data: (data) {
      if (data != null && data.token.isNotEmpty) {
        context.go('/dashboard');
      }
    });
  }

  Future<void> _resend() async {
    try {
      await ref.read(authProvider.notifier).resendOTP(widget.email);
      _startCountdown();
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('Kode OTP baru telah dikirim')),
        );
      }
    } catch (e) {
      if (mounted) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(content: Text('Gagal: $e')),
        );
      }
    }
  }

  @override
  Widget build(BuildContext context) {
    final authState = ref.watch(authProvider);

    return Scaffold(
      appBar: AppBar(title: const Text('Verifikasi Email')),
      body: SafeArea(
        child: Center(
          child: SingleChildScrollView(
            padding: const EdgeInsets.all(24),
            child: Column(
              mainAxisSize: MainAxisSize.min,
              children: [
                Text('Cek Email Kamu', style: Theme.of(context).textTheme.headlineMedium?.copyWith(fontWeight: FontWeight.w400)),
                const SizedBox(height: 12),
                Text(
                  'Kami telah mengirim kode OTP ke\n${widget.email}',
                  textAlign: TextAlign.center,
                  style: Theme.of(context).textTheme.bodyMedium?.copyWith(color: const Color(0xFF57534e)),
                ),
                const SizedBox(height: 32),
                TextField(
                  controller: _otpController,
                  keyboardType: TextInputType.number,
                  textAlign: TextAlign.center,
                  maxLength: 6,
                  style: const TextStyle(fontSize: 32, letterSpacing: 12, fontWeight: FontWeight.w700),
                  decoration: const InputDecoration(
                    counterText: '',
                    hintText: '------',
                    hintStyle: TextStyle(color: Color(0xFFd6d3d1), fontSize: 32, letterSpacing: 12),
                    border: OutlineInputBorder(),
                  ),
                  onChanged: (_) {
                    if (_otpController.text.length == 6) _verify();
                  },
                ),
                const SizedBox(height: 24),
                SizedBox(
                  width: double.infinity,
                  height: 48,
                  child: ElevatedButton(
                    onPressed: (authState.isLoading || _otpController.text.length != 6) ? null : _verify,
                    child: authState.isLoading
                        ? const SizedBox(width: 20, height: 20, child: CircularProgressIndicator(strokeWidth: 2))
                        : const Text('Verifikasi'),
                  ),
                ),
                if (authState.hasError)
                  Padding(
                    padding: const EdgeInsets.only(top: 16),
                    child: Text(
                      '${authState.error}',
                      style: const TextStyle(color: Color(0xFFdc2626)),
                      textAlign: TextAlign.center,
                    ),
                  ),
                const SizedBox(height: 24),
                if (_canResend)
                  TextButton(onPressed: _resend, child: const Text('Kirim ulang OTP'))
                else
                  Text(
                    'Kirim ulang dalam $_countdown detik',
                    style: const TextStyle(color: Color(0xFFa8a29e)),
                  ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
