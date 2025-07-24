# gRPC

![](../resource/rpc.svg)


## gprc分为四种服务方法

```proto
// Unary RPCs 一元RPCs
rpc SayHello(HelloRequest) returns (HelloResponse);

// Server streaming RPCs 服务流RPCs
rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse);

//  Client streaming RPCs 客户端流RPCs
rpc LotsOfGreetings(stream HelloRequest) returns (HelloResponse);

// Bidirectional streaming RPCs 双向流RPCs
rpc BidiHello(stream HelloRequest) returns (stream HelloResponse);
```

## 创建proto文件

生成pb.go文件和name_grpc.pb.go文件

```teminal
$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    xxx.proto
```

* `name.pb.go` 其中包含所有 protocol buffer code，用于填充、序列化和检索请求和响应消息(message)类型。
* `name_grpc.pb.gp` 其中包含以下内容：

    1. 客户端使用`service`类型服务中定义的方法调用的接口类型。

    2. 服务器要实现的接口类型，也具有`service`服务中定义的方法。

## 创建服务

1. 实现服务接口
2. 运行服务并监听请求

## 创建client

1. 创建链接
2. 调用方法