package main

import (
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/bxcodec/go-clean-arch/docs"
	"github.com/bxcodec/go-clean-arch/internal/adapter"
	"github.com/bxcodec/go-clean-arch/internal/rest/middleware"
	"github.com/bxcodec/go-clean-arch/pdx"
	"github.com/labstack/echo/v4"
	httpSwagger "github.com/swaggo/http-swagger"
)

const (
	defaultTimeout = 30
	defaultAddress = ":9090"
)

// @title PDF Service API
// @version 1.0
// @description API for handling PDF operations such as compression, merging, and splitting.
// @host localhost:9090
// @BasePath /api/v1
func main() {
	e := echo.New()
	e.Use(middleware.CORS)

	// Get context timeout from environment variables
	timeoutStr := os.Getenv("CONTEXT_TIMEOUT")
	timeout, err := strconv.Atoi(timeoutStr)
	if err != nil {
		log.Println("failed to parse timeout, using default timeout")
		timeout = defaultTimeout
	}
	timeoutContext := time.Duration(timeout) * time.Second
	e.Use(middleware.SetRequestContextWithTimeout(timeoutContext))

	// Healthcheck route for checking if the API is running
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(200, "PDF API IS RUNNING")
	})

	// Swagger route to serve Swagger UI
	e.GET("/swagger/*", echo.WrapHandler(httpSwagger.Handler()))

	// Initialize PDF adapter
	pdfAdapter := adapter.NewPDFCPUAdapter()

	// Build service Layer
	pdfSvc := pdx.New(pdfAdapter)
	_ = pdfSvc
	// Rest route handlers can be added here, such as:
	// rest.NewPDFHandler(e, pdfSvc)

	// Start the server on the configured address
	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		address = defaultAddress
	}
	log.Fatal(e.Start(address)) //nolint
}
