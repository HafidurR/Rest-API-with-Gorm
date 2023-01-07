package routes

import (
	"api-gorm/controller"
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func HandleRequests() {
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

		res := Result{Code: 405, Message: "Method not allowed"}
		response, _ := json.Marshal(res)
		w.Write(response)
	})

	myRouter.HandleFunc("/product", controller.GetAll).Methods("GET")
	myRouter.HandleFunc("/product/{id}", controller.GetDetail).Methods("GET")
	myRouter.HandleFunc("/product", controller.Create).Methods("POST")
	myRouter.HandleFunc("/product/{id}", controller.Update).Methods("PUT")
	myRouter.HandleFunc("/product/{id}", controller.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", myRouter))
}