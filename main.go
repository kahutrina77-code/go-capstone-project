// package main  // This file is the starting point of my program

// import "fmt" //I need the fmt package to print things

// func main() {   // Go starts executing my program RIGHT HERE
//     fmt.Println("Hello from Nairobi!")
//     fmt.Println("My first Go program!")
// }
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

// Response struct defines the shape of our JSON response
type Response struct {
    Message   string `json:"message"`
    Location  string `json:"location"`
    Status    string `json:"status"`
    Timestamp string `json:"timestamp"`
}

// homeHandler handles the "/" route
func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to my Go API! 🇰🇪\nVisit /api for JSON response\nVisit /about for project info\nVisit /time for current time")
}

// apiHandler handles the "/api" route - returns JSON
func apiHandler(w http.ResponseWriter, r *http.Request) {
    // Set response header to JSON
    w.Header().Set("Content-Type", "application/json")

    // Create response object
    response := Response{
        Message:   "Hello from Nairobi!",
        Location:  "Nairobi, Kenya 🇰🇪",
        Status:    "success",
        Timestamp: time.Now().Format("2006-01-02 15:04:05"),
    }

    // Convert to JSON and send
    json.NewEncoder(w).Encode(response)
}

// aboutHandler handles the "/about" route
func aboutHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Go Capstone Project\nBuilt by: Trina Luseno\nSchool: Moringa School\nTech: Go 1.21.6")
}

// timeHandler handles the "/time" route
func timeHandler(w http.ResponseWriter, r *http.Request) {
    currentTime := time.Now().Format("Monday, 02 January 2006 - 15:04:05")
    fmt.Fprintf(w, "Current Time: %s", currentTime)
}

func main() {
    // Register all routes
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/api", apiHandler)
    http.HandleFunc("/about", aboutHandler)
    http.HandleFunc("/time", timeHandler)

    // Start server
    fmt.Println("🚀 Server starting...")
    fmt.Println("📍 Visit: http://localhost:8080")
    fmt.Println("📍 API:   http://localhost:8080/api")
    fmt.Println("📍 About: http://localhost:8080/about")
    fmt.Println("📍 Time:  http://localhost:8080/time")

    http.ListenAndServe(":8080", nil)
}