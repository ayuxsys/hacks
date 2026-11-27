package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPushSubdomain(t *testing.T) {
	conn, err := Connect()
	require.NoError(t, err)
	require.NoError(t, conn.Migrate())
	subdomains := SubFinder{
		Domain: "example.com",
		Subs:   []string{"foobar.example.com", "test.example.com", "zeehaha.example.com"},
	}
	require.NoError(t, conn.Create(&subdomains).Error)
}

func TestPSQLMigrate(t *testing.T) {
	conn, err := Connect()
	require.NoError(t, err)
	require.NoError(t, conn.Migrate())
}

func TestPSQLConnect(t *testing.T) {
	_, err := Connect()
	require.NoError(t, err)
}
