package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello there")

	dat, err := os.ReadFile("./sample.kdbx")
	check(err)

	signature1 := binary.LittleEndian.Uint32(dat[:4])
	fmt.Printf("Signature 1: 0x%02X\n", signature1)

	signature2 := binary.LittleEndian.Uint32(dat[4:8])
	fmt.Printf("Signature 2: 0x%02X\n", signature2)

	version := binary.LittleEndian.Uint32(dat[8:12])
	fmt.Printf("Version: 0x%02X\n", version)

	fmt.Println("Headers:")
	fmt.Printf("ID: 0x%02X\n", dat[12])
	headerLength := binary.LittleEndian.Uint32(dat[13:17])
	fmt.Printf("Length: 0x%02X (%d)\n", headerLength, headerLength)
	fmt.Printf("Value: ")
	for _, byte := range dat[17:(17 + 16)] {
		fmt.Printf("%X", byte)
	}
	fmt.Println("")

	foo := []byte{0x00, 0x04}
	foo = append(foo, 0x12)
	fmt.Printf("%X %X %X\n", foo[0], foo[1], foo[2])
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
