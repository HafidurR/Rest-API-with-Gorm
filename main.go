package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

// Product is a representation of a product
type Product struct {
	ID    int             `form:"id" json:"id"`
	Name  string          `form:"name" json:"name"`
	Description  string 	`form:"code" json:"description"`
	Stock       int    		`json:"stock"`
}

// Result is an array of product
type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// Main
func main() {
	db, err = gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_store")

	if err != nil {
		log.Println("Connection failed", err)
	} else {
		log.Println("Connection established")
	}

	// db.AutoMigrate(&Product{})
	handleRequests()
}

func handleRequests() {
	log.Println("Start the development server at http://127.0.0.1:9000")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		res := Result{Code: 404, Message: "Method not found"}
		response, _ := json.Marshal(res)
		w.Write(response)
	})

	myRouter.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)

		res := Result{Code: 403, Message: "Method not allowed"}
		response, _ := json.Marshal(res)
		w.Write(response)
	})

	myRouter.HandleFunc("/product", getAll).Methods("GET")
	myRouter.HandleFunc("/product/{id}", getDetail).Methods("GET")
	myRouter.HandleFunc("/product", create).Methods("POST")
	myRouter.HandleFunc("/product/{id}", update).Methods("PUT")
	myRouter.HandleFunc("/product/{id}", delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", myRouter))
}

func getAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: get products")

	products := []Product{}
	db.Find(&products)

	res := Result{Code: 200, Data: products, Message: "Success get products"}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func getDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]

	var product Product

	db.First(&product, productID)

	res := Result{Code: 200, Data: product, Message: "Success get product"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func create(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var product Product
	json.Unmarshal(payloads, &product)

	db.Create(&product)

	res := Result{Code: 200, Data: product, Message: "Success create product"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var productUpdates Product
	json.Unmarshal(payloads, &productUpdates)

	var product Product
	db.First(&product, productID)
	db.Model(&product).Updates(productUpdates)

	res := Result{Code: 200, Data: product, Message: "Success update product"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]

	var product Product

	db.First(&product, productID)
	db.Delete(&product)

	res := Result{Code: 200, Message: "Success delete product"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}