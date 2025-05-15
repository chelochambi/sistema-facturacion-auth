package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	// Carga el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error al cargar .env:", err)
	}

	// Construye la cadena de conexión
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("❌ Error al abrir la conexión con PostgreSQL:", err)
	}

	// Verifica que realmente se pueda conectar
	if err := DB.Ping(); err != nil {
		log.Fatal("❌ No se pudo conectar a PostgreSQL:", err)
	}

	log.Println("✅ Conectado a PostgreSQL exitosamente")
}
