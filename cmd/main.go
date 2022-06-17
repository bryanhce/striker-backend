package main

import (
	"fmt"

	"backend/cmd/app"
)

func main() {
	fmt.Println("code is running")

	a := app.App{}
	// a.Initialize(
	// 	os.Getenv("APP_DB_USERNAME"),
	// 	os.Getenv("APP_DB_PASSWORD"),
	// 	os.Getenv("APP_DB_NAME"))
	a.Initialize()

	a.Run(":8010")

	defer a.DB.Close()
}