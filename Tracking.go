package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// /add
type AddService interface {
	PostSum(x int, y int)
	GetSum(x int, y int) (sum float64)
	PutSum(x int, y int)
	DeleteSum(x int, y int)
}

func ServeHTTP0(w http.ResponseWriter, r *http.Request) {
	num1 := r.URL.Query()["x"][0]
	num2 := r.URL.Query()["y"][0]
	long, _ := strconv.ParseFloat(num1, 64)
	lat, _ := strconv.ParseFloat(num2, 64)

	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "The closest Driver to (%f, %f) is Driver 3", float64(long), float64(lat))
	default:
		fmt.Fprintf(w, "Unsupported HTTP method!")
	}
}
