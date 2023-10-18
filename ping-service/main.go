package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	externalAPI := os.Getenv("EXTERNAL_API")
	user := os.Getenv("SECRET_USER")
	pass := os.Getenv("SECRET_PASS")

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		res := `
			I am PING service
			---
			EXTERNAL_API=` + externalAPI + `
			SECRET_USER=` + user + `
			SECRET_PASS=` + pass + `
		`

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(res))
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {

		res := `
			I am PING service
			---
			SECRET_USER=` + user + `
			SECRET_PASS=` + pass + `
		`

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(res))
	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
