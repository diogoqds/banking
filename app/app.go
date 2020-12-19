package app

import (
	"log"
	"net/http"

	"github.com/diogoqds/banking/domain"
	"github.com/diogoqds/banking/service"
	"github.com/gorilla/mux"
)

// Start the application
func Start() {

	handlers := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}
	router := mux.NewRouter()

	router.HandleFunc("/customers", handlers.getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", handlers.getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:3000", router))
}
