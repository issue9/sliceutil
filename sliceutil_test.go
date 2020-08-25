// SPDX-License-Identifier: MIT

package sliceutil

import (
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

func TestReverse(t *testing.T) {
	a := assert.New(t)

	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	Reverse(intSlice)
	a.Equal(intSlice, []int{7, 4, 0, 7, 3, 2, 1})

	byteSlice := []byte{1, 2, 3, 7, 0, 4, 7}
	Reverse(byteSlice)
	a.Equal(byteSlice, []int{7, 4, 0, 7, 3, 2, 1})
}

func TestDelete(t *testing.T) {
	a := assert.New(t)

	intSlice := []int{1, 2, 3, 7, 0, 4, 7}

	// delete
	intResult := []int{1, 2, 3, 0, 4}
	intInput := make([]int, len(intSlice))
	eq := func(i int) bool {
		return intInput[i] == 7
	}
	copy(intInput, intSlice)
	size := Delete(intInput, eq)
	a.Equal(intInput[:size], intResult)

	// quickDelete
	intResult = []int{1, 2, 3, 4, 0}
	intInput = make([]int, len(intSlice))
	eq = func(i int) bool {
		return intInput[i] == 7
	}
	copy(intInput, intSlice)
	size = QuickDelete(intInput, eq)
	a.Equal(intInput[:size], intResult)

	// 连续的相同数值

	intSlice = []int{1, 2, 3, 7, 7, 0, 4}

	// delete
	intResult = []int{1, 2, 3, 0, 4}
	intInput = make([]int, len(intSlice))
	eq = func(i int) bool {
		return intInput[i] == 7
	}
	copy(intInput, intSlice)
	size = Delete(intInput, eq)
	a.Equal(intInput[:size], intResult)

	// quickDelete
	intResult = []int{1, 2, 3, 4, 0}
	intInput = make([]int, len(intSlice))
	eq = func(i int) bool {
		return intInput[i] == 7
	}
	copy(intInput, intSlice)
	size = QuickDelete(intInput, eq)
	a.Equal(intInput[:size], intResult)

	// 删除后为空数组

	intSlice = []int{1, 2, 3, 7, 7, 0, 4}

	// delete
	intInput = make([]int, len(intSlice))
	eq = func(i int) bool {
		return true
	}
	copy(intInput, intSlice)
	size = Delete(&intInput, eq)
	a.Equal(0, size)

	// quickDelete
	intInput = make([]int, len(intSlice))
	eq = func(i int) bool {
		return true
	}
	copy(intInput, intSlice)
	size = QuickDelete(&intInput, eq)
	a.Equal(0, size)

	// 空数组

	intSlice = intSlice[:0]

	// delete
	eq = func(i int) bool {
		return intSlice[i] == 7
	}
	size = Delete(intSlice, eq)
	a.Equal(size, 0)

	// quickDelete
	eq = func(i int) bool {
		return intSlice[i] == 7
	}
	size = QuickDelete(intSlice, eq)
	a.Equal(size, 0)

	// nil

	intSlice = nil

	// delete
	eq = func(i int) bool {
		return intSlice[i] == 7
	}
	size = Delete(intSlice, eq)
	a.Equal(size, 0)

	// quickDelete
	eq = func(i int) bool {
		return intSlice[i] == 7
	}
	size = QuickDelete(intSlice, eq)
	a.Equal(size, 0)

	// array

	intArray := [7]int{1, 2, 3, 7, 7, 0, 4}
	a.Panic(func() {
		Delete(intArray, eq)
	})
	a.Panic(func() {
		QuickDelete(intArray, eq)
	})

	// object slice

	// delete
	objResult := []*obj{
		{ID: 1, Name: "5", Age: 1},
		{ID: 3, Name: "3", Age: 3},
		{ID: 4, Name: "2", Age: 4},
		{ID: 5, Name: "5", Age: 5},
		{ID: 1, Name: "1", Age: 6},
	}
	objInput := make([]*obj, len(objSlice))
	eq = func(i int) bool {
		return objInput[i].ID == 2
	}
	copy(objInput, objSlice)
	size = Delete(objInput, eq)
	a.Equal(objInput[:size], objResult)

	// quickDelete
	objResult = []*obj{
		{ID: 1, Name: "5", Age: 1},
		{ID: 1, Name: "1", Age: 6},
		{ID: 3, Name: "3", Age: 3},
		{ID: 4, Name: "2", Age: 4},
		{ID: 5, Name: "5", Age: 5},
	}
	objInput = make([]*obj, len(objSlice))
	eq = func(i int) bool {
		return objInput[i].ID == 2
	}
	copy(objInput, objSlice)
	size = QuickDelete(objInput, eq)
	a.Equal(objInput[:size], objResult)
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

	intSlice := []int{1, 2, 3, 7, 0, 4, 7, 7, 5, 7}
	a.Equal([]int{3, 6, 7, 9}, Dup(intSlice, func(i, j int) bool {
		return intSlice[i] == intSlice[j]
	}))

	// 空数组
	intSlice = []int{}
	a.Nil(Dup(intSlice, func(i, j int) bool {
		return intSlice[i] == intSlice[j]
	}))

	// 空数组
	intSlice = nil
	a.Nil(Dup(intSlice, func(i, j int) bool {
		return intSlice[i] == intSlice[j]
	}))

	intArray := [7]int{1, 2, 3, 7, 0, 4, 7}
	a.Equal([]int{3, 6}, Dup(intArray, func(i, j int) bool {
		return intArray[i] == intArray[j]
	}))

	a.Equal([]int{3, 6}, Dup(&intArray, func(i, j int) bool {
		return intArray[i] == intArray[j]
	}))

	stringSlice := []string{"a", "b", "0", "a"}
	a.Equal([]int{0, 3}, Dup(stringSlice, func(i, j int) bool {
		return stringSlice[i] == stringSlice[j]
	}))

	a.Equal([]int{0, 5}, Dup(objSlice, func(i, j int) bool {
		return objSlice[i].ID == objSlice[j].ID
	}))
	a.Equal([]int{0, 4}, Dup(objSlice, func(i, j int) bool {
		return objSlice[i].Name == objSlice[j].Name
	}))
	a.Nil(Dup(objSlice, func(i, j int) bool {
		return objSlice[i].Age == objSlice[j].Age
	}))

	a.Panic(func() {
		Dup(5, func(i, j int) bool {
			return false
		})
	})
}

func TestContains(t *testing.T) {
	a := assert.New(t)

	ints := []int{1, 2, 3, 4, 5}
	uints := []uint{1, 5, 2}
	int8s := []int8{1, 9, 7}
	floats := []float32{1.0, 9.0}
	a.True(Contains(ints, uints, func(i, j int) bool { return uint(ints[i]) == uints[j] }))
	a.False(Contains(uints, ints, func(i, j int) bool { return int(uints[i]) == ints[j] }))
	a.False(Contains(ints, int8s, func(i, j int) bool { return int8(ints[i]) == int8s[j] }))
	a.True(Contains(int8s, floats, func(i, j int) bool { return float32(int8s[i]) == floats[j] }))

	// arr vs slice
	int8Arr := [3]int8{1, 3, 5}
	a.True(Contains(ints, int8Arr, func(i, j int) bool { return int8(ints[i]) == int8Arr[j] }))

	// object
	objArr := [3]*obj{
		{ID: 2, Name: "4", Age: 2},
		{ID: 3, Name: "3", Age: 3},
		{ID: 5, Name: "5", Age: 5},
	}
	a.True(Contains(objSlice, objArr, func(i, j int) bool { return objSlice[i].ID == objArr[j].ID }))
}
