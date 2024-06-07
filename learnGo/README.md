# Go


## 创建workspace管理多mod
~~~GO
mkdir workspace && cd workspace
go work init 
go work use ./hello
go work use ./example/hello
~~~


## 本地mod重定向访问
~~~GO
go mod edit -replace example.com/greetings=../greetings
go mod tidy
~~~


## 语法
~~~GO
// 二个变量允许重复声明
f, err := os.Open(name)  
d, err := f.Stat()

// For循环的3中方式
for init; condition; post { }
for condition { }
for { }

// 遍历map array
for key, value := range oldMap {
    newMap[key] = value
}
for _, value := range array {
    sum += value
}

// Switch Condiction
switch {
case '0' <= c && c <= '9':
    return c - '0'
case 'a' <= c && c <= 'f':
    return c - 'a' + 10
case 'A' <= c && c <= 'F':
    return c - 'A' + 10
}
// Switch Variable
switch c {
case ' ', '?', '&', '=', '#', '+', '%':
    return true
default:
    return false
}

// Loop 跳出标识的循环

// Defer similary to finally,  LIFO order

// new 返回指针，初始化地址数据为空
// make for slices, maps, and channels 
var p *[]int = new([]int)       // allocates slice structure; *p == nil; rarely useful
var v  []int = make([]int, 100) // the slice v now refers to a new array of 100 ints

// Slices wrap arrays to give a more general, powerful, and convenient interface to sequences of data. Except for items with explicit dimension such as transformation matrices, most array programming in Go is done with slices rather than simple arrays.

