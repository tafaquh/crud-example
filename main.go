package main

import (
	"github.com/joho/godotenv"
	"github.com/tafaquh/crud-example/app"
)

func main() {
	_ = godotenv.Load()

	app.StartApplication()
}
