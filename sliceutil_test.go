// SPDX-License-Identifier: MIT

package sliceutil

import (
	"fmt"
	"testing"

	"github.com/issue9/assert"
)

type obj struct {
	ID   int
	Name string
	Age  int
}

var objSlice = []*obj{
	{ID: 1, Name: "5", Age: 1},
	{ID: 2, Name: "4", Age: 2},
	{ID: 3, Name: "3", Age: 3},
	{ID: 4, Name: "2", Age: 4},
	{ID: 5, Name: "5", Age: 5},
	{ID: 1, Name: "1", Age: 6},
}

func TestCount(t *testing.T) {
	a := assert.New(t)

	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	a.Equal(2, Count(intSlice, func(i int) bool {
		return intSlice[i] == 7
	}))
	a.Equal(1, Count(intSlice, func(i int) bool {
		return intSlice[i] == 0
	}))
	a.Equal(0, Count(intSlice, func(i int) bool {
		return intSlice[i] == -1000
	}))

	// 空数组
	intSlice = []int{}
	a.Equal(0, Count(intSlice, func(i int) bool {
		return intSlice[i] == 7
	}))

	// 空数组
	intSlice = nil
	a.Equal(0, Count(&intSlice, func(i int) bool {
		return intSlice[i] == 0
	}))

	a.Equal(2, Count(objSlice, func(i int) bool {
		return objSlice[i].ID == 1
	}))
	a.Equal(1, Count(objSlice, func(i int) bool {
		return objSlice[i].Name == "4"
	}))
	a.Equal(0, Count(objSlice, func(i int) bool {
		return objSlice[i].Age == 1000
	}))
}

func TestDup(t *testing.T) {
	a := assert.New(t)

	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	a.Equal(6, Dup(intSlice, func(i, j int) bool {
		return intSlice[i] == intSlice[j]
	}))

	// 空数组
	intSlice = []int{}
	a.Equal(-1, Dup(intSlice, func(i, j int) bool {
		return intSlice[i] == intSlice[j]
	}))

	// 空数组
	intSlice = nil
	a.Equal(-1, Dup(intSlice, func(i, j int) bool {
		return intSlice[i] == intSlice[j]
	}))

	intArray := [7]int{1, 2, 3, 7, 0, 4, 7}
	a.Equal(6, Dup(intArray, func(i, j int) bool {
		return intArray[i] == intArray[j]
	}))

	a.Equal(6, Dup(&intArray, func(i, j int) bool {
		return intArray[i] == intArray[j]
	}))

	stringSlice := []string{"a", "b", "0", "a"}
	a.Equal(3, Dup(stringSlice, func(i, j int) bool {
		return stringSlice[i] == stringSlice[j]
	}))

	a.Equal(5, Dup(objSlice, func(i, j int) bool {
		return objSlice[i].ID == objSlice[j].ID
	}))
	a.Equal(4, Dup(objSlice, func(i, j int) bool {
		return objSlice[i].Name == objSlice[j].Name
	}))
	a.Equal(-1, Dup(objSlice, func(i, j int) bool {
		return objSlice[i].Age == objSlice[j].Age
	}))

	a.Panic(func() {
		Dup(5, func(i, j int) bool {
			return false
		})
	})
}

func ExampleDup() {
	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	fmt.Println(Dup(intSlice, func(i, j int) bool {
		return intSlice[i] == intSlice[j]
	}))

	// Output: 6
}
