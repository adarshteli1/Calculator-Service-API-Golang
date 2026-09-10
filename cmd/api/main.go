package main

import (
	"Calculator/internal/routes"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error Loading the .env File")
	}


	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.StaticRoutes(router)

	fmt.Println("Starting Router")
	router.Run(":" + port)

}
