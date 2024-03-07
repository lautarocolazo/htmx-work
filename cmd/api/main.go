package main

import (
	"fmt"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"

	"htmx-fiber/internal/server"
)

func main() {

	server := server.New()

	server.SetupStatic()

	server.RegisterFiberRoutes()
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	err := server.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
