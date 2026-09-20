package main

import (
	"log"

	"github.com/halimi/gopher-social/internal/db"
	"github.com/halimi/gopher-social/internal/env"
	"github.com/halimi/gopher-social/internal/store"
)

func main() {
	addr := env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/social?sslmode=disable")
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	store := store.NewStorage(conn)
	db.Seed(store)
}
