package main

import (
	"fmt"
	"math"
)

func main() {
	hash := make(map[string]int)
	for i := 0; i < 3; i++ {
		h, _ := hash["a"]
		fmt.Println(h)
		h += 1
		hash["a"] = h
	}

	fmt.Println(hash["a"])
	var x int32
	var y int32
	x = 2
	y = 2
	fmt.Println(math.Pow(float64(x), float64(y)))
}
