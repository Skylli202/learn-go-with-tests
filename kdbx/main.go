package main

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/Skylli202/learn-go-with-tests/kdbx/kdbx"
)

func main() {
	data, err := fs.ReadFile(os.DirFS("."), "sample.kdbx")
	if err != nil {
		panic(err)
	}

	kdbx := kdbx.KDBX{
		Data: data,
	}

	major, minor := kdbx.FormatVersion()

	fmt.Printf(`Signature 1: %X
Signature 2: %X
Format version: %d.%d
`, kdbx.ReadSignture1(), kdbx.ReadSignture2(), major, minor)

	headers := kdbx.ReadHeaders()
	for i, header := range headers {
		fmt.Printf(`[%d] ID: %d, Length: %d,
    Value: %X
`, i, header.ID, header.Length, header.Value)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
