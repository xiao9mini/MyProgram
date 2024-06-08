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

//  2D slice 声明，避免slice扩容产生的内存变化
// Allocate the top-level slice. Loop over the rows, allocating the slice for each row.
picture := make([][]uint8, YSize) // One row per unit of y.
for i := range picture {
    picture[i] = make([]uint8, XSize)
}

// 不定长参数
func Min(a ...int) int {
	min := int(^uint(0) >> 1) // largest int
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

// slices, maps hold references to an underlying data structure. 数据修改同步
// apeend slices
x := []int{1,2,3}
y := []int{4,5,6}
x = append(x, y...)

~~~
# 输出格式化
~~~GO
// fmt数据包： 标准输出、指定输出、字符串
Printf, Fprintf and Sprintf
Println, Fprintln, and Sprintln
Print, Fprint, and Sprint
// 通用参数
// %v, just fmt.Println(timeZone)
// %T, which prints the type of a value.
// %q, quoted string format

// 结构体重写string
func (t *T) String() string {
    return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
}
~~~
# 数值结构体
~~~GO
type ByteSize float64   // 单一数值结构体,无变量名称
const (
    _           = iota // ignore first value by assigning to blank identifier
    KB ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)
// 重写ByteSize的String
func (b ByteSize) String() string {
    switch {
    case b >= YB:
        return fmt.Sprintf("%.2fYB", b/YB)
    case b >= ZB:
        return fmt.Sprintf("%.2fZB", b/ZB)
    case b >= EB:
        return fmt.Sprintf("%.2fEB", b/EB)
    case b >= PB:
        return fmt.Sprintf("%.2fPB", b/PB)
    case b >= TB:
        return fmt.Sprintf("%.2fTB", b/TB)
    case b >= GB:
        return fmt.Sprintf("%.2fGB", b/GB)
    case b >= MB:
        return fmt.Sprintf("%.2fMB", b/MB)
    case b >= KB:
        return fmt.Sprintf("%.2fKB", b/KB)
    }
    return fmt.Sprintf("%.2fB", b)
}

~~~
# 环境变量
~~~GO
home   := os.Getenv("HOME")
numCPU := runtime.NumCPU()

// 文件初始化函数，常用：真正执行开始之前验证或修复程序状态的正确性。
func init() {

}

~~~
# 服务器接口 - 类型
~~~GO
// in http Any object that implements Handler can serve HTTP requests.
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
-------------------------------
type Counter struct {
    n int
}
func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    ctr.n++
    fmt.Fprintf(w, "counter = %d\n", ctr.n)
}
ctr := new(Counter)
http.Handle("/counter", ctr)

// Simpler counter server.
type Counter int

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    *ctr++
    fmt.Fprintf(w, "counter = %d\n", *ctr)
}

// with chan
// A channel that sends a notification on each visit.
// (Probably want the channel to be buffered.)
type Chan chan *http.Request
func (ch Chan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    ch <- req
    fmt.Fprint(w, "notification sent")
}
~~~
# 服务器接口 - 函数接口
~~~GO
// in http , ServeHTTP calls f(w, req).
type HandlerFunc func(ResponseWriter, *Request)
func (f HandlerFunc) ServeHTTP(w ResponseWriter, req *Request) {
    f(w, req)
}
-------------------------------
// Argument server.
func ArgServer(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(w, os.Args)
}
http.Handle("/args", http.HandlerFunc(ArgServer))



~~~
# 并发
~~~GO
不要通过共享内存进行通信；相反，通过通信来共享内存。

// Goroutines
go list.Sort()  // run list.Sort concurrently; don't wait for it.

// Channels
ci := make(chan int)            // unbuffered channel of integers
cj := make(chan int, 0)         // unbuffered channel of integers
cs := make(chan *os.File, 100)  // buffered channel of pointers to Files
~~~
## 并发 - 主线程内，通过channel等待Goroutines完成
~~~GO
c := make(chan int)  // Allocate a channel.
// Start the sort in a goroutine; when it completes, signal on the channel.
go func() {
    list.Sort()
    c <- 1  // Send a signal; value does not matter.
}()
doSomethingForAWhile()
<-c   // Wait for sort to finish; discard sent value.
~~~
## 并发 - 缓冲通道可以像信号量一样使用，例如限制吞吐量。
~~~GO
并发使用go goroutine时，注意变量复用的问题
// for loop, the loop variable is reused for each iteration, so the req variable is shared across all goroutines.
// We need to make sure that req is unique for each goroutine.
var sem = make(chan int, MaxOutstanding)
func Serve(queue chan *Request) {
    for req := range queue {
        req := req // Create new instance of req for the goroutine.
        sem <- 1
        go func() {
            process(req)
            <-sem
        }()
    }
}

~~~
## 并发 - 创建固定数量的goroutine处理请求
~~~GO
// 处理请求的函数，持续获取请求
func handle(queue chan *Request) {
    for r := range queue {
        process(r)
    }
}
// 创建多线程
func Serve(clientRequests chan *Request, quit chan bool) {
    for i := 0; i < MaxOutstanding; i++ {
        go handle(clientRequests)
    }
    <-quit  // Wait to be told to exit.
}

~~~
## 并发 - A rate-limited, parallel, non-blocking RPC system
~~~GO
// 参数、函数、通信管道
type Request struct {
    args        []int
    f           func([]int) int
    resultChan  chan int
}

// 处理请求的函数：持续获取请求，执行请求中的函数，返回结果到客户端
func handle(queue chan *Request) {
    for req := range queue {
        req.resultChan <- req.f(req.args)
    }
}

// 创建多线程
func Serve(clientRequests chan *Request, quit chan bool) {
    for i := 0; i < MaxOutstanding; i++ {
        go handle(clientRequests)
    }
    <-quit  // Wait to be told to exit.
}

// Build request
request := &Request{[]int{3, 4, 5}, sum, make(chan int)}
// Send request
clientRequests <- request
// Wait for response.
fmt.Printf("answer: %d\n", <-request.resultChan)
~~~
# 内存复用
~~~GO
var freeList = make(chan *Buffer, 100)
var serverChan = make(chan *Buffer)

func client() {
    for {
        var b *Buffer
        // Grab a buffer if available; allocate if not.
        select {
        case b = <-freeList:
            // Got one; nothing more to do.
        default:
            // None free, so allocate a new one.
            b = new(Buffer)
        }
        load(b)              // Read next message from the net.
        serverChan <- b      // Send to server.
    }
}

func server() {
    for {
        b := <-serverChan    // Wait for work.
        process(b)
        // Reuse buffer if there's room.
        select {
        case freeList <- b:
            // Buffer on free list; nothing more to do.
        default:
            // Free list full, just carry on.
        }
    }
}

