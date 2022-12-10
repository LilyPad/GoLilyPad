package packet

import (
	"encoding/binary"
	"math/bits"
)

type BitSet struct {
	words      []uint64
	wordsInUse int
}

const addressBitsPerWord = 6
const bitsPerWord = 1 << addressBitsPerWord

func NewBitSet(initialSize int) *BitSet {
	return &BitSet{words: make([]uint64, wordIndex(initialSize-1)+1)}
}

func NewBitSetFrom(bytes []byte) *BitSet {
	remaining := len(bytes)
	for ; remaining > 0 && bytes[remaining-1] == 0; remaining-- {
	}
	words := make([]uint64, (remaining+7)/8)
	i, position := 0, 0
	for remaining >= 8 {
		words[i] = binary.LittleEndian.Uint64(bytes[position : position+8])
		i++
		position += 8
		remaining -= 8
	}
	for j := 0; j < remaining; j++ {
		words[i] |= uint64(bytes[position]&0xff) << uint64((8*j)%64)
		position++
	}
	return &BitSet{
		words:      words,
		wordsInUse: len(words),
	}
}

func (this *BitSet) Get(bitIndex int) bool {
	wordIndex := wordIndex(bitIndex)
	return wordIndex < this.wordsInUse && (this.words[wordIndex]&(uint64(1)<<uint64(bitIndex%64))) != 0
}

func (this *BitSet) Set(bitIndex int) {
	wordIndex := wordIndex(bitIndex)
	this.expandTo(wordIndex)
	this.words[wordIndex] |= uint64(1) << uint64(bitIndex%64)
}

func (this *BitSet) ToByteArray() []byte {
	if this.wordsInUse == 0 {
		return make([]byte, 0)
	}
	length := 8 * (this.wordsInUse - 1)
	for x := this.words[this.wordsInUse-1]; x != 0; x >>= 8 {
		length++
	}
	bytes := make([]byte, length)
	pos := 0
	for i := 0; i < this.wordsInUse-1; i++ {
		binary.LittleEndian.PutUint64(bytes[pos:pos+8], this.words[i])
		pos += 8
	}
	for x := this.words[this.wordsInUse-1]; x != 0; x >>= 8 {
		bytes[pos] = (byte)(x & 0xff)
		pos++
	}
	return bytes
}

func (this *BitSet) Len() int {
	if this.wordsInUse == 0 {
		return 0
	}
	return bitsPerWord*(this.wordsInUse-1) +
		(bitsPerWord - bits.LeadingZeros64(this.words[this.wordsInUse-1]))
}

func (this *BitSet) ensureCapacity(wordsRequired int) {
	if len(this.words) < wordsRequired {
		newWords := make([]uint64, max(2*len(this.words), wordsRequired))
		copy(newWords, this.words)
		this.words = newWords
	}
}

func (this *BitSet) expandTo(wordIndex int) {
	wordsRequired := wordIndex + 1
	if this.wordsInUse < wordsRequired {
		this.ensureCapacity(wordsRequired)
		this.wordsInUse = wordsRequired
	}
}

func wordIndex(bitIndex int) int {
	return bitIndex >> addressBitsPerWord
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
