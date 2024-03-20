package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type cartProduct struct {
	ProductId string `json:"product_id"`
	Qty       int    `json:"qty"`
}

// handler
func cartProductsHandler(w http.ResponseWriter, r *http.Request) {
	products := []cartProduct{
		{
			ProductId: "A0001",
			Qty:       1,
		},
		{
			ProductId: "A0002",
			Qty:       5,
		},
		{
			ProductId: "B0001",
			Qty:       1,
		},
	}

	bytes, err := json.Marshal(products)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)

}

type Genre int

const (
	UNKNOWN = iota
	CD
	DVD
	BOOK
	GOODS
)

type product struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Genre Genre  `json:"genre"`
	Price int    `json:"price"`
}

func productInfoHandler(w http.ResponseWriter, r *http.Request) {
	products := map[string]product{
		"A0001": {
			Id:    "A0001",
			Name:  "世界一流エンジニアの思考法",
			Genre: BOOK,
			Price: 1700,
		},
		"A0002": {
			Id:    "A0002",
			Name:  "Clean Architecture 達人に学ぶソフトウェアの構造と設計",
			Genre: BOOK,
			Price: 3168,
		},
		"B0001": {
			Id:    "B0001",
			Name:  "映画『THE FIRST SLAM DUNK』STANDARD EDITION",
			Genre: DVD,
			Price: 3523,
		},
	}

	vars := mux.Vars(r)
	bytes, err := json.Marshal(products[vars["id"]])
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/cart/products", cartProductsHandler)
	r.HandleFunc("/products/{id}", productInfoHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8081",
	}
	log.Fatal(srv.ListenAndServe())
	log.Printf("start server port: %v", "8081")
}
