package main

import (
	"net/http"
)

func main() {
	ss := &cCache{
		cashes: map[key]float64{},
	}

	http.HandleFunc("/add", ss.ServeHTTP0)
	http.HandleFunc("/sub", ss.ServeHTTP1)
	http.HandleFunc("/div", ss.ServeHTTP2)
	http.HandleFunc("/mul", ss.ServeHTTP3)
	http.ListenAndServe("localhost:9911", nil)
}
