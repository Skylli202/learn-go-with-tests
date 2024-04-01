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

const (
	VARIANT_DICTIONARY_TYPE_UINT32     = 0x04
	VARIANT_DICTIONARY_TYPE_UINT64     = 0x05
	VARIANT_DICTIONARY_TYPE_BOOLEAN    = 0x08
	VARIANT_DICTIONARY_TYPE_INT32      = 0x0C
	VARIANT_DICTIONARY_TYPE_INT64      = 0x0D
	VARIANT_DICTIONARY_TYPE_STRING     = 0x18
	VARIANT_DICTIONARY_TYPE_BYTE_ARRAY = 0x42
)

var VARIANT_DICTIONARY_TYPE_NAME = map[uint8]string{
	VARIANT_DICTIONARY_TYPE_UINT32:     "Uint32",
	VARIANT_DICTIONARY_TYPE_UINT64:     "Uint64",
	VARIANT_DICTIONARY_TYPE_BOOLEAN:    "Boolean",
	VARIANT_DICTIONARY_TYPE_INT32:      "Int32",
	VARIANT_DICTIONARY_TYPE_INT64:      "Int64",
	VARIANT_DICTIONARY_TYPE_STRING:     "String",
	VARIANT_DICTIONARY_TYPE_BYTE_ARRAY: "Byte[]",
}

type VariantDictionaryItem struct {
	Name      string
	Value     []byte
	NameSize  int32
	ValueSize int32
	Version   uint16
	ValueType byte
}

func DecodeVariantDictionaryItem(data []byte) VariantDictionaryItem {
	item := VariantDictionaryItem{}
	var i uint8 = 0

	item.Version = binary.BigEndian.Uint16(data[:2])
	i += 2

	item.ValueType = uint8(data[i])
	i += 1
	_, ok := VARIANT_DICTIONARY_TYPE_NAME[item.ValueType]
	if !ok {
		panic("TODO: Unsupported variant dictionary type.")
	}

	item.NameSize = int32(binary.LittleEndian.Uint32(data[i:(i + 4)]))
	i += 4

	item.Name = string(data[i:(i + uint8(item.NameSize))])
	i += uint8(item.NameSize)

	item.ValueSize = int32(binary.LittleEndian.Uint32(data[i:(i + 4)]))
	i += 4

	item.Value = data[i:(i + uint8(item.ValueSize))]
	i += uint8(item.ValueSize)

	return item
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
		case 2, 3, 4, 7, 11:
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
