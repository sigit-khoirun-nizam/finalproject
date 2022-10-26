package main

import (
	"finalProject/database"
	"finalProject/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
