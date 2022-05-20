package main

import (
	"os"

	"github.com/djarum76-bot/taskplanner_golang/db"
	"github.com/djarum76-bot/taskplanner_golang/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	port := os.Getenv("PORT")

	e.Logger.Fatal(e.Start(":" + port))
}
