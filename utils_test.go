package sm4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShift32ToGet4(t *testing.T) {
	word := uint32(0x12345678)
	assert.Equal(t, uint8(0x01), shift32ToGet4(word, 1))
	assert.Equal(t, uint8(0x02), shift32ToGet4(word, 2))
	assert.Equal(t, uint8(0x03), shift32ToGet4(word, 3))
	assert.Equal(t, uint8(0x04), shift32ToGet4(word, 4))
	assert.Equal(t, uint8(0x05), shift32ToGet4(word, 5))
	assert.Equal(t, uint8(0x06), shift32ToGet4(word, 6))
	assert.Equal(t, uint8(0x07), shift32ToGet4(word, 7))
	assert.Equal(t, uint8(0x08), shift32ToGet4(word, 8))
}

func TestRotate32byN(t *testing.T) {
	word := uint32(0x12345678)
	assert.Equal(t, uint32(0x12345678), rotate32byN(word, 0))
	assert.Equal(t, uint32(0x23456781), rotate32byN(word, 4))
	assert.Equal(t, uint32(0x34567812), rotate32byN(word, 8))
	assert.Equal(t, uint32(0x45678123), rotate32byN(word, 12))
	assert.Equal(t, uint32(0x56781234), rotate32byN(word, 16))
	assert.Equal(t, uint32(0x67812345), rotate32byN(word, 20))
	assert.Equal(t, uint32(0x78123456), rotate32byN(word, 24))
	assert.Equal(t, uint32(0x81234567), rotate32byN(word, 28))
	assert.Equal(t, uint32(0x12345678), rotate32byN(word, 32))
}

func TestConcatenate8ToGet16(t *testing.T) {
	word1 := uint8(0x12)
	word2 := uint8(0x34)
	assert.Equal(t, uint16(0x1234), concatenate8ToGet16(word1, word2))
}

func TestConcatenate16ToGet32(t *testing.T) {
	word1 := uint16(0x1234)
	word2 := uint16(0x5678)
	assert.Equal(t, uint32(0x12345678), concatenate16ToGet32(word1, word2))
}
