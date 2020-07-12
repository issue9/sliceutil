// SPDX-License-Identifier: MIT

package sliceutil

import "testing"

func BenchmarkDelete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intSlice := []int{1, 2, 3, 7, 0, 4, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 8, 7}
		eq := func(i int) bool {
			return intSlice[i] == 7
		}
		Delete(intSlice, eq)
	}
}

func BenchmarkQuickDelete(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intSlice := []int{1, 2, 3, 7, 0, 4, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 8, 7}
		eq := func(i int) bool {
			return intSlice[i] == 7
		}
		QuickDelete(intSlice, eq)
	}
}
