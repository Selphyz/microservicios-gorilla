package handlers

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"server/data"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, _ *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
	p.l.Println("Handle GET Product")
}
func (p *Products) PostProduct(_rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(&prod)
}
func (p *Products) PutProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to parse ID", http.StatusBadRequest)
	}
	p.l.Println("Handle PUT Product", id)
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct{}

func (p Products) MiddleWareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}
		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, fmt.Sprintf("Error reading product: %s", err), http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}
