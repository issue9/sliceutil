// SPDX-License-Identifier: MIT

package sliceutil

import "fmt"

func ExampleIndex() {
	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	fmt.Println(Index[int](intSlice, func(e int) bool {
		return e == 7
	}))

	// Output: 3
}

func ExampleDup() {
	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	fmt.Println(Dup[int](intSlice, func(i, j int) bool {
		return i == j
	}))

	// Output: [3 6]
}

func ExampleCount() {
	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	fmt.Println(Count[int](intSlice, func(e int) bool {
		return e == 7
	}))

	// Output: 2
}

func ExampleDelete() {
	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	rslt := Delete[int](intSlice, func(e int) bool {
		return e == 7
	})
	fmt.Println("Delete:", rslt)

	intSlice = []int{1, 2, 3, 7, 0, 4, 7}
	rslt = QuickDelete[int](intSlice, func(e int) bool {
		return e == 7 || e == 2
	})
	fmt.Println("QuickDelete:", rslt)

	// Output: Delete: [1 2 3 0 4]
	// QuickDelete: [1 4 3 0]
}

func ExampleUnique() {
	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	rslt := Unique[int](intSlice, func(i, j int) bool {
		return i == j
	})
	fmt.Println(rslt)

	// Output: [1 2 3 7 0 4]
}

func ExampleContains() {
	ints1 := []int{1, 2, 3, 4, 5}
	ints2 := []int{1, 5, 2}
	fmt.Println(Contains[int](ints1, ints2, func(i, j int) bool {
		return i == j
	}))

	// Output: true
}
