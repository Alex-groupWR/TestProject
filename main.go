package main

import (
	"testProject/app/routes"
	database "testProject/config"
)

func main() {
	db := database.InitDatabaseConnection()

	routes.SetupRouter(db)
}
