package main

import (
    "fmt"
    "net/http"
)

func createNewPoint(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "I will create a new point entry")
}

func getPoints(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        w.WriteHeader(http.StatusBadRequest)
        fmt.Fprint(w, "Invalid request method")
        return
    }

    fmt.Fprintf(w, "I will get all the points for %s", r.FormValue("username"))
}