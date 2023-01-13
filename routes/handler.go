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

func middleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println(r.Method, r.RequestURI)
			next.ServeHTTP(w, r)
	})
}

func middleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("Ini middleware khusus")
			next.ServeHTTP(w, r)
	})
}

func HandleRequests() {
	log.Println("Start the development server at http://127.0.0.1:9000")

	router := mux.NewRouter().StrictSlash(true)
	router.Use(middleware1)
	
	pref := router.PathPrefix("/product").Subrouter()
	pref.Use(middleware2)

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		res := Result{Code: 404, Message: "Method not found"}
		response, _ := json.Marshal(res)
		w.Write(response)
	})

	router.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)

		res := Result{Code: 405, Message: "Method not allowed"}
		response, _ := json.Marshal(res)
		w.Write(response)
	})

	pref.HandleFunc("/", controller.GetAll).Methods("GET")
	pref.HandleFunc("/{id}", controller.GetDetail).Methods("GET")
	pref.HandleFunc("/post", controller.Create).Methods("POST")
	pref.HandleFunc("/{id}", controller.Update).Methods("PUT")
	pref.HandleFunc("/{id}", controller.Delete).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", router))
}