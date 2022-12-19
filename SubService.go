package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// /sub

type SubService interface {
	PostSub(x int, y int)
	GetSub(x int, y int) (sub float64)
	PutSub(x int, y int)
	DeleteSub(x int, y int)
}

func (ss *cCache) PostSub(x int, y int) {

	res := float64(x - y)
	ss.Post(x , y , 1,res)
}

func (ss *cCache) GetSub(x int, y int) (sub float64) {
	
	return ss.Get(x, y, 1)
}

func (ss *cCache) DeleteSub(x int, y int) {
	ss.Delete(x, y, 1)
}

func (ss *cCache) ServeHTTP1(w http.ResponseWriter, r *http.Request) {
	num1 := r.URL.Query()["x"][0]
	num2 := r.URL.Query()["y"][0]
	x, _ := strconv.Atoi(num1)
	y, _ := strconv.Atoi(num2)

	switch r.Method {
	case "POST":
		ss.PostSub(x, y)
		fmt.Fprintf(w, "New sub saved!")
	case "GET":
		sub := ss.GetSub(x, y)
		if sub == -1 {
			fmt.Fprintf(w, "Sub does not exist in cache!")
		} else {
			fmt.Fprintf(w, "Sub = %v", sub)
		}
	case "PUT":
	case "DELETE":
		ss.DeleteSub(x, y)
		fmt.Fprintf(w, "Sub deleted!")
	default:
		fmt.Fprintf(w, "Unsupported HTTP method!")
	}
}
