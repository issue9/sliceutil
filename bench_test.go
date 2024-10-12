// SPDX-FileCopyrightText: 2020-2024 caixw
//
// SPDX-License-Identifier: MIT

package sliceutil

import "testing"

func BenchmakIndexes(b *testing.B) {
	intSlice := []int{1, 2, 3, 7, 0, 4, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 8, 7}
	eq := func(i, _ int) bool {
		return intSlice[i] == 7
	}

	for i := 0; i < b.N; i++ {
		Indexes(intSlice, eq)
	}
}

func BenchmarkDelete(b *testing.B) {
	b.Run("delete", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			intSlice := []int{1, 2, 3, 7, 0, 4, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 8, 7}
			eq := func(i, _ int) bool {
				return intSlice[i] == 7
			}
			Delete(intSlice, eq)
		}
	})

	b.Run("quickDelete", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			intSlice := []int{1, 2, 3, 7, 0, 4, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 8, 7}
			eq := func(i, _ int) bool {
				return intSlice[i] == 7
			}
			QuickDelete(intSlice, eq)
		}
	})
}

func BenchmarkFilter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intSlice := []int{1, 2, 3, 7, 0, 4, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 8, 7}
		eq := func(i, _ int) bool {
			return intSlice[i] == 7
		}
		Filter(intSlice, eq)
	}
}

func BenchmarkSafeFilter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		intSlice := []int{1, 2, 3, 7, 0, 4, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 8, 7}
		eq := func(i, _ int) bool {
			return intSlice[i] == 7
		}
		SafeFilter(intSlice, eq)
	}
}
