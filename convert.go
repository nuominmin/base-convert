package baseconvert

import (
	"errors"
	"strings"
)

// BaseNCodec 结构体
type BaseNCodec struct {
	Alphabet string // 编码表，如 BASE62 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	Base     uint64 // 编码的进制数，如 BASE = 62
}

// NewBaseNCodec 创建一个 BaseNCodec 实例
func NewBaseNCodec(alphabet string) (*BaseNCodec, error) {
	base := uint64(len(alphabet))
	if base < 2 {
		return nil, errors.New("alphabet length must be at least 2")
	}
	return &BaseNCodec{
		Alphabet: alphabet,
		Base:     base,
	}, nil
}

func NewBase62Codec() *BaseNCodec {
	baseNCodec, _ := NewBaseNCodec("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	return baseNCodec
}

func NewBase52Codec() *BaseNCodec {
	baseNCodec, _ := NewBaseNCodec("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	return baseNCodec
}

func (bc *BaseNCodec) Encode(number uint64) (string, error) {
	if number == 0 {
		return string(bc.Alphabet[0]), nil
	}

	var result strings.Builder
	for number > 0 {
		result.WriteByte(bc.Alphabet[number%bc.Base])
		number /= bc.Base
	}

	encoded := result.String()

	// 翻转字符串，因为结果是从低位到高位的
	runes := []rune(encoded)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes), nil
}

// Decode 将一个 base 字符串解码为整数
func (bc *BaseNCodec) Decode(encoded string) (uint64, error) {
	mapAlphabetIdx := make(map[rune]int)
	for i, alphabet := range bc.Alphabet {
		mapAlphabetIdx[alphabet] = i
	}

	var result uint64
	for _, char := range encoded {
		index, exists := mapAlphabetIdx[char]
		if !exists {
			return 0, errors.New("invalid character in encoded string")
		}
		result = result*bc.Base + uint64(index)
	}

	return result, nil
}
