package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func init() {
	DB = createDatabaseConnection()
}

func createDatabaseConnection() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=postgres password=postgres sslmode=disable host=api-db-1")

	if err != nil {
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("*****************************")
	log.Println("** ✔ Connected to database **")
	log.Println("*****************************")

	return db
}
