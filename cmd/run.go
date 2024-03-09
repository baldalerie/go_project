package cmd

import (
	"log"
	"net/http"

	"HW5_RESTFULL_API.GO/handlers"
)

func Run() {

	http.HandleFunc("/item", handlers.Item)
	log.Println("server started on: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
