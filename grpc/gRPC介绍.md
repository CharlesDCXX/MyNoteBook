# gRPC

在gRPC里，客户端和服务端接口定义语言（interface Definition Language）和底层消息交换格式使用的是protocol buffers，所以可以用java写服务端，go、python等语言写客户端。客户端可以直接去调用不同服务器上不同语言的服务应用。这样可以更简单的创建分布式应用。

![](resource/rpc.svg)

# gRPC 核心概念、架构及生命周期

## 服务定义

gRPC给予定义服务的思想，指定可以通过其参数和返回类型远程调用的方法。
## 服务特点
低延迟、高度可扩展、分布式

## gprc分为四种服务方法

```proto
rpc SayHello(HelloRequest) returns (HelloResponse);
rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse);
rpc LotsOfGreetings(stream HelloRequest) returns (HelloResponse);
rpc BidiHello(stream HelloRequest) returns (stream HelloResponse);
```

## 创建proto文件

生成pb.go文件和name_grpc.pb.go文件
如果rpc方法中使用了流式服务，客户端服务和服务端服务方法是不一样的。

```teminal
$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    routeguide/route_guide.proto
```

* `name.pb.go` 其中包含所有 protocol buffer code，用于填充、序列化和检索请求和响应消息(message)类型。
* `name_grpc.pb.gp` 其中包含以下内容：

    1. 客户端使用`service`类型服务中定义的方法调用的接口类型（或存根stud）。

    2. 服务器要实现的接口类型，也具有`service`服务中定义的方法。

## 创建服务

1. 实现服务接口
2. 运行服务并监听请求

## 创建client

1. 创建链接
2. 调用方法