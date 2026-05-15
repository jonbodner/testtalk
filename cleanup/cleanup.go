package main

import (
	"context"
	"database/sql"
)

type Address struct {
	Country string
}

func FindAddress(ctx context.Context, db *sql.DB, name string) (Address, error) {
	return Address{
		Country: "Australia",
	}, nil
}
