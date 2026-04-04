-- Migration: 20260403_134922_add_refresh_token_to_users
-- Description: add_refresh_token_to_users
-- Created: 2026-04-03 13:49:22
-- Direction: UP

-- Add refresh_token column to users table for storing JWT refresh tokens
ALTER TABLE users ADD COLUMN IF NOT EXISTS refresh_token TEXT;

