package utils

import (
	"strings"

	"github.com/google/uuid"
)

// UUIDWithPrefix creates a unique public identifier with a prefix
// Format: prefix_uuid (e.g., "usr_123e4567e89b12d3a456426614174000")
// Removes all hyphens and underscores from UUID for clean format
func UUIDWithPrefix(prefix string) string {
	id := uuid.New().String()
	id = prefix + "_" + id
	id = strings.ReplaceAll(id, "-", "")
	return id
}
