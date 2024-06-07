# Go


## 创建workspace管理多mod
    mkdir workspace && cd workspace
    go work init 
    go work use ./hello
    go work use ./example/hello



## 本地mod重定向访问
    go mod edit -replace example.com/greetings=../greetings
    go mod tidy