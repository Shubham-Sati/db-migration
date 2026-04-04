-- Migration: 20260403_134922_add_refresh_token_to_users
-- Description: add_refresh_token_to_users
-- Created: 2026-04-03
-- Direction: DOWN

-- Rollback: Remove refresh_token column from users table
ALTER TABLE users DROP COLUMN IF EXISTS refresh_token;