package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/load_test", forLoadTesting).Methods("POST")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8001",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

type Request struct {
	Password string `json:"password"`
}

func forLoadTesting(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var request Request
	err := decoder.Decode(&request)
	if err != nil {
		panic(err)
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), 5)
	if err != nil {
		panic(err)
	}

	fmt.Println("password generated: ", string(bytes))

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
