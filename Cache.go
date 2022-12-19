package main

type cCache struct {
	cashes map[key]float64
}

type key struct {
	x, y, z int
}