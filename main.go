package main

import (
	"fmt"
	"log"
	config "myapp/Config"
	models "myapp/Models"
	routes "myapp/Router"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = config.InitDB()
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}

	err = config.DB.AutoMigrate(models.Employee{}, models.Department{})
	if err != nil {
		fmt.Printf("faild to migrate schema: %v", err)
	}

	app := gin.Default()
	routes.RegisterRouter(app)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8090"
	}

	log.Printf("server listening on port %s", port)
	if err := http.ListenAndServe(":"+port, app); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
