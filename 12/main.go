package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	arr := [...]string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(arr)

	//var count int
	for _, val := range arr {
		m[val] += 1
	}

	var set []string
	for key, _ := range m {
		if m[key] > 1 {
			for i := 0; i < m[key]; i++ {
				set = append(set, key)
			}
		}
	}

	fmt.Println(m)
	fmt.Println(set)
}
