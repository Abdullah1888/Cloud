package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// mul

type MulService interface {
	PostMul(x int, y int)
	GetMul(x int, y int) (mul float64)
	PutMul(x int, y int)
	DeleteMul(x int, y int)
}

func (ss *cCache) PostMul(x int, y int) {

	res := float64(x * y)
	  ss.Post(x , y , 3,res)
}

func (ss *cCache) GetMul(x int, y int) (mul float64) {
	return ss.Get(x, y, 3)
}

func (ss *cCache) DeleteMul(x int, y int) {
	ss.Delete(x, y, 3)
}

func (ss *cCache) ServeHTTP3(w http.ResponseWriter, r *http.Request) {
	num1 := r.URL.Query()["x"][0]
	num2 := r.URL.Query()["y"][0]
	x, _ := strconv.Atoi(num1)
	y, _ := strconv.Atoi(num2)

	switch r.Method {
	case "POST":
		ss.PostMul(x, y)
		fmt.Fprintf(w, "New mul saved!")
	case "GET":
		mul := ss.GetMul(x, y)
		if mul == -1 {
			fmt.Fprintf(w, "Mul does not exist in cache!")
		} else {
			fmt.Fprintf(w, "Mul = %v", mul)
		}
	case "PUT":
	case "DELETE":
		ss.DeleteMul(x, y)
		fmt.Fprintf(w, "Mul deleted!")
	default:
		fmt.Fprintf(w, "Unsupported HTTP method!")
	}
}
