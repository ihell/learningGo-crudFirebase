package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"go-firebase-crud/firebase"
	"go-firebase-crud/models"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateProduk(w http.ResponseWriter, r *http.Request) {
	var produk models.Produk
	json.NewDecoder(r.Body).Decode(&produk)

	_, _, err := firebase.Client.Collection("produk").Add(context.Background(), produk)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Produk ditambahkan"})
}

func GetAllProduk(w http.ResponseWriter, r *http.Request) {
	var hasil []models.Produk
	docs, err := firebase.Client.Collection("produk").Documents(context.Background()).GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, doc := range docs {
		var p models.Produk
		doc.DataTo(&p)
		hasil = append(hasil, p)
	}
	json.NewEncoder(w).Encode(hasil)
}

func UpdateProduk(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var produk models.Produk
	json.NewDecoder(r.Body).Decode(&produk)

	_, err := firebase.Client.Collection("produk").Doc(id).Set(context.Background(), produk)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Produk diperbarui"})
}

func DeleteProduk(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := firebase.Client.Collection("produk").Doc(id).Delete(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Produk dihapus"})
}
