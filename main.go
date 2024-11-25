package main

import (
	"log"
	"net/http"
)

func main() {
	// Register API routes
	http.HandleFunc("/receipts/process", ProcessReceipts)
	http.HandleFunc("/receipts/", GetPoints)

	// Start the server
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
