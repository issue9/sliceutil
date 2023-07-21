// SPDX-License-Identifier: MIT

// Package sliceutil 提供对切片的相关功能
package sliceutil

// At 从 slice 中查找符合 eq 的元素
func At[T any](slice []T, eq func(T, int) bool) (T, bool) {
	if index := Index(slice, eq); index > -1 {
		return slice[index], true
	}

	var v T
	return v, false
}

// Index 从 slice 查找符合 eq 的第一个元素并返回其在数组中的元素
func Index[T any](slice []T, eq func(T, int) bool) (index int) {
	for i, e := range slice {
		if eq(e, i) {
			return i
		}
	}
	return -1
}

// Indexes 返回所有符合条件的索引
func Indexes[T any](slice []T, eq func(T, int) bool) (indexes []int) {
	indexes = make([]int, 0, 10)
	for i, e := range slice {
		if eq(e, i) {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

// Exists 判断 slice 中是否存在符合 eq 的元素存在
func Exists[T any](slice []T, eq func(T, int) bool) bool { return Index(slice, eq) > -1 }

// Reverse 反转数组中的元素
func Reverse[T any](slice []T) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Delete 删除 slice 中符合 eq 条件的元素
//
// eq 对比函数，用于确定指定的元素是否可以删除，返回 true 表示可以删除；
func Delete[T any](slice []T, eq func(T, int) bool) []T {
	l := len(slice)
	var cnt int
	last := l - 1

	for i := 0; i <= last; i++ {
		if !eq(slice[i], i) {
			continue
		}

		for j := i; j < last; j++ {
			slice[j], slice[j+1] = slice[j+1], slice[j]
		}
		cnt++
		i--
		last--
	}

	return slice[:l-cnt]
}

// QuickDelete 删除 slice 中符合 eq 条件的元素
//
// 功能与 Delete 相同，但是性能相对 Delete 会好一些，同时也不再保证元素顺序与原数组相同。
func QuickDelete[T any](slice []T, eq func(T, int) bool) []T {
	l := len(slice)
	var cnt int
	last := l - 1

	for i := 0; i <= last; i++ {
		if !eq(slice[i], i) {
			continue
		}

		slice[i], slice[last] = slice[last], slice[i]
		cnt++
		last--
		i--
	}

	return slice[:l-cnt]
}

// Count 检测数组中指定值的数量
func Count[T any](slice []T, eq func(T, int) bool) (count int) {
	for i, e := range slice {
		if eq(e, i) {
			count++
		}
	}
	return
}

// Unique 提取 slice 中的所有唯一值
//
// NOTE: 此操作会改变 slice 元素顺序。
func Unique[T any](slice []T, eq func(i, j T) bool) []T {
	var cnt int
	l := len(slice)
	last := l - 1
	for i := 0; i <= last; i++ {
		for j := i + 1; j <= last; j++ {
			if eq(slice[i], slice[j]) {
				slice[j], slice[last] = slice[last], slice[j]
				last--
				cnt++
			}
		}
	}

	return slice[:l-cnt]
}

// Dup 检测数组或是切片中是否包含重复的值
//
// 在存在相同元素时，会返回该相同元素的下标列表，
// 当有多组相同元素时，仅返回第一组相同元素的下标。
func Dup[T any](slice []T, eq func(i, j T) bool) (indexes []int) {
	l := len(slice)

	for i := 0; i < l && len(indexes) == 0; i++ {
		for j := i + 1; j < l; j++ {
			if eq(slice[i], slice[j]) {
				if len(indexes) == 0 {
					indexes = append(indexes, i)
				}
				indexes = append(indexes, j)
			}
		}
	}

	return indexes
}

// Contains container 是否包含了 sub 中的所有元素
//
// container 与 sub 都必须是数组或是切片类型。
// 如果只是需要判断某一个值是否在 container 中，可以使用 Count() 函数。
// eq 用于判断两个数组或是切的某个元素是否相等，其原型为：
//
//	func(i, j int) bool
//
// i 表示 sub 的第 i 个元素，j 表示 container 的第 j 个元素，两者顺序不能乱。
func Contains[T any](container, sub []T, eq func(i, j T) bool) bool {
	cl := len(container)
	sl := len(sub)
	if sl > cl {
		return false
	}

LOOP:
	for i := 0; i < sl; i++ {
		for j := 0; j < cl; j++ {
			if eq(sub[i], container[j]) {
				continue LOOP
			}
		}
		return false
	}
	return true
}

// Min 查找最小值
//
// less 用于判断 i 是否小于 j
func Min[T any](slices []T, less func(i, j T) bool) T {
	min := slices[0]
	for i := 1; i < len(slices); i++ {
		if !less(min, slices[i]) {
			min = slices[i]
		}
	}
	return min
}

// Max 查找最大值
//
// less 用于判断 i 是否小于 j
func Max[T any](slices []T, less func(i, j T) bool) T {
	max := slices[0]
	for i := 1; i < len(slices); i++ {
		if less(max, slices[i]) {
			max = slices[i]
		}
	}
	return max
}

// Filter 过滤数据
//
// NOTE: 这是基于对原有数据 slices 的修改。
func Filter[T any](slices []T, f func(T, int) bool) []T {
	i := 0
	for index, elem := range slices {
		if f(elem, index) {
			slices[i] = elem
			i++
		}
	}
	return slices[:i]
}

// SafeFilter 过滤数据
func SafeFilter[T any](slices []T, f func(T, int) bool) []T {
	items := make([]T, 0, len(slices))
	for i, elem := range slices {
		if f(elem, i) {
			items = append(items, elem)
		}
	}
	return items
}
