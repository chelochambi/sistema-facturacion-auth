package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	var err error
	DB, err = sql.Open("postgres", "postgres://chelo:12345678@192.168.109.241:5432/sofi_facturacion?sslmode=disable")
	if err != nil {
		log.Fatal("Error de conexión a PostgreSQL:", err)
	}
}
