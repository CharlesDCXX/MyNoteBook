package main

import (
	"context"
	"fmt"
	"unsafe"
)

type V struct {
	i int32
	j int64
}

func (this V) PutI() {
	fmt.Printf("i=%d\n", this.i)
}

func (this V) PutJ() {
	fmt.Printf("j=%d\n", this.j)
}

type SizeOfC struct {
	A byte  // 1字节
	C int32 // 4字节
}

func main() {
	fmt.Println(unsafe.Sizeof(int32(0)))
	var v *V = new(V)
	var i *int32 = (*int32)(unsafe.Pointer(v))
	*i = int32(98)
	var j *int64 = (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + unsafe.Sizeof(int32(0))))
	*j = int64(763)
	v.PutI()
	v.PutJ()
	fmt.Println(unsafe.Alignof(SizeOfC{0, 0}))
	context.Background()
}
