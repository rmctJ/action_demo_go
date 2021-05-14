package main

import (
	calc "action_demo_go/calculator"
	"fmt"
	"log"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/pi" {
		fmt.Fprint(w, calc.Pi())
	} else {
		http.NotFound(w, r)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Default port %s", port)
	} else {
		log.Printf("Listening port %s", port)
	}

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
	
}
