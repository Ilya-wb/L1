package main

import (
	"fmt"
	"math/rand"
)

func main() {
	m := make(map[int]int)
	var arr1 [10]int
	var arr2 [5]int

	fmt.Printf("arr1: {")
	for i, _ := range arr1 {
		arr1[i] = rand.Intn(10)
		fmt.Printf(" %d", arr1[i])
	}
	fmt.Printf("}\narr2: {")

	for i, _ := range arr2 {
		arr2[i] = rand.Intn(10)
		fmt.Printf(" %d", arr2[i])
	}
	fmt.Printf("}\n")

	//var count int
	for _, val := range arr1 {
		m[val] = 1
	}
	for _, val := range arr2 {
		m[val] += 1
	}

	fmt.Println(m)
}
