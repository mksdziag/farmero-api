package main

import (
	"github.com/mksdziag/farmer-api/api"
	"github.com/mksdziag/farmer-api/db"
)

func main() {
	db.CreateDatabaseConnection()
	api.StartServer()
}
