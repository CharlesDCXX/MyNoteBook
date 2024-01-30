package main

import (
	pb "MyNoteBook/protocol"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
)

func main() {
	p := pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.PhoneType_PHONE_TYPE_HOME},
		},
	}
	book := &pb.AddressBook{}

	book.People = []*pb.Person{&p}
	// ...

	// Write the new address book back to disk.
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	fname := "./protocol/file.txt"
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	book1 := &pb.AddressBook{}
	if err := proto.Unmarshal(in, book1); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}
	fmt.Println(book1.String())
}
