-- Migration: 20260404_184448_add_email_verification_to_users
-- Description: add_email_verification_to_users
-- Created: 2026-04-04 18:44:48
-- Direction: DOWN

-- Remove email verification status column from users table
ALTER TABLE users DROP COLUMN IF EXISTS is_verified;
