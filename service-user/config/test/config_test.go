package config_test

import (
	"service-user/config"
	"testing"
)

func TestGetPostgresDB(t *testing.T) {
	db := config.GetPostgresDB()
	if db == nil {
		t.Error("Expected non-nil database connection, got nil")
	}
}
