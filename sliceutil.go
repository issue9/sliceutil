// SPDX-License-Identifier: MIT

// Package sliceutil 提供对数组和切片的相关功能
package sliceutil

import (
	"fmt"
	"reflect"
)

// Count 检测数组中指定值的数量
//
// slice 需要检测的数组或是切片，其它类型会 panic；
// eq 对比函数，i 表示数组的下标，需要在函数将该下标表示的值与你需要的值进行比较是否相等；
func Count(slice interface{}, eq func(i int) bool) (count int) {
	v := getSliceValue(slice)
	l := v.Len()

	for i := 0; i < l; i++ {
		if eq(i) {
			count++
		}
	}
	return
}

// Dup 检测数组中是否包含重复的值
//
// slice 需要检测的数组或是切片，其它类型会 panic；
// eq 对比数组中两个值是否相等，相等需要返回 true；
// 返回值表示存在相等值时，第二个值在数组中的下标值；
func Dup(slice interface{}, eq func(i, j int) bool) int {
	v := getSliceValue(slice)
	l := v.Len()
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if eq(i, j) {
				return j
			}
		}
	}

	return -1
}

func getSliceValue(slice interface{}) reflect.Value {
	v := reflect.ValueOf(slice)
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		panic(fmt.Sprint("参数 slice 只能是 slice 或是 array"))
	}

	return v
}
