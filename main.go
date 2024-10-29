package main

import (
	"fmt"

	"serengeti.app/go-rest-template/pkg/api"
	"serengeti.app/go-rest-template/pkg/config"
	"serengeti.app/go-rest-template/pkg/db"
)

func main() {
	config.InitConfig()

	if err := db.InitDB(); err != nil {
		fmt.Println("Failed to initialize database:", err)
		return
	}

	fmt.Println("Starting application...")
	api.StartServer()
}
