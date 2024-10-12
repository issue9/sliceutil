// SPDX-FileCopyrightText: 2024 caixw
//
// SPDX-License-Identifier: MIT

package sliceutil

import (
	"slices"
	"testing"

	"github.com/issue9/assert/v4"
)

func TestFilterSeq(t *testing.T) {
	a := assert.New(t, false)

	ints1 := []int{1, 2, 3, 4, 5}
	seq := FilterSeq(slices.Values(ints1), func(e int) bool { return e == 3 || e == 4 })
	ret := slices.Collect(FilterSeq(seq, func(e int) bool { return e == 3 }))
	a.Equal(ret, []int{3})
}
