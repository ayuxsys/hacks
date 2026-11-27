package main

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.M) {
	_ = godotenv.Load(".env")
}

func TestPushSubdomain(t *testing.T) {
	require.NoError(t, Connect(os.Getenv("DB_URL")))
	require.NoError(t, databaseConn.Migrate())
	subdomains := SubFinder{
		Domain: "example.com",
		Subs:   []string{"foobar.example.com", "test.example.com", "zeehaha.example.com"},
	}
	require.NoError(t, databaseConn.Create(&subdomains).Error)
}

func TestPSQLMigrate(t *testing.T) {
	require.NoError(t, Connect(os.Getenv("DB_URL")))
	require.NoError(t, databaseConn.Migrate())
}

func TestPSQLConnect(t *testing.T) {
	require.NoError(t, Connect(os.Getenv("DB_URL")))
}
