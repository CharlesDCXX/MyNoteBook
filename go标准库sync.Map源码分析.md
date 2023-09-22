# sync.Map源码分析(UTF-8 with BOM)

原生的 Go Map 在并发读写场景下经常会遇到 panic 的情况。造成的原因是 map 是非线性安全的，并发读写过程中 map 的数据会被写乱。

而一般情况下，解决并发读写 map 的思路是加锁，或者把一个 map 切分成若干个小 map，对 key 进行哈希。
在业界中使用最多并发指出的模式分别是：

- 原生 map + 互斥锁 或者 读写锁
- 标准库 sync.Map (Go 1.9 及之后)

```go
/**
Map 类似于 Go 的 map[interface{}]interface{}，但可以安全地由多个 goroutine 并发使用，无需额外的锁定或协调。

loads、stores和deletes在分摊常量时间内运行。

Map 类型是特殊的 map。 
大多数代码应该使用普通的 Go 映射，并具有单独的锁定或协调，以获得更好的类型安全性，并更容易维护其他不变量以及映射内容。

Map 类型针对两种常见用例进行了优化：
(1) 当给定键的条目仅写入一次但读取多次时，如在只会增长的缓存中，
(2) 当多个 goroutine 读取、写入和覆盖不相交的键集的条目时。
在这两种情况下，与单独的 Mutex 或 RWMutex 配对的 Go Map 相比，使用 Map 可以显着减少锁争用。
*/
type Map struct {
    mu Mutex

    read atomic.Pointer[readOnly]

    dirty map[any]*entry

    misses int
}
```