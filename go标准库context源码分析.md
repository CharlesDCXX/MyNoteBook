# go标准库context源码分析

## context接口方法

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool) // 返回 context 的过期时间
    Done() <-chan struct{} // 返回 context 中的 channel
    Err() error // 返回错误
    Value(key any) any // 返回 context 中的对应 key 的值
}
```
