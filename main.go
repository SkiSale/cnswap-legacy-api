package main

import (
	"os"

	"github.com/SkiSale/cnswap-legacy-api/config"
	"github.com/SkiSale/cnswap-legacy-api/router"
	"github.com/joho/godotenv"

	_ "github.com/joho/godotenv/autoload"
)

func init() {
	godotenv.Load()
}

func main() {
	port := os.Getenv("PORT")

	init := config.Init()
	app := router.Init(init)

	app.Run(":" + port)
}
