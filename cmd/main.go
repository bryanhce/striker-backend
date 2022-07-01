package main

import (
	"backend/cmd/app"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("code is running")

	a := app.App{}

	a.Initialize()

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("could not load env file")
	}
	port := os.Getenv("PORT")
	a.Run(":" + port)

	defer a.DB.Close()
}