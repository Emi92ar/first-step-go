package main

import (
	"FromGithub/first-step-go/src/api/app"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := app.RunServer(port); err != nil {
		println("error occurs running.", err)
	}
}