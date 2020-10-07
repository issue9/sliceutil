sliceutil
[![Build Status](https://github.com/issue9/sliceutil/workflows/Go/badge.svg)](https://github.com/issue9/sliceutil/actions?query=workflow%3AGo)
[![license](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat)](https://opensource.org/licenses/MIT)
[![codecov](https://codecov.io/gh/issue9/sliceutil/branch/master/graph/badge.svg)](https://codecov.io/gh/issue9/sliceutil)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/issue9/sliceutil)](https://pkg.go.dev/github.com/issue9/sliceutil)
======

sliceutil 提供了针对数组和切片的功能

- Reverse 反转数组中的元素；
- Delete 删除符合条件的切片元素；
- QuickDelete 删除符合条件的切片元素，性能稍高于 Delete；
- Count 统计数组或切片中包含指定什的数量；
- Dup 查看数组或切片中是否包含重得的值；
- Contains 判断一个数组或是切片是否包含了另一个的所有元素；

```go
intSlice := []int{1, 2, 3, 7, 0, 4, 7}
intArr := [3]int{1, 7, 0}

// index == [3, 7]
index := Dup(intSlice, func(i, j int) bool {
    return intSlice[i] == intSlice[j]
})

// 返回 7 的数量
count := Count(intSlice, func(i int) bool {
    return intSlice[i] == 7
})

// 会重新调整切片的内容，将删除后的数据在放最前端，并返回数切片的大小。
// 通过 intSlice[:size] 即为删除后的内容
size := Delete(intSlice, func(i int) bool {
    return intSlice[i] == 7
})

// ok == true
ok := Contains(intSlice, intArr, func(i, j int) bool {
    return int8(intSlice[i]) == int8Arr[j]
})
```

安装
----

```shell
go get github.com/issue9/sliceutil
```

版权
----

本项目采用 [MIT](http://opensource.org/licenses/MIT) 开源授权许可证，完整的授权说明可在 [LICENSE](LICENSE) 文件中找到。
