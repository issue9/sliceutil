// SPDX-License-Identifier: MIT

package sliceutil

import (
	"testing"

	"github.com/issue9/assert/v2"
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

func TestIndex(t *testing.T) {
	a := assert.New(t, false)

	a.Equal(-1, Index[*obj](objSlice, func(o *obj) bool { return o.ID == 100 }))
	a.Equal(1, Index[*obj](objSlice, func(o *obj) bool { return o.ID == 2 }))
	a.Equal(4, Index[*obj](objSlice, func(o *obj) bool { return o.ID == 5 }))
}

func TestReverse(t *testing.T) {
	a := assert.New(t, false)

	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	Reverse(intSlice)
	a.Equal(intSlice, []int{7, 4, 0, 7, 3, 2, 1})

	byteSlice := []byte{1, 2, 3, 7, 0, 4, 7}
	Reverse(byteSlice)
	a.Equal(byteSlice, []int{7, 4, 0, 7, 3, 2, 1})
}

func TestDelete(t *testing.T) {
	a := assert.New(t, false)

	intSlice := []int{1, 2, 3, 7, 0, 4, 7}

	// delete
	intResult := []int{1, 2, 3, 0, 4}
	intInput := make([]int, len(intSlice))
	eq := func(e int) bool {
		return e == 7
	}
	copy(intInput, intSlice)
	rslt := Delete[int](intInput, eq)
	a.Equal(rslt, intResult)

	// quickDelete
	intResult = []int{1, 2, 3, 4, 0}
	intInput = make([]int, len(intSlice))
	eq = func(e int) bool {
		return e == 7
	}
	copy(intInput, intSlice)
	rslt = QuickDelete[int](intInput, eq)
	a.Equal(rslt, intResult)

	// 未找到元素

	intSlice = []int{1, 2, 3, 7, 7, 0, 4}

	// delete
	intResult = []int{1, 2, 3, 7, 7, 0, 4}
	intInput = make([]int, len(intSlice))
	eq = func(e int) bool {
		return e == -1
	}
	copy(intInput, intSlice)
	rslt = Delete[int](intInput, eq)
	a.Equal(rslt, intResult)

	// quickDelete
	intResult = []int{1, 2, 3, 7, 7, 0, 4}
	intInput = make([]int, len(intSlice))
	eq = func(e int) bool {
		return e == -1
	}
	copy(intInput, intSlice)
	rslt = QuickDelete[int](intInput, eq)
	a.Equal(rslt, intResult)

	// 连续的相同数值

	intSlice = []int{1, 2, 3, 7, 7, 0, 4}

	// delete
	intResult = []int{1, 2, 3, 0, 4}
	intInput = make([]int, len(intSlice))
	eq = func(e int) bool {
		return e == 7
	}
	copy(intInput, intSlice)
	rslt = Delete[int](intInput, eq)
	a.Equal(rslt, intResult)

	// quickDelete
	intResult = []int{1, 2, 3, 4, 0}
	intInput = make([]int, len(intSlice))
	eq = func(e int) bool {
		return e == 7
	}
	copy(intInput, intSlice)
	rslt = QuickDelete[int](intInput, eq)
	a.Equal(rslt, intResult)

	// 删除后为空数组

	intSlice = []int{1, 2, 3, 7, 7, 0, 4}

	// 空数组

	intSlice = intSlice[:0]

	// delete
	eq = func(e int) bool {
		return e == 7
	}
	rslt = Delete[int](intSlice, eq)
	a.Equal(len(rslt), 0)

	// quickDelete
	eq = func(e int) bool {
		return e == 7
	}
	rslt = QuickDelete[int](intSlice, eq)
	a.Equal(len(rslt), 0)

	// nil

	intSlice = nil

	// delete
	eq = func(e int) bool {
		return e == 7
	}
	rslt = Delete[int](intSlice, eq)
	a.Equal(len(rslt), 0)

	// quickDelete
	eq = func(e int) bool {
		return e == 7
	}
	rslt = QuickDelete[int](intSlice, eq)
	a.Equal(len(rslt), 0)

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
	eq1 := func(e *obj) bool {
		return e.ID == 2
	}
	copy(objInput, objSlice)
	rslt2 := Delete[*obj](objInput, eq1)
	a.Equal(rslt2, objResult)

	// quickDelete
	objResult = []*obj{
		{ID: 1, Name: "5", Age: 1},
		{ID: 1, Name: "1", Age: 6},
		{ID: 3, Name: "3", Age: 3},
		{ID: 4, Name: "2", Age: 4},
		{ID: 5, Name: "5", Age: 5},
	}
	objInput = make([]*obj, len(objSlice))
	eq1 = func(e *obj) bool {
		return e.ID == 2
	}
	copy(objInput, objSlice)
	rslt2 = QuickDelete[*obj](objInput, eq1)
	a.Equal(rslt2, objResult)
}

func TestUnique(t *testing.T) {
	a := assert.New(t, false)

	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	rslt := Unique[int](intSlice, func(i, j int) bool { return i == j })
	a.Equal(rslt, []int{1, 2, 3, 7, 0, 4})

	intSlice = []int{1, 2, 3, 7, 3, 4, 4, 0, 4, 7}
	rslt = Unique[int](intSlice, func(i, j int) bool { return i == j })
	a.Equal(rslt, []int{1, 2, 3, 7, 4, 0})

	// 空数组
	intSlice = []int{}
	rslt = Unique[int](intSlice, func(i, j int) bool { return i == j })
	a.Empty(rslt)

	// 空数组
	intSlice = nil
	rslt = Unique[int](intSlice, func(i, j int) bool { return i == j })
	a.Empty(rslt)
}

func TestCount(t *testing.T) {
	a := assert.New(t, false)

	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	a.Equal(2, Count[int](intSlice, func(e int) bool {
		return e == 7
	}))
	a.Equal(3, Count[int](intSlice, func(e int) bool {
		return e == 7 || e == 2
	}))
	a.Equal(1, Count[int](intSlice, func(e int) bool {
		return e == 0
	}))
	a.Equal(0, Count[int](intSlice, func(e int) bool {
		return e == -1000
	}))

	// 空数组
	intSlice = []int{}
	a.Equal(0, Count[int](intSlice, func(e int) bool {
		return e == 7
	}))

	// 空数组

	a.Equal(2, Count[*obj](objSlice, func(e *obj) bool {
		return e.ID == 1
	}))
	a.Equal(1, Count[*obj](objSlice, func(e *obj) bool {
		return e.Name == "4"
	}))
	a.Equal(0, Count[*obj](objSlice, func(e *obj) bool {
		return e.Age == 1000
	}))
}

func TestDup(t *testing.T) {
	a := assert.New(t, false)

	intSlice := []int{1, 2, 3, 7, 0, 4, 7, 7, 5, 7}
	a.Equal([]int{3, 6, 7, 9}, Dup[int](intSlice, func(i, j int) bool {
		return i == j
	}))

	// 空数组
	intSlice = []int{}
	a.Nil(Dup[int](intSlice, func(i, j int) bool {
		return i == j
	}))

	// 空数组
	intSlice = nil
	a.Nil(Dup[int](intSlice, func(i, j int) bool {
		return i == j
	}))

	stringSlice := []string{"a", "b", "0", "a"}
	a.Equal([]int{0, 3}, Dup[string](stringSlice, func(i, j string) bool {
		return i == j
	}))

	a.Equal([]int{0, 5}, Dup[*obj](objSlice, func(i, j *obj) bool {
		return i.ID == j.ID
	}))
	a.Equal([]int{0, 4}, Dup[*obj](objSlice, func(i, j *obj) bool {
		return i.Name == j.Name
	}))
	a.Nil(Dup[*obj](objSlice, func(i, j *obj) bool {
		return i.Age == j.Age
	}))
}

func TestContains(t *testing.T) {
	a := assert.New(t, false)

	ints1 := []int{1, 2, 3, 4, 5}
	ints2 := []int{1, 5, 2}
	ints3 := []int{1, 9, 7}
	a.True(Contains[int](ints1, ints2, func(i, j int) bool { return i == j }))
	a.False(Contains[int](ints2, ints1, func(i, j int) bool { return i == j }))
	a.False(Contains[int](ints1, ints3, func(i, j int) bool { return i == j }))

	// object
	objArr := []*obj{
		{ID: 2, Name: "4", Age: 2},
		{ID: 3, Name: "3", Age: 3},
		{ID: 5, Name: "5", Age: 5},
	}
	a.True(Contains[*obj](objSlice, objArr, func(i, j *obj) bool { return i.ID == j.ID }))
}
