-- Add OTP verification support to users table
ALTER TABLE users ADD COLUMN otp_code TEXT;
ALTER TABLE users ADD COLUMN otp_expires_at TEXT;
ALTER TABLE users ADD COLUMN otp_verified_at TEXT;

-- Users start as 'pending' until OTP verified, then become 'active'
-- Update existing users to be active
UPDATE users SET status = 'active' WHERE status = 'active';
