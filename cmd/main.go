package main

import (
	"backend/cmd/app"
	"fmt"
	"os"
)

func main() {
	fmt.Println("code is running")

	a := app.App{}

	a.Initialize()

	port := os.Getenv("PORT")
	// port := "8000"
	a.Run(":" + port)

	defer a.DB.Close()
}