package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func redirect(w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "http://www.google.com", 301)
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/redirect", redirect)
    http.ListenAndServe(":8080", nil)
}
