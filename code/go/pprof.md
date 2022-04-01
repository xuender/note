# pprof 性能测试

```go
import  _ "net/http/pprof"
```

## heap

```shell
wget http://localhost:8080/debug/pprof/heap
go tool pprof -http=:8081 heap
```

## 常用命令

查看堆栈调用信息
```shell
go tool pprof http://localhost:8080/debug/pprof/heap
```

查看 30 秒内的 CPU 信息
```shell
go tool pprof "http://localhost:8080/debug/pprof/profile?seconds=30"
```

查看 goroutine 阻塞
```shell
go tool pprof http://localhost:8080/debug/pprof/block
```

收集 5 秒内的执行路径
```shell
go tool pprof "http://localhost:8080/debug/pprof/trace?seconds=5"
```

争用互斥持有者的堆栈跟踪
```shell
go tool pprof http://localhost:8080/debug/pprof/mutex
```

## 内存对比

不同时间抓取2个heap

```shell
go tool pprof -http=:8081 -base heap1.pprof heap2.pprof
```
