package main

import (
	"goGram/database"
	"goGram/routes"
	"os"
)

func main() {
	database.StartDB()
	r := routes.StartApp()
	r.Run(":" + os.Getenv("PORT"))
}
