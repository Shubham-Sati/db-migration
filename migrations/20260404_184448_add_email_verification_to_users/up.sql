-- Migration: 20260404_184448_add_email_verification_to_users
-- Description: add_email_verification_to_users
-- Created: 2026-04-04 18:44:48
-- Direction: UP

-- Add email verification status column to users table
ALTER TABLE users ADD COLUMN is_verified BOOLEAN DEFAULT FALSE NOT NULL;
