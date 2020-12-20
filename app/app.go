package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/diogoqds/banking/domain"
	"github.com/diogoqds/banking/service"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// Start the application

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" ||
		os.Getenv("DB_USER") == "" ||
		os.Getenv("DB_PASSWORD") == "" ||
		os.Getenv("DB_HOST") == "" ||
		os.Getenv("DB_PORT") == "" ||
		os.Getenv("DB_NAME") == "" {
		log.Fatal("Environment variable not defined")
	}
}

func Start() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sanityCheck()
	handlers := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	router := mux.NewRouter()

	router.HandleFunc("/customers", handlers.getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", handlers.getCustomer).Methods(http.MethodGet)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}
