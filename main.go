package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/tracking", ServeHTTP0)
	http.ListenAndServe("localhost:9911", nil)
}
