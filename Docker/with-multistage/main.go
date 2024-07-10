package main

import ( 
    "fmt"
    "net/http"
)
func helloDocker(w http.ResponseWriter, r *http.Request){ 
    fmt.Fprintf(w, "Hello, go!")
}
func main() {
    http.HandleFunc("/", helloDocker)
    fmt.Println("Server running on port 8080")
    http.ListenAndServe(":8080", nil)
}