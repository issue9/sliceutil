// SPDX-FileCopyrightText: 2020-2024 caixw
//
// SPDX-License-Identifier: MIT

package sliceutil

import (
	"sort"
	"testing"

	"github.com/issue9/assert/v4"
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

func TestAt(t *testing.T) {
	a := assert.New(t, false)

	v, found := At(objSlice, func(o *obj, _ int) bool { return o.ID == 100 })
	a.False(found).Nil(v)

	v, found = At(objSlice, func(o *obj, _ int) bool { return o.ID == 2 })
	a.True(found).Equal(objSlice[1], v)

	v, found = At(objSlice, func(o *obj, _ int) bool { return o.ID == 5 })
	a.True(found).Equal(objSlice[4], v)

	v2, found := At([]string{"1", "2"}, func(o string, _ int) bool { return o == "-1" })
	a.False(found).Equal("", v2)
}

func TestIndex(t *testing.T) {
	a := assert.New(t, false)

	a.Equal(-1, Index(objSlice, func(o *obj, _ int) bool { return o.ID == 100 }))
	a.Equal(1, Index(objSlice, func(o *obj, _ int) bool { return o.ID == 2 }))
	a.Equal(4, Index(objSlice, func(o *obj, _ int) bool { return o.ID == 5 }))
}

func TestExists(t *testing.T) {
	a := assert.New(t, false)

	a.False(Exists(objSlice, func(o *obj, _ int) bool { return o.ID == 100 }))
	a.True(Exists(objSlice, func(o *obj, _ int) bool { return o.ID == 2 }))
	a.True(Exists(objSlice, func(o *obj, _ int) bool { return o.ID == 5 }))
}

func TestIndexes(t *testing.T) {
	a := assert.New(t, false)

	intSlice := []int{1, 2, 3, 7, 0, 4, 7} // 奇数个数
	indexes := Indexes(intSlice, func(v int, _ int) bool { return v == 7 })
	a.Equal(indexes, []int{3, 6})
}

func TestReverse(t *testing.T) {
	a := assert.New(t, false)

	intSlice := []int{1, 2, 3, 7, 0, 4, 7} // 奇数个数
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
	eq := func(e int, _ int) bool {
		return e == 7
	}
	copy(intInput, intSlice)
	rslt := Delete(intInput, eq)
	a.Equal(rslt, intResult)

	// quickDelete
	intResult = []int{1, 2, 3, 4, 0}
	intInput = make([]int, len(intSlice))
	eq = func(e int, _ int) bool {
		return e == 7
	}
	copy(intInput, intSlice)
	rslt = QuickDelete(intInput, eq)
	a.Equal(rslt, intResult)

	// 未找到元素

	intSlice = []int{1, 2, 3, 7, 7, 0, 4}

	// delete
	intResult = []int{1, 2, 3, 7, 7, 0, 4}
	intInput = make([]int, len(intSlice))
	eq = func(e int, _ int) bool {
		return e == -1
	}
	copy(intInput, intSlice)
	rslt = Delete(intInput, eq)
	a.Equal(rslt, intResult)

	// quickDelete
	intResult = []int{1, 2, 3, 7, 7, 0, 4}
	intInput = make([]int, len(intSlice))
	eq = func(e int, _ int) bool {
		return e == -1
	}
	copy(intInput, intSlice)
	rslt = QuickDelete(intInput, eq)
	a.Equal(rslt, intResult)

	// 连续的相同数值

	intSlice = []int{1, 2, 3, 7, 7, 0, 4}

	// delete
	intResult = []int{1, 2, 3, 0, 4}
	intInput = make([]int, len(intSlice))
	eq = func(e, _ int) bool {
		return e == 7
	}
	copy(intInput, intSlice)
	rslt = Delete(intInput, eq)
	a.Equal(rslt, intResult)

	// quickDelete
	intResult = []int{1, 2, 3, 4, 0}
	intInput = make([]int, len(intSlice))
	eq = func(e int, _ int) bool {
		return e == 7
	}
	copy(intInput, intSlice)
	rslt = QuickDelete(intInput, eq)
	a.Equal(rslt, intResult)

	// 删除后为空数组

	intSlice = []int{1, 2, 3, 7, 7, 0, 4}

	// 空数组

	intSlice = intSlice[:0]

	// delete
	eq = func(e int, _ int) bool {
		return e == 7
	}
	rslt = Delete(intSlice, eq)
	a.Equal(len(rslt), 0)

	// quickDelete
	eq = func(e int, _ int) bool {
		return e == 7
	}
	rslt = QuickDelete(intSlice, eq)
	a.Equal(len(rslt), 0)

	// nil

	intSlice = nil

	// delete
	eq = func(e int, _ int) bool {
		return e == 7
	}
	rslt = Delete(intSlice, eq)
	a.Equal(len(rslt), 0)

	// quickDelete
	eq = func(e int, _ int) bool {
		return e == 7
	}
	rslt = QuickDelete(intSlice, eq)
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
	eq1 := func(e *obj, _ int) bool {
		return e.ID == 2
	}
	copy(objInput, objSlice)
	rslt2 := Delete(objInput, eq1)
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
	eq1 = func(e *obj, _ int) bool {
		return e.ID == 2
	}
	copy(objInput, objSlice)
	rslt2 = QuickDelete(objInput, eq1)
	a.Equal(rslt2, objResult)
}

func TestUnique(t *testing.T) {
	a := assert.New(t, false)

	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	rslt := Unique(intSlice, func(i, j int) bool { return i == j })
	a.Equal(rslt, []int{1, 2, 3, 7, 0, 4})

	intSlice = []int{1, 2, 3, 7, 3, 4, 4, 0, 4, 7}
	rslt = Unique(intSlice, func(i, j int) bool { return i == j })
	a.Equal(rslt, []int{1, 2, 3, 7, 4, 0})

	// 空数组
	intSlice = []int{}
	rslt = Unique(intSlice, func(i, j int) bool { return i == j })
	a.Empty(rslt)

	// 空数组
	intSlice = nil
	rslt = Unique(intSlice, func(i, j int) bool { return i == j })
	a.Empty(rslt)
}

func TestCount(t *testing.T) {
	a := assert.New(t, false)

	intSlice := []int{1, 2, 3, 7, 0, 4, 7}
	a.Equal(2, Count(intSlice, func(e int, _ int) bool {
		return e == 7
	}))
	a.Equal(3, Count(intSlice, func(e int, _ int) bool {
		return e == 7 || e == 2
	}))
	a.Equal(1, Count(intSlice, func(e int, _ int) bool {
		return e == 0
	}))
	a.Equal(0, Count(intSlice, func(e int, _ int) bool {
		return e == -1000
	}))

	// 空数组
	intSlice = []int{}
	a.Equal(0, Count(intSlice, func(e int, _ int) bool {
		return e == 7
	}))

	// 空数组

	a.Equal(2, Count(objSlice, func(e *obj, _ int) bool {
		return e.ID == 1
	}))
	a.Equal(1, Count(objSlice, func(e *obj, _ int) bool {
		return e.Name == "4"
	}))
	a.Equal(0, Count(objSlice, func(e *obj, _ int) bool {
		return e.Age == 1000
	}))
}

func TestDup(t *testing.T) {
	a := assert.New(t, false)

	intSlice := []int{1, 2, 3, 7, 0, 4, 7, 7, 5, 7}
	a.Equal([]int{3, 6, 7, 9}, Dup(intSlice, func(i, j int) bool {
		return i == j
	}))

	// 空数组
	intSlice = []int{}
	a.Nil(Dup(intSlice, func(i, j int) bool {
		return i == j
	}))

	// 空数组
	intSlice = nil
	a.Nil(Dup(intSlice, func(i, j int) bool {
		return i == j
	}))

	stringSlice := []string{"a", "b", "0", "a"}
	a.Equal([]int{0, 3}, Dup(stringSlice, func(i, j string) bool {
		return i == j
	}))

	a.Equal([]int{0, 5}, Dup(objSlice, func(i, j *obj) bool {
		return i.ID == j.ID
	}))
	a.Equal([]int{0, 4}, Dup(objSlice, func(i, j *obj) bool {
		return i.Name == j.Name
	}))
	a.Nil(Dup(objSlice, func(i, j *obj) bool {
		return i.Age == j.Age
	}))
}

func TestContains(t *testing.T) {
	a := assert.New(t, false)

	ints1 := []int{1, 2, 3, 4, 5}
	ints2 := []int{1, 5, 2}
	ints3 := []int{1, 9, 7}
	a.True(Contains(ints1, ints2, func(i, j int) bool { return i == j }))
	a.False(Contains(ints2, ints1, func(i, j int) bool { return i == j }))
	a.False(Contains(ints1, ints3, func(i, j int) bool { return i == j }))

	// object
	objArr := []*obj{
		{ID: 2, Name: "4", Age: 2},
		{ID: 3, Name: "3", Age: 3},
		{ID: 5, Name: "5", Age: 5},
	}
	a.True(Contains(objSlice, objArr, func(i, j *obj) bool { return i.ID == j.ID }))
}

func TestMinMax(t *testing.T) {
	a := assert.New(t, false)

	ints1 := []int{1, 2, 3, 4, 5}
	a.Equal(Min(ints1, func(i, j int) bool { return i < j }), 1)
	a.Equal(Max(ints1, func(i, j int) bool { return i < j }), 5)

	ints1 = []int{1, 2, 0, 0, 7, 4, 5}
	a.Equal(Min(ints1, func(i, j int) bool { return i < j }), 0)
	a.Equal(Max(ints1, func(i, j int) bool { return i < j }), 7)

	ints1 = []int{1, 2, -9, 0, 7, 4, 5}
	a.Equal(Min(ints1, func(i, j int) bool { return i < j }), -9)
	a.Equal(Max(ints1, func(i, j int) bool { return i < j }), 7)
}

func TestFilter(t *testing.T) {
	a := assert.New(t, false)

	ints1 := []int{1, 2, 3, 4, 5}
	ret := Filter(ints1, func(e, _ int) bool { return e == 3 || e == 4 })
	a.Equal(ret, []int{3, 4}).
		Equal(ints1[:2], []int{3, 4})
}

func TestSafeFilter(t *testing.T) {
	a := assert.New(t, false)

	ints1 := []int{1, 2, 3, 4, 5}
	ret := SafeFilter(ints1, func(e, _ int) bool { return e == 3 || e == 4 })
	a.Equal(ret, []int{3, 4}).
		Equal(ints1[:2], []int{1, 2})
}

func TestMapKeys(t *testing.T) {
	a := assert.New(t, false)
	m := map[string]string{"1": "1", "2": "2", "0": "0"}
	keys := MapKeys(m)
	sort.Strings(keys)
	a.Equal(keys, []string{"0", "1", "2"})
}

func TestMapVals(t *testing.T) {
	a := assert.New(t, false)
	m := map[string]string{"1": "1", "2": "2", "0": "0"}
	vals := MapVals(m)
	sort.Strings(vals)
	a.Equal(vals, []string{"0", "1", "2"})
}
