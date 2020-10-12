// SPDX-License-Identifier: MIT

package sliceutil

import "fmt"

func ExampleDup() {
	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	fmt.Println(Dup(intSlice, func(i, j int) bool {
		return intSlice[i] == intSlice[j]
	}))

	// Output: [3 6]
}

func ExampleCount() {
	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	fmt.Println(Count(intSlice, func(i int) bool {
		return intSlice[i] == 7
	}))

	// Output: 2
}

func ExampleDelete() {
	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	size := Delete(intSlice, func(i int) bool {
		return intSlice[i] == 7
	})
	fmt.Println("Delete:", intSlice[:size])

	intSlice = []int{1, 2, 3, 7, 0, 4, 7}
	size = QuickDelete(intSlice, func(i int) bool {
		return intSlice[i] == 7 || intSlice[i] == 2
	})
	fmt.Println("QuickDelete:", intSlice[:size])

	// Output: Delete: [1 2 3 0 4]
	// QuickDelete: [1 4 3 0]
}

func ExampleContains() {
	ints := []int{1, 2, 3, 4, 5}
	uints := []uint{1, 5, 2}
	fmt.Println(Contains(ints, uints, func(i, j int) bool {
		return uint(ints[i]) == uints[j]
	}))

	// Output: true
}
