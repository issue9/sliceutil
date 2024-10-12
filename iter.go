// SPDX-FileCopyrightText: 2024 caixw
//
// SPDX-License-Identifier: MIT

package sliceutil

import "iter"

// Filter 过滤数据
func FilterSeq[T any](seq iter.Seq[T], f func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for elem := range seq {
			if f(elem) {
				if !yield(elem) {
					break
				}
			}
		}
	}
}
