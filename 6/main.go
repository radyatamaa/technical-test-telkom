package main

import (
	"log"
	"os"
	"strconv"
	"technical-test-telkom/dbconn"
	"technical-test-telkom/helper/logger"
	"technical-test-telkom/middleware"
	"technical-test-telkom/models"
	"time"

	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"

	_cartHttpHandler "technical-test-telkom/cart/delivery/http"
	_cartUsecase "technical-test-telkom/cart/usecase"
	_cartRepo "technical-test-telkom/cart/repository"

	_ "technical-test-telkom/swagger"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Echo Swagger technical-test-telkom API
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:5000
// @BasePath /
// @schemes http
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbConn := dbconn.DB()

	// db auto migrate dev environment
	if err := dbConn.AutoMigrate(
		&models.CartProduk{}); err != nil {
		panic(err)
	}

	l := logger.L
	appPort := os.Getenv("APP_PORT")
	timeout, _ := strconv.Atoi(os.Getenv("APP_TIMEOUT"))

	e := echo.New()
	middL := middleware.InitMiddleware()
	e.Use(middL.Log)
	e.Use(middL.CORS)

	cartRepo := _cartRepo.NewCartRepository(dbConn, l)

	timeoutContext := time.Duration(timeout) * time.Second

	cartUsecase := _cartUsecase.NewCartUsecase(cartRepo, l, timeoutContext)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	gapis := e.Group("/api")
	_cartHttpHandler.NewCartHandler(gapis,e, cartUsecase, l)

	log.Fatal(e.Start(":" + appPort))
}
