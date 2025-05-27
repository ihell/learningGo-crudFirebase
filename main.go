package main

import (
	"learningGo-crudFirebase/firebase"
	"learningGo-crudFirebase/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	firebase.InitFirestore()
	router := mux.NewRouter()

	router.HandleFunc("/produk", handlers.CreateProduk).Methods("POST")
	router.HandleFunc("/produk", handlers.GetAllProduk).Methods("GET")
	router.HandleFunc("/produk/{id}", handlers.UpdateProduk).Methods("PUT")
	router.HandleFunc("/produk/{id}", handlers.DeleteProduk).Methods("DELETE")

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
