# [protocol buffers定义](https://protobuf.dev/getting-started/gotutorial/)

协议缓冲区是 Google 用于序列化结构化数据的语言中立、平台中立、可扩展的机制 - 想想 XML，但更小、更快、更简单。 您只需定义一次数据的结构化方式，然后就可以使用特殊生成的源代码，使用各种语言轻松地在各种数据流中写入和读取结构化数据。

# Protocol Buffer 解决了什么问题？

协议缓冲区为大小高达几兆字节的类型化结构化数据包提供序列化格式。 该格式适用于短暂的网络流量和长期数据存储。 协议缓冲区可以使用新信息进行扩展，而无需使现有数据无效或需要更新代码。

协议缓冲区是 Google 最常用的数据格式。 它们广泛用于服务器间通信以及磁盘上数据的归档存储。 Protocol buffer 消息和服务由工程师编写的 .proto 文件描述。 下面显示了一条消息示例：

```proto
message Person {
  optional string name = 1;
  optional int32 id = 2;
  optional string email = 3;
}
```

# 在go语言中的快速上手使用

使用协议缓冲区语言的 proto3 版本，介绍如何使用协议缓冲区。通过创建一个简单的示例应用程序，展示如何
- 在 .proto 文件中定义消息格式。
- 使用协议缓冲区编译器。
- 使用Go协议缓冲区API来写入和读取消息。

示例是一个非常简单的“地址簿”应用程序，它可以在文件中读取和写入人们的联系方式。地址簿中的每个人都有姓名、ID、电子邮件地址和联系电话号码。

# 定义协议格式

## 首先创建一个addressbook.proto文件

```proto

syntax = "proto3";
package tutorial;

import "google/protobuf/timestamp.proto";

```

### go_package 选项定义包的导入路径，该路径将包含该文件的所有生成代码。 Go 包名称将是导入路径的最后一个路径组成部分。 例如，我们的示例将使用包名称“tutorialpb”。


```proto
option go_package = "github.com/protocolbuffers/protobuf/examples/go/tutorialpb";
```

### 接下来，是消息定义 
消息只是包含一组类型化字段的聚合。 许多标准简单数据类型都可用作字段类型，包括 bool、int32、float、double 和 string。 您还可以通过使用其他消息类型作为字段类型来向消息添加进一步的结构。

```proto
message Person {
  string name = 1;
  int32 id = 2;  // Unique ID number for this person.
  string email = 3;

  message PhoneNumber {
    string number = 1;
    PhoneType type = 2;
  }

  repeated PhoneNumber phones = 4;

  google.protobuf.Timestamp last_updated = 5;
}

enum PhoneType {
  PHONE_TYPE_UNSPECIFIED = 0;
  PHONE_TYPE_MOBILE = 1;
  PHONE_TYPE_HOME = 2;
  PHONE_TYPE_WORK = 3;
}

// Our address book file is just one of these.
message AddressBook {
  repeated Person people = 1;
}
```

在上面的示例中，Person 消息包含 PhoneNumber 消息，而 AddressBook 消息包含 Person 消息。 您甚至可以定义嵌套在其他消息中的消息类型 - 正如您所看到的，PhoneNumber 类型是在 Person 内部定义的。 如果您希望字段之一具有预定义值列表之一，您还可以定义枚举类型 - 此处您希望指定电话号码可以是 PHONE_TYPE_MOBILE、PHONE_TYPE_HOME 或 PHONE_TYPE_WORK 之一。

每个元素上的“= 1”、“= 2”标记标识该字段在二进制编码中使用的唯一“标记”。 标签编号 1-15 需要比编号更高的标签少一个字节进行编码，因此作为一种优化，您可以决定将这些标签用于常用或重复的元素，而将标签 16 和更高的标签留给不太常用的可选元素。 重复字段中的每个元素都需要重新编码标签编号，因此重复字段特别适合此优化。

如果未设置字段值，则使用默认值：数字类型为零，字符串为空字符串，布尔值为 false。 对于嵌入消息，默认值始终是消息的“默认实例”或“原型”，其未设置任何字段。 调用访问器来获取尚未显式设置的字段的值始终返回该字段的默认值。

如果字段重复，则该字段可以重复任意次数（包括零次）。 重复值的顺序将保留在协议缓冲区中。 将重复字段视为动态大小的数组。

```bash
protoc -I=./protocol --go_out=./protocol ./protocol/addressbook.proto
```

### The Protocol Buffer API
Generating addressbook.pb.go gives you the following useful types:

An AddressBook structure with a People field.
A Person structure with fields for Name, Id, Email and Phones.
A Person_PhoneNumber structure, with fields for Number and Type.
The type Person_PhoneType and a value defined for each value in the Person.PhoneType enum.
You can read more about the details of exactly what’s generated in the Go Generated Code guide, but for the most part you can treat these as perfectly ordinary Go types.

Here’s an example from the list_people command’s unit tests of how you might create an instance of Person:
```go
p := pb.Person{
    Id:    1234,
    Name:  "John Doe",
    Email: "jdoe@example.com",
    Phones: []*pb.Person_PhoneNumber{
        {Number: "555-4321", Type: pb.Person_PHONE_TYPE_HOME},
    },
}
```

### Writing a Message
The whole purpose of using protocol buffers is to serialize your data so that it can be parsed elsewhere. In Go, you use the proto library’s Marshal function to serialize your protocol buffer data. A pointer to a protocol buffer message’s struct implements the proto.Message interface. Calling proto.Marshal returns the protocol buffer, encoded in its wire format. For example, we use this function in the add_person command:
```go
book := &pb.AddressBook{}
// ...

// Write the new address book back to disk.
out, err := proto.Marshal(book)
if err != nil {
    log.Fatalln("Failed to encode address book:", err)
}
if err := ioutil.WriteFile(fname, out, 0644); err != nil {
    log.Fatalln("Failed to write address book:", err)
}
```
### Reading a Message
To parse an encoded message, you use the proto library’s Unmarshal function. Calling this parses the data in in as a protocol buffer and places the result in book. So to parse the file in the list_people command, we use:

// Read the existing address book.
in, err := ioutil.ReadFile(fname)
if err != nil {
    log.Fatalln("Error reading file:", err)
}
book := &pb.AddressBook{}
if err := proto.Unmarshal(in, book); err != nil {
    log.Fatalln("Failed to parse address book:", err)
}