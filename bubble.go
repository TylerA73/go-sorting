package main

import "fmt"

func bubbleSort(a []int) {
	sorted := false // Start sorted as false
	fmt.Println(a)  // Display the initial stage of the array

	// While not sorted, continue to operate on the array
	// When sorted is true, the array is sorted
	for !sorted {

		// Assume that when we start that the array is sorted
		sorted = true

		// Loop through the entire array, comparing a value at position i to the next neighbour
		for i := 0; i < len(a)-1; i++ {

			// If we make a swap, the array is not sorted
			if a[i] > a[i+1] {
				sorted = false
				tmp := a[i+1]
				a[i+1] = a[i]
				a[i] = tmp
			}

		}

		// Display each iteration of the bubble sort
		fmt.Println(a)
	}
}
