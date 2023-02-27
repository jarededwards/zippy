package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// This endpoint will show any content in the "public" directory
	// and will also serve the index.html file as the homepage
	mux.HandleFunc("/", http.FileServer(http.Dir("public")).ServeHTTP)

	// This is an imaginary endpoint to handle subroutes under /api
	mux.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is my API"))
	})

	mux.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		var content struct {
			Name string `json:"name"`
		}

		// I decode the request body into the content variable
		// and if it fails, I return a 400 Bad Request
		if err := json.NewDecoder(r.Body).Decode(&content); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Generate a custom response
		resp := map[string]interface{}{
			"message": "Hello " + content.Name,
		}

		// I encode the content variable into the response body
		// and if it fails, I return a 500 Internal Server Error
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// I create a server, no framework, just pure Go
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// I start the server, if fails, I log the error
	log.Fatal(srv.ListenAndServe())
}
