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

func (ss *cCache) PostSum(x int, y int) {

	  res := float64(x + y)
	  ss.Post(x , y , 0,res)
}

func (ss *cCache) GetSum(x int, y int) (sum float64) {

	return ss.Get(x, y, 0)
}

func (ss *cCache) DeleteSum(x int, y int) {
	ss.Delete(x, y, 0)
}

func (ss *cCache) ServeHTTP0(w http.ResponseWriter, r *http.Request) {
	num1 := r.URL.Query()["x"][0]
	num2 := r.URL.Query()["y"][0]
	x, _ := strconv.Atoi(num1)
	y, _ := strconv.Atoi(num2)

	switch r.Method {
	case "POST":
		ss.PostSum(x, y)
		fmt.Fprintf(w, "New sum saved!")
	case "GET":
		sum := ss.GetSum(x, y)
		if sum == -1 {
			fmt.Fprintf(w, "Sum does not exist in cache!")
		} else {
			fmt.Fprintf(w, "Sum = %v", sum)
		}
	case "PUT":
	case "DELETE":
		ss.DeleteSum(x, y)
		fmt.Fprintf(w, "Sum deleted!")
	default:
		fmt.Fprintf(w, "Unsupported HTTP method!")
	}
}
