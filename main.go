package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/kulmenevaa/go-start/app/routes"
	"os"
)

func main() {
	port := os.Getenv("APP_PORT")
	prefix := os.Getenv("ROUTE_PREFIX")

	router := gin.New()
	routes.ApiRoutes(prefix, router)
	router.Run(":" + port)
}
