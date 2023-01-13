package controller

import (
	"api-gorm/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

var err error

type Product struct {
	ID          int    `form:"id" json:"id"`
	Name        string `form:"name" json:"name"`
	Description string `form:"code" json:"description"`
	Stock       int    `json:"stock"`
}

type Result struct {
	Success 	bool   		`json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type NotFoundResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

var db = config.Connect()

func check(productID string) bool {
	var product Product
	db.First(&product, productID)
	if product.ID == 0 {
		return false
	}
	return true
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	products := []Product{}
	db.Find(&products)

	res := Result{Success: true, Message: "Success get products", Data: products}
	results, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func GetDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]

	var product Product

	if check(productID) == false {
		res := NotFoundResponse{Success: false, Message: "Data not found"}
		result, err := json.Marshal(res)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(result)
		return
	}

	db.First(&product, productID)

	res := Result{Success: true, Message: "Success get detail product", Data: product}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func Create(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var product Product
	json.Unmarshal(payloads, &product)

	db.Create(&product)

	res := Result{Success: true, Message: "Success create product", Data: product}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(result)
}

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var productUpdates Product
	json.Unmarshal(payloads, &productUpdates)

	if check(productID) == false {
		res := NotFoundResponse{Success: false, Message: "Data not found"}
		result, err := json.Marshal(res)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(result)
		return
	}

	var product Product
	db.First(&product, productID)
	db.Model(&product).Updates(productUpdates)

	res := Result{Success: true, Message: "Success update product", Data: product}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productID := vars["id"]

	if check(productID) == false {
		res := NotFoundResponse{Success: false, Message: "Data not found"}
		result, err := json.Marshal(res)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(result)
		return
	}

	var product Product
	db.First(&product, productID)
	db.Delete(&product)

	res := Result{Success: true, Message: "Success delete products", Data: product}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
