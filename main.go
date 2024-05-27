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
	http.HandleFunc("/branches/add", handlers.AddBranch)
	http.HandleFunc("/branches/update/", handlers.UpdateBranch)
	http.HandleFunc("/branches/delete/", handlers.DeleteBranch)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
