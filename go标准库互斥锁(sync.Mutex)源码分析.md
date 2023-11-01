# sync.Mutex源码分析(UTF-8 with BOM)

[sync.Mutex](https://pkg.go.dev/sync#Mutex)是Go标准库中常用的一个排外锁。当一个 goroutine 获得了这个锁的拥有权后， 其它请求锁的 goroutine 就会阻塞在 Lock 方法的调用上，直到锁被释放。

```go
type Counter struct {
    mu    sync.Mutex
    count int
}

func (c *Counter) Inc(wg *sync.WaitGroup) {
    wg.Add(1)
    c.mu.Lock()
    defer func ()  {
        c.mu.Unlock()
        wg.Done()
    } ()
    
    c.count++
}

func (c *Counter) Count() int {
    return c.count
}

func main() {
    var wg sync.WaitGroup
    counter := Counter{}
    for i := 0; i < 1000; i++ {
        go counter.Inc(&wg)
    }
    wg.Wait()
    fmt.Println(counter.Count())
}

```

在上面的代码中，我们定义了一个 Counter 结构体，该结构体包含一个 sync.Mutex 类型的字段 mu 和一个整数类型的字段 count。Inc 方法用来增加计数器的值，Count 方法用来获取计数器的值。在 Inc 和 Count 方法中，我们都使用了 mu.Lock() 和 mu.Unlock() 方法来加锁和解锁。

在 main 函数中，我们启动了 1000 个 goroutine 来增加计数器的值，最后输出计数器的值。由于使用了互斥锁，每个 goroutine 在访问计数器时都会先获得锁，这样就避免了竞争条件和数据竞争问题。

```go
/**
互斥锁是一种互斥锁。 互斥体的零值是未锁定的互斥体。

首次使用后不得复制互斥体。

在 Go 内存模型的术语中，对于任意 n < m，第 n 次 Unlock 调用“同步于”第 m 次 Lock 调用之前。 成功调用 TryLock 相当于调用 Lock。 对 TryLock 的失败调用根本不会建立任何“同步之前”关系。
*/
type Mutex struct {
    state int32
    sem  a  uint32
}

/**
Lock 锁 m。如果锁已被使用，则调用 Goroutine 会阻塞，直到互斥锁可用。
*/
func (m *Mutex) Lock() {
    // Fast path: grab unlocked mutex.
    if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
        if race.Enabled {
            race.Acquire(unsafe.Pointer(m))
        }
        return
    }
    // Slow path (outlined so that the fast path can be inlined)
    m.lockSlow()
}

func (m *Mutex) Unlock() {
    if race.Enabled {
        _ = m.state
        race.Release(unsafe.Pointer(m))
    }

    // Fast path: drop lock bit.
    new := atomic.AddInt32(&m.state, -mutexLocked)
    if new != 0 {
        // Outlined slow path to allow inlining the fast path.
        // To hide unlockSlow during tracing we skip one extra frame when tracing GoUnblock.
        m.unlockSlow(new)
    }
}
```

>互斥锁有两种状态：正常状态和饥饿状态。
>
>在正常状态下，所有等待锁的goroutine按照FIFO顺序等待。唤醒的goroutine不会直接拥有锁，而是会和新请求锁的goroutine竞争锁的拥有。新请求锁的goroutine具有优势：它正在CPU上执行，而且可能有好几个，所以刚刚唤醒的goroutine有很大可能在锁竞争中失败。在这种情况下，这个被唤醒的goroutine会加入到等待队列的前面。 如果一个等待的goroutine超过1ms没有获取锁，那么它将会把锁转变为饥饿模式。
>
>在饥饿模式下，锁的所有权将从unlock的gorutine直接交给交给等待队列中的第一个。新来的goroutine将不会尝试去获得锁，即使锁看起来是unlock状态, 也不会去尝试自旋操作，而是放在等待队列的尾部。
>
>如果一个等待的goroutine获取了锁，并且满足一以下其中的任何一个条件：(1)它是队列中的最后一个；(2)它等待的时候小于1ms。它会将锁的状态转换为正常状态。
>
>正常状态有很好的性能表现，饥饿模式也是非常重要的，因为它能阻止尾部延迟的现象。
