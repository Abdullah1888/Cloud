package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/tracking", ServeHTTP0)
	http.ListenAndServe(":9911", nil)
}
