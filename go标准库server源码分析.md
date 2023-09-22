# go 启动 server 服务接收 http 连接(UTF-8 with BOM)

## 启动 Servers

[下面是官网例子](https://pkg.go.dev/net/http#hdr-Servers)

```go
// 将url与处理函数建立映射关系
http.Handle("/foo", fooHandler) 

// 将url与匿名处理函数之间建立映射关系
http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}) 

// 在8080端口启动监听服务和监听到url调用对应处理函数服务
log.Fatal(http.ListenAndServe(":8080", nil))
```

http.Handle 和 http.HandleFunc 两种方式都是实现 URL 路径与处理函数的映射关系，但是使用方式不同。

第一种方式需要显式地定义一个实现了 http.Handler 接口的处理函数，而第二种方式则可以直接使用匿名函数来定义处理逻辑。

一般来说，如果处理逻辑比较简单，且不需要在多个地方复用，那么可以使用第二种方式；如果处理逻辑比较复杂，或者需要在多个地方复用，那么可以使用第一种方式。

ListenAndServe 使用给定的地址和处理程序启动 HTTP 服务器。处理程序通常为 nil，这意味着使用 DefaultServeMux.Handle 和 HandleFunc 将处理程序添加到 DefaultServeMux 。

## http.Handle 在 server 中实现过程

Handle 函数将 fooHandler 函数注册在了当前 */foo* 路径下面，即当访问 "localhost:8080/foo" 时，会调用 fooHandler 函数。注册使用的是默认的 serveMux 即 DefaultServeMux 。

```go
func Handle(pattern string, handler Handler) { 
    // 默认的serveMux对其进行注册
    DefaultServeMux.Handle(pattern, handler) 
}
```

DefaultServeMux 是 *ServeMux 类型的指针。

```go
/**
DefaultServeMux 是一个全局变量，因此可以在整个应用程序中使用。
它提供了一个默认的 HTTP 请求多路复用器，可以方便地注册和匹配 HTTP 请求路由，
而不需要在每个处理器中都显式地创建一个独立的多路复用器。
在使用 net/http 包提供的服务器时，如果没有指定自定义的多路复用器，
就会使用 DefaultServeMux 作为默认的多路复用器。
*/
var DefaultServeMux = &defaultServeMux

var defaultServeMux ServeMux

type ServeMux struct {
    mu    sync.RWMutex
    m     map[string]muxEntry
    es    []muxEntry // 从最长到最短排序的切片。
    hosts bool       // 任何模式是否包含主机名
}

type muxEntry struct {
    h       Handler
    pattern string
}

```

serveMux这个结构体的作用比较复杂，
它是一个 HTTP 请求多路复用器。

它匹配每个输入请求的 URL 并调用最相似 URL 的注册模式（patterns）的处理函数来处理这个请求。
模式命名固定的根路径，例如“/favicon.ico”，或根子树，例如“/images/”（注意尾部斜杠）。
较长的模式优先于较短的模式，因此，如果同时为“/images/”和“/images/thumbnails/”注册了处理程序，则将为以“/images/thumbnails/”开头的路径调用后一个处理程序，而前者则将被调用 将接收对“/images/”子树中任何其他路径的请求。

注意以斜杆“/”结尾的模式（pattern）命名为根子树（subtree），因为模式 "/" 匹配的是所有未被其他模式匹配到的路径，而不仅仅是 Path 字段为 "/" 的路径。即如果要匹配所有的路径，可以使用模式 "/"，它会匹配所有未被其他模式匹配到的路径。这个规则在编写 HTTP 服务器时非常有用，可以方便地处理所有未被其他模式匹配到的请求。

如果子树已注册，并且收到命名子树根但不带尾部斜杠的请求，则 ServeMux 会将该请求重定向到子树根（添加尾部斜杠）。
可以通过单独注册不带尾部斜杠的路径来覆盖此行为。 例如，注册“/images/”会导致 ServeMux 将对“/images”的请求重定向到“/images/”，除非“/images”已单独注册。
模式可以选择以主机名开头，仅匹配该主机上的 URL。 特定于主机的模式优先于一般模式，因此处理程序可以注册两个模式“/codesearch”和“codesearch.google.com/”，而无需接管对“http://www.google.com/”的请求 ”。

即以主机名开头的模式，例如 "example.com/path"。这种模式只会匹配指定主机名的请求，而不会匹配其他主机名的请求。如果有多个模式都匹配了同一个请求，那么以主机名开头的模式会优先匹配，即主机名匹配优先于通用匹配。

意思是，如果要限制某一个模式只匹配特定的主机名，可以在模式开头加上主机名。这样可以方便地处理不同主机名下的请求，避免请求被错误地路由到其他服务器上。同时，还要注意主机名匹配的优先级，以避免出现匹配错误的情况。

>假设我们有一个 HTTP 服务器，它可以处理两个不同的主机名：example.com 和 example.net。我们想要为这两个主机分别注册不同的处理函数来处理请求。
>如果我们想要为 example.com 注册一个处理函数来处理路径为 /path 的请求，可以使用模式 "example.com/path"。而对于 example.net 的请求，则不会被这个模式匹配到。
>如果我们还想要为 example.net 注册一个处理函数，来处理路径为 /path 的请求，可以使用模式 "example.net/path"。这个模式只会匹配 example.net 的请求，不会匹配 example.com 的请求。
>如果我们还想要为所有主机名注册一个处理函数，来处理路径为 /path 的请求，可以使用模式 "/path"。这个模式会匹配所有主机名的请求，但如果有和主机名匹配的模式存在，那么主机名匹配的模式会优先匹配。例如，如果有一个模式 "example.com/"，则请求 "<http://example.com/path>" 会被这个模式匹配到，而不是通用模式 "/path"。

让我们看看 DefaultServeMux 的 Handle 方法具体实现

```go
// Handle 方法注册给定url的处理函数。
// 如果url的处理程序已经存在，则抛出panic。
func (mux *ServeMux) Handle(pattern string, handler Handler) {
    mux.mu.Lock()
    defer mux.mu.Unlock()

    if pattern == "" {
        panic("http: invalid pattern")
    }
    if handler == nil {
        panic("http: nil handler")
    }
    if _, exist := mux.m[pattern]; exist {
        panic("http: multiple registrations for " + pattern)
    }

    if mux.m == nil {
        mux.m = make(map[string]muxEntry)
    }
    e := muxEntry{h: handler, pattern: pattern}
    mux.m[pattern] = e
    if pattern[len(pattern)-1] == '/' {
        mux.es = appendSorted(mux.es, e)
    }

    if pattern[0] != '/' {
        mux.hosts = true
    }
}
```

其中主要逻辑是将 handler 方法与 pattern 放进 mux 的 map 结构体，key 为 pattern 。

```go
func (mux *ServeMux) Handle(pattern string, handler Handler) {
    e := muxEntry{h: handler, pattern: pattern}
    mux.m[pattern] = e
}
```

接下来就是  

```go
http.ListenAndServe(":8080", nil)
```

看一下 ListenAndServe 方法的源码

```go

func ListenAndServe(addr string, handler Handler) error {
    server := &Server{Addr: addr, Handler: handler}
    return server.ListenAndServe()
}

/**
ListenAndServe 监听TCP网络地址 srv.Addr 然后调用 Serve 和 handle 来处理连接上的请求
接受的连接配置为启用 TCP keep-alives。
ListenAndServe 始终返回非零错误。

*/
func (srv *Server) ListenAndServe() error {
    if srv.shuttingDown() {
        return ErrServerClosed
    }
    addr := srv.Addr
    if addr == "" {
        addr = ":http"
    }
    ln, err := net.Listen("tcp", addr)
    if err != nil {
        return err
    }
    return srv.Serve(ln)
}

/**
Serve 接受 Listener l 上的传入连接，为每个连接创建一个新的服务 goroutine。
服务 goroutine 读取请求，然后调用 srv.Handler 来回复它们。 仅当侦听器返回 *tls 时才会启用 HTTP/2 支持。
Conn 连接，它们在 TLS Config.NextProtos 中配置为“h2”。 Serve 总是返回一个非零错误并关闭 l。 Shutdown或Close后，返回的错误为ErrServerClosed。
*/
func (srv *Server) Serve(l net.Listener) error {
    if fn := testHookServerServe; fn != nil {
        fn(srv, l) // call hook with unwrapped listener
    }

    origListener := l
    l = &onceCloseListener{Listener: l}
    defer l.Close()

    if err := srv.setupHTTP2_Serve(); err != nil {
        return err
    }

    if !srv.trackListener(&l, true) {
        return ErrServerClosed
    }
    defer srv.trackListener(&l, false)

    baseCtx := context.Background()
    if srv.BaseContext != nil {
        baseCtx = srv.BaseContext(origListener)
        if baseCtx == nil {
            panic("BaseContext returned a nil context")
        }
    }

    var tempDelay time.Duration // how long to sleep on accept failure

    ctx := context.WithValue(baseCtx, ServerContextKey, srv)
    for {
        rw, err := l.Accept()
        if err != nil {
            if srv.shuttingDown() {
                return ErrServerClosed
            }
            if ne, ok := err.(net.Error); ok && ne.Temporary() {
                if tempDelay == 0 {
                    tempDelay = 5 * time.Millisecond
                } else {
                    tempDelay *= 2
                }
                if max := 1 * time.Second; tempDelay > max {
                    tempDelay = max
                }
                srv.logf("http: Accept error: %v; retrying in %v", err, tempDelay)
                time.Sleep(tempDelay)
                continue
            }
            return err
        }
        connCtx := ctx
        if cc := srv.ConnContext; cc != nil {
            connCtx = cc(connCtx, rw)
            if connCtx == nil {
                panic("ConnContext returned nil")
            }
        }
        tempDelay = 0
        c := srv.newConn(rw)
        c.setState(c.rwc, StateNew, runHooks) // before Serve can return
        go c.serve(connCtx)
    }
}
```
