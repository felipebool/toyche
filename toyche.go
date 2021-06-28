package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var cache = NewCache()

func locationUpdateHandler(Cache *Cache) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {}
}

func locationDeleteHandler(Cache *Cache) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {}
}

func locationRetrieveHandler(Cache *Cache) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {}
}

func main() {
	cache := NewCache()

	r := mux.NewRouter()

	r.HandleFunc("/location/{order_id}", locationUpdateHandler(cache)).Methods(http.MethodPut)
	r.HandleFunc("/location/{order_id}", locationDeleteHandler(cache)).Methods(http.MethodDelete)
	r.HandleFunc("/location/{order_id}", locationRetrieveHandler(cache)).Methods(http.MethodGet)

	http.ListenAndServe(":8080", r)
}
