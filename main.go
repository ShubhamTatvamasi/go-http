package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
)


func handler(w http.ResponseWriter, r *http.Request) {
    number := 0
    number++
    fmt.Fprintf(w, "Welcome to go http " + strconv.Itoa(number) + "\n")
}

func main() {


    http.HandleFunc("/", handler)

    log.Fatal(http.ListenAndServe(":80", nil))

}
