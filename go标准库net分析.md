# net 库(UTF-8 with BOM)

> net 库 提供了一个轻便的接口给 网络I/O,包括 TCP/IP, UDP, domain name resolution, and Unix domain sockets.

## 方法

### 建立与服务器的连接

```go
conn, err := net.Dial("tcp", "golang.org:80")
if err != nil {
    // handle error
}
fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
status, err := bufio.NewReader(conn).ReadString('\n')
```

### 使用 Listen 函数创建一个服务器

```go
ln, err := net.Listen("tcp", ":8080")
if err != nil {
    // handle error
}
for {
    conn, err := ln.Accept()
    if err != nil {
        // handle error
    }
    go handleConnection(conn)
}
```

Accept在未接受到连接时，会被堵塞，for循环会停在这里。

### 