package main

import (
	"miniProject/config"
	"miniProject/database"
	"miniProject/routes"
)

func main() {
	database.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(config.PORT))

	
}
