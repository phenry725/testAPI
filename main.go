package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
  "encoding/json"
  "log"
)

type Response struct {
	Name string `json:"name"`
}

func main() {

	//API
	router := mux.NewRouter()
	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	//routes for leagues
	router.HandleFunc("/", testMethod).Methods("GET")

	log.Fatal(srv.ListenAndServe())
}

func testMethod(w http.ResponseWriter, r *http.Request) {
  log.Println("Request")
	// log.Println( )
	test := Response{Name: "Blah"}
	json.NewEncoder(w).Encode(test)
}
