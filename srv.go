package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("%v\n", r)
    })
    fmt.Printf("Starting the Server\n")
    http.ListenAndServe(":3000", nil)

}
