package main

import "fmt"

func selectionSort(a []int) {
	var index int
	var tmp int

	fmt.Println(a)

	for i := 0; i < len(a)-1; i++ {
		tmp = a[i]
		index = i

		for j := i + 1; j < len(a); j++ {
			if tmp > a[j] {
				tmp = a[j]
				index = j
			}
		}

		a[index] = a[i]
		a[i] = tmp

		fmt.Println(a)
	}
}
