package sm4

func rotate32byN(x uint32, N uint8) uint32 {
	return (x << N) | (x >> (32 - N))
}

func concatenate8ToGet16(a, b uint8) uint16 {
	return (uint16(a) << 8) | uint16(b)
}

func concatenate16ToGet32(a, b uint16) uint32 {
	return (uint32(a) << 16) | uint32(b)
}

func shift32ToGet4(x uint32, ind uint8) uint8 {
	// Indexing must start at 1
	return uint8((x << ((ind - 1) * 4)) >> 28)
}
