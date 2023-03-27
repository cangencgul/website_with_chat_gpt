package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    // Define routes
    http.HandleFunc("/", helloWorld)
    http.HandleFunc("/auth", authenticate(authenticated))

    // Get the port number from the PORT environment variable
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default to port 8080
    }

    // Start server
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, World!")
}

func authenticate(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        username, password, ok := r.BasicAuth()
        if !ok || username != os.Getenv("AUTH_USERNAME") || password != os.Getenv("AUTH_PASSWORD") {
            w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        next.ServeHTTP(w, r)
    }
}

func authenticated(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, authenticated user!")
}
