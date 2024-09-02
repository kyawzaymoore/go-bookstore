// @title Go Bookstore API
// @version 1.0
// @description This is a sample server for a bookstore.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/kzm/go-bookstore/api/config"
	_ "github.com/kzm/go-bookstore/api/docs"
	"github.com/kzm/go-bookstore/api/routes"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

func main() {
	appEnv := os.Getenv("APP_ENV")
	errorPath := os.Getenv("CONFIG_PATH")
	configPath := os.Getenv("ERROR_PATH")

	fmt.Printf("AppEnv %s\n ErrorPath %s\n ConfigPath %s\n", appEnv, errorPath, configPath)

	// Load the configuration
	config.LoadConfig()

	// Access the configuration values
	fmt.Printf("Server is running on %s:%d\n", config.AppConfig.Server.Host, config.AppConfig.Server.Port)
	fmt.Printf("Database: %s@%s:%d/%s\n", config.AppConfig.Database.User, config.AppConfig.Database.Host, config.AppConfig.Database.Port, config.AppConfig.Database.DBName)

	r := mux.NewRouter()

	// Serve Swagger UI
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	routes.BookRoutes(r)
	routes.CategoryRoutes(r)
	routes.AuthRoutes(r)
	routes.ProfileRoutes(r)

	// Enable CORS for all routes
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"*"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"*"},
        AllowCredentials: true,
    })

	// Wrap the router with the CORS middleware
    handler := c.Handler(r)

	//http.Handle("/", r)

	// Build the address string using the host and port from the config
	address := fmt.Sprintf("%s:%d", config.AppConfig.Server.Host, config.AppConfig.Server.Port)

	log.Fatal(http.ListenAndServe(address, handler))
}
