package stringutil

func Encrypt(v uint64, k [4]uint32) uint64 {
	var v0 uint32 = uint32(v >> 32)
	var v1 uint32 = uint32(v & 0xffffffff)
	var sum uint32 = 0x0
	const delta uint32 = 0x9e3779b9
	var k0 uint32 = k[0]
	var k1 uint32 = k[1]
	var k2 uint32 = k[2]
	var k3 uint32 = k[3]
	for i := 0; i < 32; i++ {
		sum += delta
		v0 += ((v1 << 4) + k0) ^ (v1 + sum) ^ ((v1 >> 5) + k1)
		v1 += ((v0 << 4) + k2) ^ (v0 + sum) ^ ((v0 >> 5) + k3)
	}
	return (uint64(v0) << 32) | uint64(v1)
}

// void decrypt (uint32_t* v, uint32_t* k) {
//     uint32_t v0=v[0], v1=v[1], sum=0xC6EF3720, i;  /* set up */
//     uint32_t delta=0x9e3779b9;                     /* a key schedule constant */
//     uint32_t k0=k[0], k1=k[1], k2=k[2], k3=k[3];   /* cache key */
//     for (i=0; i<32; i++) {                         /* basic cycle start */
//         v1 -= ((v0<<4) + k2) ^ (v0 + sum) ^ ((v0>>5) + k3);
//         v0 -= ((v1<<4) + k0) ^ (v1 + sum) ^ ((v1>>5) + k1);
//         sum -= delta;
//     }                                              /* end cycle */
//     v[0]=v0; v[1]=v1;
// }
