package baseconvert

import (
	"errors"
	"math"
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

func (bc *BaseNCodec) EncodeString(input string) (string, error) {
	var result strings.Builder
	maxLen := bc.maxEncodedLength()
	for _, char := range input {
		encodedChar, err := bc.Encode(uint64(char))
		if err != nil {
			return "", err
		}
		result.WriteString(leftPad(encodedChar, maxLen, rune(bc.Alphabet[0])))
	}
	return result.String(), nil
}

func (bc *BaseNCodec) DecodeString(encoded string) (string, error) {
	var result strings.Builder
	maxLen := bc.maxEncodedLength()
	for i := 0; i < len(encoded); i += maxLen {
		if i+maxLen > len(encoded) {
			return "", errors.New("invalid encoded string length")
		}
		encodedChar := encoded[i : i+maxLen]
		decodedChar, err := bc.Decode(encodedChar)
		if err != nil {
			return "", err
		}
		result.WriteRune(rune(decodedChar))
	}
	return result.String(), nil
}

func leftPad(str string, length int, pad rune) string {
	var result strings.Builder
	for i := 0; i < length-len(str); i++ {
		result.WriteRune(pad)
	}
	result.WriteString(str)
	return result.String()
}

func (bc *BaseNCodec) maxEncodedLength() int {
	// log(Base, 256) 计算得到每个字符编码后的最大长度。
	return int(math.Ceil(math.Log2(256) / math.Log2(float64(bc.Base))))
}
