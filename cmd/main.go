package main

import (
	"fmt"

	"backend/cmd/app"
)

func main() {
	fmt.Println("code is running")

	a := app.App{}

	a.Initialize()

	a.Run(":8000")

	defer a.DB.Close()
}