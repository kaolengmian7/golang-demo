# GO的指针

## 普通指针

用于传递对象地址，不能进行指针运算

## unsafe.Pointer

通用指针类型，用于转换不同类型的指针，不能进行指针运算，不能读取内存中存储的值（必须转换到某一类型的普通指针）

## uintptr

用于指针运算。GC不把uintptr当指针，uintptr类型的目标会被回收。uintptr无法持有对象。

# type ArbitraryType int

ArbitraryType 代表GO的任意表达式，它并不属于unsafe包的一部分

> arbirary adj 任意的

# type Pointer *ArbitraryType

Pointer 是一个 ArbitraryType类型的指针。也就是说Pointer代表任意类型的指针。

Pointer类型支持其他类型所不支持的操作：

- 任何指针类型可以转换成Pointer类型
- Pointer类型可以转换成任何指针类型
- uintptr类型可以转换成Pointer类型
- Pointer类型可以转换成uintptr类型

==Pointer允许程序绕过类型系统去读写内存，故操作Pointer时要极其小心==

下面介绍几种使用Pointer的模版

> 如果不使用如下几种模版来编写Pointer代码，那么很用可能代码是无效的，或者说可能在未来变的有效

> 使用 go vet 命令可以查找出哪些关于Pointer的代码没有遵循这些模式。但是不能**以“go vet"没有查出来**作为符合这些模式的标准

## 模式一：指针类型的转换（*T1 to Pointer to *T2）

要求：

- T2 不大于 T1
- T2 和 T1 享有相同的内存布局

例子：

```go
math.Float64bits:

func Float64bits(f float64) uint64 {   
	return *(*uint64)(unsafe.Pointer(&f))
}
```

解释：

- float64 和 uint64 都是64位，符合要求

## 模式二：Pointer类型 to uintptr类型

- Pointer类型转换成uintptr类型会生成一个该值指向的内存地址
- 通常情况下，不允许将uintptr类型转换回Pointer类型
- unitptr类型是一个整形，并不是一个指针
- 即使uintptr保存着某个对象的地址，GC不会更新uintptr的值。如果对象被移动，uintptr也不会阻止GC回收该对象。