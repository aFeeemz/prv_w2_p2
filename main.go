package main

import (
	"log"
	"net/http"

	"preview/config"
	"preview/handlers"
)

func main() {
	config.InitDB()
	defer config.CloseDB()

	http.HandleFunc("/branches", handlers.GetBranches)
	http.HandleFunc("/branches/", handlers.GetBranch)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
