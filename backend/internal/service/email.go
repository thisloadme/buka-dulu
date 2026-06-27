package service

import (
	"fmt"
	"log/slog"
	"net/smtp"
	"strings"
)

type EmailConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	From     string
}

type EmailService struct {
	config EmailConfig
}

func NewEmailService(config EmailConfig) *EmailService {
	return &EmailService{config: config}
}

func (s *EmailService) SendOTP(to, otp string, expiryMinutes int) error {
	subject := "Kode OTP BukaDulu"
	body := fmt.Sprintf(`
<html>
<body style="font-family: Arial, sans-serif; background: #fafaf9; padding: 24px;">
  <div style="max-width: 480px; margin: 0 auto; background: #fff; border-radius: 8px; padding: 32px; border: 1px solid #e7e5e4;">
    <div style="text-align: center; margin-bottom: 24px;">
      <span style="font-size: 24px; font-weight: 300; color: #ea580c;">BukaDulu</span>
    </div>
    <h2 style="color: #1c1917; font-weight: 300; font-size: 20px; margin: 0 0 8px;">Verifikasi Email</h2>
    <p style="color: #57534e; font-size: 14px; line-height: 1.5; margin: 0 0 24px;">
      Gunakan kode OTP berikut untuk memverifikasi email kamu.
      Kode berlaku selama %d menit.
    </p>
    <div style="text-align: center; margin-bottom: 24px;">
      <span style="font-size: 36px; letter-spacing: 8px; font-weight: 700; color: #ea580c; background: #fff7ed; padding: 12px 24px; border-radius: 6px;">%s</span>
    </div>
    <p style="color: #57534e; font-size: 12px; line-height: 1.5; text-align: center;">
      Jika kamu tidak mendaftar akun BukaDulu, abaikan email ini.
    </p>
    <hr style="border: none; border-top: 1px solid #e7e5e4; margin: 24px 0;" />
    <p style="color: #a8a29e; font-size: 11px; text-align: center;">
      BukaDulu — F&B Validation Execution System
    </p>
  </div>
</body>
</html>`, expiryMinutes, otp)

	return s.send(to, subject, body)
}

func (s *EmailService) send(to, subject, htmlBody string) error {
	if s.config.Password == "" {
		slog.Warn("SMTP password not configured — skipping email send", "to", to, "subject", subject)
		return nil
	}

	addr := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
	auth := smtp.PlainAuth("", s.config.User, s.config.Password, s.config.Host)

	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%s",
		s.config.From, to, subject, htmlBody)

	if err := smtp.SendMail(addr, auth, s.config.From, []string{to}, []byte(msg)); err != nil {
		return fmt.Errorf("send email: %w", err)
	}

	slog.Info("email sent", "to", to, "subject", subject)
	return nil
}

// IsConfigured returns true if SMTP password is set
func (s *EmailService) IsConfigured() bool {
	return strings.TrimSpace(s.config.Password) != ""
}
