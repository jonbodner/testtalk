package main

import (
	"database/sql"
	"testing"

	_ "github.com/proullon/ramsql/driver"
)

func TestAddressQuery(t *testing.T) {
	db := setupDB(t, "address")
	address, err := FindAddress(t.Context(), db, "Fred")
	if err != nil {
		t.Error(err)
	}
	if address.Country != "Australia" {
		t.Error("wrong country")
	}
}

func setupDB(t *testing.T, tableName string) *sql.DB {
	t.Helper()
	db, err := sql.Open("ramsql", tableName)
	if err != nil {
		t.Fatal("error creating database", err)
	}
	t.Cleanup(func() { db.Close() })
	return db
}
