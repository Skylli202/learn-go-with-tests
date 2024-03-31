package kdbx_test

import (
	"testing"

	kdbx "github.com/Skylli202/learn-go-with-tests/kdbx/kdbx"
)

// KDBX File Format Specification
// https://keepass.info/help/kb/kdbx.html

var KDBX_BINARY_SAMPLE []byte = []byte{
	// Signature 1 (uint32)
	0x03, 0xD9, 0xA2, 0x9A,
	// Signature 2 (uint32)
	0x67, 0xFB, 0x4B, 0xB5,
	// Version (uint32)
	0x01, 0x00, 0x04, 0x00,
	// Header 1
	// ID:2, Length: 16, value: 31C1F2E6BF714350BE5805216AFC5AFF
	0x02, 0x10, 0x00, 0x00, 0x00, 0x31, 0xC1, 0xF2, 0xE6, 0xBF, 0x71, 0x43, 0x50, 0xBE, 0x58, 0x05, 0x21, 0x6A, 0xFC, 0x5A, 0xFF,
}

func TestKDBX(t *testing.T) {
	t.Run("Extract Signature 1", func(t *testing.T) {
		db := kdbx.KDBX{
			Data: KDBX_BINARY_SAMPLE,
		}

		got := db.ReadSignture1()

		if got != 0x9AA2d903 {
			t.Errorf("got %02X, want 0x9AA2d903", got)
		}
	})

	t.Run("Extract Signature 2", func(t *testing.T) {
		db := kdbx.KDBX{
			Data: KDBX_BINARY_SAMPLE,
		}

		got := db.ReadSignture2()

		if got != 0xB54BFB67 {
			t.Errorf("got %02X, want 0xB54BFB67", got)
		}
	})

	t.Run("Extract format version", func(t *testing.T) {
		db := kdbx.KDBX{
			Data: KDBX_BINARY_SAMPLE,
		}

		major, minor := db.FormatVersion()

		var want_major uint32 = 0x0004
		var want_minor uint32 = 0x0001

		if major != want_major {
			t.Errorf("got %02X, want %02X", major, want_major)
		}

		if minor != want_minor {
			t.Errorf("got %02X, want %02X", minor, want_minor)
		}
	})
}
