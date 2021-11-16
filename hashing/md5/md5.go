package md5

import "encoding/binary"

var s = [64]uint32{
	7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22, 7, 12, 17, 22,
	5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20, 5, 9, 14, 20,
	4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23, 4, 11, 16, 23,
	6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21, 6, 10, 15, 21,
}

var K = [64]uint32{
	0xd76aa478, 0xe8c7b756, 0x242070db, 0xc1bdceee,
	0xf57c0faf, 0x4787c62a, 0xa8304613, 0xfd469501,
	0x698098d8, 0x8b44f7af, 0xffff5bb1, 0x895cd7be,
	0x6b901122, 0xfd987193, 0xa679438e, 0x49b40821,
	0xf61e2562, 0xc040b340, 0x265e5a51, 0xe9b6c7aa,
	0xd62f105d, 0x02441453, 0xd8a1e681, 0xe7d3fbc8,
	0x21e1cde6, 0xc33707d6, 0xf4d50d87, 0x455a14ed,
	0xa9e3e905, 0xfcefa3f8, 0x676f02d9, 0x8d2a4c8a,
	0xfffa3942, 0x8771f681, 0x6d9d6122, 0xfde5380c,
	0xa4beea44, 0x4bdecfa9, 0xf6bb4b60, 0xbebfbc70,
	0x289b7ec6, 0xeaa127fa, 0xd4ef3085, 0x04881d05,
	0xd9d4d039, 0xe6db99e5, 0x1fa27cf8, 0xc4ac5665,
	0xf4292244, 0x432aff97, 0xab9423a7, 0xfc93a039,
	0x655b59c3, 0x8f0ccc92, 0xffeff47d, 0x85845dd1,
	0x6fa87e4f, 0xfe2ce6e0, 0xa3014314, 0x4e0811a1,
	0xf7537e82, 0xbd3af235, 0x2ad7d2bb, 0xeb86d391,
}

func pad(message *[]byte) {
	originalLength := make([]byte, 8)
	binary.BigEndian.PutUint64(originalLength, uint64(len(*message)))
	*message = append(*message, 0x80) // 0x80 = '1000 0000' append a 1 bit, then 0 bits
	for len(*message)%64 < 56 {       // append 0 bits until 64 bits (8 bytes) before 512 bits (64 bytes) multiple
		*message = append(*message, 0x00) // 0x00 = '0000 0000'
	}
	*message = append(*message, originalLength...)
}

func process(message []byte) (digest [16]byte) {
	// Process per 512 bits (64 bytes) chunk
	a0 := uint32(0x67452301)
	b0 := uint32(0xefcdab89)
	c0 := uint32(0x98badcfe)
	d0 := uint32(0x10325476)
	for i := uint32(0); i < uint32(len(message)); i += 64 {
		// Initialize hash value for this chunk
		A := a0
		B := b0
		C := c0
		D := d0

		for j := uint32(0); j < 64; j++ {
			var F uint32
			var g uint32
			switch {
			case j < 16:
				F = (B & C) | ((^B) & D)
				g = j
			case j >= 16 && j < 32:
				F = (D & B) | ((^D) & C)
				g = (5*j + 1) % 16
			case j >= 32 && j < 48:
				F = B ^ C ^ D
				g = (3*j + 5) % 16
			case j >= 48:
				F = C ^ (B | (^D))
				g = (7 * j) % 16
			}
			F = F + A + K[j] + binary.LittleEndian.Uint32(message[i+g*4:i+(g+1)*4]) // select a 32 bit (4 bytes) block
			A = D
			D = C
			C = B
			B = B + F<<s[j]
		}
		// Add this chunk's hash to result so far:
		a0 += A
		b0 += B
		c0 += C
		d0 += D
	}

	// Write result
	binary.LittleEndian.PutUint32(digest[:4], a0)
	binary.LittleEndian.PutUint32(digest[4:8], b0)
	binary.LittleEndian.PutUint32(digest[8:12], c0)
	binary.LittleEndian.PutUint32(digest[12:16], d0)

	return
}

func Md5(message []byte) (digest [16]byte) {
	pad(&message)
	return process(message)
}
