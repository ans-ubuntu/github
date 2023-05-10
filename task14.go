package main

import (
    "fmt"
    "net/http"
)

func main() {
    // Serve static files
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/", fs)

    // API endpoint
    http.HandleFunc("/api", handleAPIRequest)

    // Start server
    fmt.Println("Server listening on port 8080...")
    http.ListenAndServe(":8080", nil)
}

func handleAPIRequest(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        // Retrieve name from query parameter
        name := r.URL.Query().Get("name")

        // Send response
        response := fmt.Sprintf("Hello, %s!", name)
        fmt.Fprintf(w, response)
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

