-- SaaS freemium + KlikQris payments + Google SSO support
-- idempotent: uses ADD COLUMN IF NOT EXISTS / CREATE TABLE IF NOT EXISTS

-- Auth provider tracking (email vs google) + SSO identifier
ALTER TABLE users ADD COLUMN IF NOT EXISTS provider TEXT NOT NULL DEFAULT 'email';
ALTER TABLE users ADD COLUMN IF NOT EXISTS provider_uid TEXT;
-- Google SSO users have no password; allow nullable by leaving password_hash as-is
-- and treating empty provider_uid as native email users.
CREATE UNIQUE INDEX IF NOT EXISTS idx_users_provider_uid ON users(provider, provider_uid) WHERE provider_uid IS NOT NULL;

-- Freemium quota: how many free idea validations the user has consumed
ALTER TABLE users ADD COLUMN IF NOT EXISTS free_quota_used INT NOT NULL DEFAULT 0;

-- Payment orders (KlikQris dynamic QRIS)
CREATE TABLE IF NOT EXISTS orders (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    venture_id TEXT REFERENCES ventures(id) ON DELETE SET NULL,
    purpose TEXT NOT NULL,                     -- 'idea_validation'
    amount INT NOT NULL,                       -- base price (10000)
    total_amount NUMERIC(12,2),                -- from KlikQris (amount + unique code)
    status TEXT NOT NULL DEFAULT 'pending' CHECK(status IN ('pending','paid','expired','failed')),
    qris_url TEXT,
    qris_image TEXT,                           -- base64 PNG
    signature TEXT,                            -- KlikQris signature (for webhook validation)
    klikqris_order_id TEXT,                    -- order_id sent to KlikQris
    expired_at TEXT,
    paid_at TEXT,
    fulfilled BOOLEAN NOT NULL DEFAULT FALSE,  -- whether the paid action was executed
    created_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP),
    updated_at TEXT NOT NULL DEFAULT (CURRENT_TIMESTAMP)
);
CREATE INDEX IF NOT EXISTS idx_orders_user ON orders(user_id);
CREATE INDEX IF NOT EXISTS idx_orders_status ON orders(status);
CREATE INDEX IF NOT EXISTS idx_orders_venture ON orders(venture_id);
