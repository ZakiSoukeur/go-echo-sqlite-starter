package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-starter/internal/db"
	_ "github.com/go-starter/internal/docs"
	"github.com/go-starter/internal/router"
	"github.com/go-starter/internal/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server for using Swagger with Echo.
// @host localhost:3000
// @BasePath /
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("no .env file")
	}
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "example.sqlite"
	}
	// init connection
	err = db.InitDB(dsn)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	conn := db.GetDB()
	queries := db.New(conn)
	authorService := service.NewAuthorService(queries)

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           `[${time_custom}] ${method} ${uri} - ${status} - ${latency_human} - ${bytes_in} in / ${bytes_out} out - error: ${error}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
		Output:           os.Stdout,
	}))

	// All author routes grouped in router package
	router.RegisterAuthorRoutes(e, authorService)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}

	e.GET("/docs/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(port))
}
