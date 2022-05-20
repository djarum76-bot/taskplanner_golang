package main

import (
	"github.com/djarum76-bot/taskplanner_golang/db"
	"github.com/djarum76-bot/taskplanner_golang/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":8080"))
}
