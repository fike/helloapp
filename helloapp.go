package main

import (
    "fmt"
    "net/http"
    "os"
)

var count int

func main() {
    name, err := os.Hostname()

    if err != nil {
      panic(err)
    }

    http.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Counter reset\r\n")
        count = 0
    })

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
      count = count + 1

      fmt.Fprint(w, name,"\r\n")
      fmt.Fprint(w, "The counter is: ", count, "\r\n")
    })

    http.ListenAndServe("0.0.0.0:8080", nil)

}
