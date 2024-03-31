package kdbx

import (
	"encoding/binary"
)

// KDBX File Format Specification
// https://keepass.info/help/kb/kdbx.html

type KDBX struct {
	Data []byte
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
