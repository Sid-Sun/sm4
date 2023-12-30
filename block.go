package sm4

import (
	"encoding/binary"
)

func cryptBlock(subkeys []uint32, dst, src []byte, decrypt bool) {
	left := binary.BigEndian.Uint32(src[0:4])
	r1 := binary.BigEndian.Uint32(src[4:8])
	r2 := binary.BigEndian.Uint32(src[8:12])
	r3 := binary.BigEndian.Uint32(src[12:16])

	if !decrypt {
		for i := 0; i < 32; i++ {
			// One Round of Feistel without key whitening baby
			t := r3
			r3 = left ^ feistelFunction(r1^r2^r3^subkeys[i])
			left = r1
			r1 = r2
			r2 = t
		}
	} else {
		for i := 31; i >= 0; i-- {
			// One Round of Feistel without key whitening baby
			t := r3
			r3 = left ^ feistelFunction(r1^r2^r3^subkeys[i])
			left = r1
			r1 = r2
			r2 = t
		}
	}

	// Write Output with reverse transformation
	binary.BigEndian.PutUint32(dst[0:4], r3)
	binary.BigEndian.PutUint32(dst[4:8], r2)
	binary.BigEndian.PutUint32(dst[8:12], r1)
	binary.BigEndian.PutUint32(dst[12:16], left)
}

func feistelFunction(input uint32) uint32 {
	// non-linear substitution
	substitutedWord := nonLinearSubstitution(input)
	// linear substitutiuon (diffusion)
	linearSubstitutedWord := linearSubstitute(substitutedWord)
	return linearSubstitutedWord
}

func generateSubKeys(key [16]byte) [32]uint32 {
	var interimKeys [4 + 32]uint32 // 4 whitened input keys and 32 intermediate ones

	// Put the initial key in interimKeys after XOR with a word of <TBD>
	for i := 0; i < 4; i++ {
		interimKeys[i] = binary.BigEndian.Uint32(key[i*4:(i*4)+4]) ^ initialSystemParams[i] // 0 -> 4 : 4 -> 8 : 8 -> 12 : 12 -> 16
	}

	numberOfRounds := 32
	for i := 0; i < numberOfRounds; i++ {
		roundSubKey := interimKeys[i] ^ tDash(interimKeys[i+1]^interimKeys[i+2]^interimKeys[i+3]^keyScheduleParams[i])
		interimKeys[i+4] = roundSubKey
	}

	subkeys := interimKeys[4:]
	return [32]uint32(subkeys)
}

func tDash(word uint32) uint32 {
	// non-linear substitution
	substitutedWord := nonLinearSubstitution(word)
	// linear substitutiuon (diffusion)
	linearSubstitutedWord := linearSubstituteDash(substitutedWord)
	return linearSubstitutedWord
}

func nonLinearSubstitution(word uint32) uint32 {
	var wordBytes [4]uint8

	wordBytes[0] = SboxTable[shift32ToGet4(word, 1)][shift32ToGet4(word, 2)]
	wordBytes[1] = SboxTable[shift32ToGet4(word, 3)][shift32ToGet4(word, 4)]
	wordBytes[2] = SboxTable[shift32ToGet4(word, 5)][shift32ToGet4(word, 6)]
	wordBytes[3] = SboxTable[shift32ToGet4(word, 7)][shift32ToGet4(word, 8)]

	result := concatenate16ToGet32(concatenate8ToGet16(wordBytes[0], wordBytes[1]), concatenate8ToGet16(wordBytes[2], wordBytes[3]))
	return result
}

func linearSubstitute(word uint32) uint32 { // i.e. diffusion.
	return word ^ rotate32byN(word, 2) ^ rotate32byN(word, 10) ^ rotate32byN(word, 18) ^ rotate32byN(word, 24)
}

func linearSubstituteDash(word uint32) uint32 { // i.e. diffusion.
	return word ^ rotate32byN(word, 13) ^ rotate32byN(word, 23) // 13 and 23 are magic values AFAIK
}
