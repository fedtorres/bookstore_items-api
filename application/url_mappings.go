package application

import (
	"github.com/fedtorres/bookstore_items-api/controllers"
	"net/http"
)

func mapUrls() {
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)

	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
}
