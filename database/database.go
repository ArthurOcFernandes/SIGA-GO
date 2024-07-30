package database

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	_ "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connection() *sql.DB {

	err := godotenv.Load("Properties/db_connection.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("user=%s dbname=%s password=%s host=localhost sslmode=disable", dbUser, dbName, dbPass)
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func SelectAllUserJoinAddress(userType string) *sql.Rows {
	db := Connection()

	connectionString := fmt.Sprintf(`SELECT * FROM %s inner join addresses using (id_address)`, userType)

	rows, err := db.Query(connectionString)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	return rows
}
