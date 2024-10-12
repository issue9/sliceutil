// SPDX-FileCopyrightText: 2020-2024 caixw
//
// SPDX-License-Identifier: MIT

package sliceutil

import (
	"slices"
	"testing"
)

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
	intSlice := []int{1, 2, 3, 7, 0, 4, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 8, 7}
	eq := func(i, _ int) bool {
		return intSlice[i] == 7
	}

	for i := 0; i < b.N; i++ {
		Filter(intSlice, eq)
	}
}

func BenchmarkSafeFilter(b *testing.B) {
	intSlice := []int{1, 2, 3, 7, 0, 4, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 8, 7}
	eq := func(i, _ int) bool {
		return intSlice[i] == 7
	}

	for i := 0; i < b.N; i++ {
		SafeFilter(intSlice, eq)
	}
}

func BenchmarkFilterSeq(b *testing.B) {
	intSlice := []int{1, 2, 3, 7, 0, 4, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 8, 7}
	eq := func(i int) bool {
		return i == 7
	}

	b.Run("filterSeq", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			FilterSeq(slices.Values(intSlice), eq)
		}
	})

	b.Run("collect", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			slices.Collect(FilterSeq(slices.Values(intSlice), eq))
		}
	})
}
