package kdbx

import (
	"encoding/binary"
	"fmt"
)

// KDBX File Format Specification
// https://keepass.info/help/kb/kdbx.html

type KDBX struct {
	Data []byte
}

type Header struct {
	Value  []byte
	ID     uint8
	Length uint32
}

func (kdbx *KDBX) ReadSignture1() uint32 {
	return binary.LittleEndian.Uint32(kdbx.Data[0:4])
}

func (kdbx *KDBX) ReadSignture2() uint32 {
	return binary.LittleEndian.Uint32(kdbx.Data[4:8])
}

func (kdbx *KDBX) FormatVersion() (uint32, uint32) {
	major := binary.LittleEndian.Uint16(kdbx.Data[10:12])
	minor := binary.LittleEndian.Uint16(kdbx.Data[8:10])

	return uint32(major), uint32(minor)
}

func (kdbx *KDBX) ReadHeaders() []Header {
	headers := make([]Header, 0)

	var i uint32 = 12
	id := uint8(kdbx.Data[i])
	i++

	for id != 0 {
		header := Header{}
		header.ID = id

		switch header.ID {
		case 0:
			fmt.Println("TODO")
		case 1, 5, 6, 8, 9, 10:
			fmt.Println("Deprecated header found. Ignoring it.")
		case 2, 3, 4, 7:
			header.Length = binary.LittleEndian.Uint32(kdbx.Data[i:(i + 4)])
			i += 4
			header.Value = kdbx.Data[i:(i + header.Length)]
			i += header.Length
		default:
			fmt.Println("TODO")
		}

		headers = append(headers, header)
		if i < uint32(len(kdbx.Data)) {
			id = uint8(kdbx.Data[i])
			i += 1
		} else {
			id = 0
		}
	}

	return headers
}
