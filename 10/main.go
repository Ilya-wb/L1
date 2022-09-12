package main

import (
	"fmt"
)

func main() {
	values := [...]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	var group int
	m := make(map[int][]float64)
	for _, val := range values {
		valInt := int64(val)
		fmt.Printf("%d ", valInt)
		group = 0
		for {
			if valInt < 0 {
				if (valInt + 10) > -10 {
					group -= 10
					fmt.Printf(" break ")
					break
				} else {
					fmt.Printf(" gr++ ")
					group -= 10
					valInt += 10
				}
			} else {
				if (valInt - 10) < 10 {
					group += 10
					fmt.Printf(" break ")
					break
				} else {
					fmt.Printf(" gr++ ")
					group += 10
					valInt -= 10
				}
			}
		}
		m[group] = append(m[group], val)
		fmt.Printf("group: %d\n", group)
	}
	fmt.Println(m)
}
