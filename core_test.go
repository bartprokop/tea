package tea

import "testing"

func TestEncrypt1(t *testing.T) {
	cases := []struct {
		v0, v1         uint32
		k0, k1, k2, k3 uint32
		e0, e1         uint32
	}{
		{0x00000000, 0x00000000, 0, 0, 0, 0, 0x41ea3a0a, 0x94baa940},
		{0x01020304, 0x05060708, 0, 0, 0, 0, 0x6a2f9cf3, 0xfccf3c55},
		{0x00000000, 0x00000000, 0x01234567, 0x12345678, 0x23456789, 0x3456789a, 0x34e943b0, 0x900f5dcb},
		{0x01020304, 0x05060708, 0x01234567, 0x12345678, 0x23456789, 0x3456789a, 0x773dc179, 0x878a81c0},
	}
	for _, c := range cases {
		e0, e1 := Encrypt(c.v0, c.v1, c.k0, c.k1, c.k2, c.k3)
		if e0 != c.e0 || e1 != c.e1 {
			t.Errorf("Broken test vector")
		}
	}
}

// func TestEncrypt2(t *testing.T) {
// 	cases := []struct {
// 		v uint64
// 		k      [4]uint32
// 		e      uint64
// 	}{
// 		{0x0000000000000000, [4]uint32{0, 0, 0, 0}, 0x41ea3a0a94baa940},
// 		{0x0102030405060708, [4]uint32{0, 0, 0, 0}, 0x6a2f9cf3fccf3c55},
// 		{0x0000000000000000, [4]uint32{0x01234567, 0x12345678, 0x23456789, 0x3456789a}, 0x34e943b0900f5dcb},
// 		{0x0102030405060708, [4]uint32{0x01234567, 0x12345678, 0x23456789, 0x3456789a}, 0x773dc179878a81c0},
// 	}
// 	for _, c := range cases {
// 		got := Encrypt(c.v, c.k)
// 		if got != c.e {
// 			t.Errorf("Ecrypt(%q, %q) == %q, want %q", c.v, c.k, got, c.e)
// 		}
// 	}
// }
