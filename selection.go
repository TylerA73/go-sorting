package main

import "fmt"

func selectionSort(a []int) {

	var index int // Index of the smallest value
	var tmp int   // Temporary variable used for swapping

	// Print the initial array
	fmt.Println(a)

	// Loop through the entire array once
	for i := 0; i < len(a)-1; i++ {
		tmp = a[i] // The temporary variable becomes the next value we are comparing initially
		index = i  // index becomes the index of that temporary value

		// Loop through the array starting from 1 position in front of our i index to compare all values to it
		// Everything behind the i index is sorted
		for j := i + 1; j < len(a); j++ {

			// if tmp is larger than the current compared value in the array, then we have a new smallest value
			if tmp > a[j] {
				tmp = a[j]
				index = j
			}
		}

		// Swap the current value at i with the next smallest value
		// If the smallest value is already in this location, the swap will just put it back where it was
		a[index] = a[i]
		a[i] = tmp

		// Print the array at each step to show that it works like selection sort should
		fmt.Println(a)
	}
}
