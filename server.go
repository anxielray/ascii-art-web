package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// read the html file
	html, err := os.ReadFile("index.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading HTML file: %v", err), http.StatusInternalServerError)
	}

	// Write the HTML content as response
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(html)
}
