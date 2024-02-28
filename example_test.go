// SPDX-FileCopyrightText: 2020-2024 caixw
//
// SPDX-License-Identifier: MIT

package sliceutil

import "fmt"

func ExampleAt() {
	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	v, found := At(intSlice, func(e, _ int) bool {
		return e == 7
	})
	fmt.Println(found, v)

	// Output: true 7
}

func ExampleIndex() {
	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	fmt.Println(Index(intSlice, func(e, _ int) bool {
		return e == 7
	}))

	// Output: 3
}

func ExampleExists() {
	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	fmt.Println(Exists(intSlice, func(e, _ int) bool {
		return e == 7
	}))

	// Output: true
}

func ExampleDup() {
	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	fmt.Println(Dup(intSlice, func(i, j int) bool {
		return i == j
	}))

	// Output: [3 6]
}

func ExampleCount() {
	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	fmt.Println(Count(intSlice, func(e, _ int) bool {
		return e == 7
	}))

	// Output: 2
}

func ExampleDelete() {
	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	rslt := Delete(intSlice, func(e, _ int) bool {
		return e == 7
	})
	fmt.Println("Delete:", rslt)

	intSlice = []int{1, 2, 3, 7, 0, 4, 7}
	rslt = QuickDelete(intSlice, func(e, _ int) bool {
		return e == 7 || e == 2
	})
	fmt.Println("QuickDelete:", rslt)

	// Output: Delete: [1 2 3 0 4]
	// QuickDelete: [1 4 3 0]
}

func ExampleUnique() {
	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	rslt := Unique(intSlice, func(i, j int) bool {
		return i == j
	})
	fmt.Println(rslt)

	// Output: [1 2 3 7 0 4]
}

func ExampleContains() {
	ints1 := []int{1, 2, 3, 4, 5}
	ints2 := []int{1, 5, 2}
	fmt.Println(Contains(ints1, ints2, func(i, j int) bool {
		return i == j
	}))

	// Output: true
}
