package main

import (
    "net/http"
)

func main() {
    http.HandleFunc("/api/point/create", createNewPoint)
    http.HandleFunc("/api/point/get", getPoints)

    http.ListenAndServe(":8080", nil)
}