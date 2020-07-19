sliceutil
[![Build Status](https://img.shields.io/endpoint.svg?url=https%3A%2F%2Factions-badge.atrox.dev%2Fissue9%2Fsliceutil%2Fbadge%3Fref%3Dmaster&style=flat)](https://actions-badge.atrox.dev/issue9/sliceutil/goto?ref=master)
[![license](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat)](https://opensource.org/licenses/MIT)
[![codecov](https://codecov.io/gh/issue9/sliceutil/branch/master/graph/badge.svg)](https://codecov.io/gh/issue9/sliceutil)
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

// 返回第二个 7 所在的索引位置
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
```

安装
----

```shell
go get github.com/issue9/sliceutil
```

文档
----

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/issue9/sliceutil)

版权
----

本项目采用 [MIT](http://opensource.org/licenses/MIT) 开源授权许可证，完整的授权说明可在 [LICENSE](LICENSE) 文件中找到。
