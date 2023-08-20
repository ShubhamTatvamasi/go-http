package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var number int = 0

func handler(w http.ResponseWriter, r *http.Request) {
	number++
	fmt.Fprintf(w, "Welcome to go http "+strconv.Itoa(number)+"\n")
}

func main() {

	http.HandleFunc("/", handler)

	fmt.Println("Server is listening...")
	log.Fatal(http.ListenAndServe(":80", nil))

}
