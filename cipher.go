package sm4

import (
	"crypto/cipher"
	"strconv"
)

type sm4Cipher struct {
	subkeys [32]uint32
}

const BlockSize = 16

type KeySizeError int

func (k KeySizeError) Error() string {
	return "SM4: invalid key size " + strconv.Itoa(int(k))
}

func NewCipher(key []byte) (cipher.Block, error) {

	switch len(key) {
	case 16:
		break
	default:
		return nil, KeySizeError(len(key))
	}

	c := new(sm4Cipher)
	c.subkeys = generateSubKeys([16]byte(key))

	return c, nil
}

func (s sm4Cipher) BlockSize() int {
	return BlockSize
}

func (s sm4Cipher) Encrypt(dst, src []byte) {
	if len(src) < BlockSize {
		panic("SM4: input not full block")
	}
	cryptBlock(s.subkeys[:], dst, src, false)
}

func (s sm4Cipher) Decrypt(dst, src []byte) {
	if len(src) < BlockSize {
		panic("SM4: input not full block")
	}
	cryptBlock(s.subkeys[:], dst, src, true)
}
