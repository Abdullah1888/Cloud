package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// div

type DivService interface {
	PostDiv(x int, y int)
	GetDiv(x int, y int) (div float64)
	PutDiv(x int, y int)
	DeleteDiv(x int, y int)
}

func (ss *cCache) PostDiv(x int, y int) {

	res := float64 (x) / float64 (y)
	ss.Post(x , y , 2,res)
}

func (ss *cCache) GetDiv(x int, y int) (div float64) {
	
	return ss.Get(x, y, 2)
}

func (ss *cCache) DeleteDiv(x int, y int) {
	ss.Delete(x, y, 2)
}

func (ss *cCache) ServeHTTP2(w http.ResponseWriter, r *http.Request) {
	num1 := r.URL.Query()["x"][0]
	num2 := r.URL.Query()["y"][0]
	x, _ := strconv.Atoi(num1)
	y, _ := strconv.Atoi(num2)

	switch r.Method {
	case "POST":
		ss.PostDiv(x, y)
		fmt.Fprintf(w, "New div saved!")
	case "GET":
		div := ss.GetDiv(x, y)
		if div == -1 {
			fmt.Fprintf(w, "Div does not exist in cache!")
		} else {
			fmt.Fprintf(w, "Div = %v", div)
		}
	case "PUT":
	case "DELETE":
		ss.DeleteDiv(x, y)
		fmt.Fprintf(w, "Div deleted!")
	default:
		fmt.Fprintf(w, "Unsupported HTTP method!")
	}
}
