package stringutil

import "testing"

func TestEncrypt1(t *testing.T) {
	cases := []struct {
		v uint64
		k [4]uint32
		e uint64
	}{
		{0x0000000000000000, [4]uint32{0, 0, 0, 0}, 0x41ea3a0a94baa940},
		{0x0102030405060708, [4]uint32{0, 0, 0, 0}, 0x6a2f9cf3fccf3c55},
		{0x0000000000000000, [4]uint32{0x01234567, 0x12345678, 0x23456789, 0x3456789a}, 0x34e943b0900f5dcb},
		{0x0102030405060708, [4]uint32{0x01234567, 0x12345678, 0x23456789, 0x3456789a}, 0x773dc179878a81c0},
	}
	for _, c := range cases {
		got := Encrypt(c.v, c.k)
		if got != c.e {
			t.Errorf("Ecrypt(%q, %q) == %q, want %q", c.v, c.k, got, c.e)
		}
	}
}
