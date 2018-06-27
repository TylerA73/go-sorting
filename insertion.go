package main

import "fmt"

func insertionSort(a []int) {

	// Print out the initial array
	fmt.Println(a)

	// We start at one position ahead of the front of the array
	// This is because insertion sort works backwards, and we have nothing before the front to compare the front to
	for i := 1; i < len(a); i++ {
		tmp := a[i] // tmp becomes the next value at position i for comparison

		// While j >= 0 and the value at index j is greater than tmp, continue moving the values forward one spot
		for j = i - 1; j >= 0 && a[j] > tmp; j-- {
			a[j+1] = a[j]
		}

		a[j+1] = tmp

		// Print out the next step to show that insertion is working as it should
		fmt.Println(a)
	}

}
