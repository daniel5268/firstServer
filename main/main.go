package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/sayhi", handleSayHi)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handleSayHi(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hi!\n")
}
