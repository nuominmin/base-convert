package baseconvert

import (
	"errors"
	"math"
	"strings"
)

// BaseNCodec 结构体
type BaseNCodec struct {
	Alphabet string // 编码表，如 BASE62 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	Base     int    // 编码的进制数，如 BASE = 62
}

// NewBaseNCodec 创建一个 BaseNCodec 实例
func NewBaseNCodec(alphabet string) (*BaseNCodec, error) {
	base := len(alphabet)
	if base < 2 {
		return nil, errors.New("alphabet length must be at least 2")
	}
	return &BaseNCodec{
		Alphabet: alphabet,
		Base:     base,
	}, nil
}

// Encode 将一个整数编码为指定 base 的字符串
func (bc *BaseNCodec) Encode(number int) (string, error) {
	if number < 0 {
		return "", errors.New("number must be non-negative")
	}

	if number == 0 {
		return string(bc.Alphabet[0]), nil
	}

	var result []byte
	for number > 0 {
		result = append(result, bc.Alphabet[number%bc.Base])
		number /= bc.Base
	}

	// 翻转结果，因为我们从最低位开始追加字符
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result), nil
}

// Decode 将一个 base 字符串解码为整数
func (bc *BaseNCodec) Decode(encoded string) (int, error) {
	var result int
	for i, char := range encoded {
		index := strings.IndexRune(bc.Alphabet, char)
		if index == -1 {
			return 0, errors.New("invalid character in encoded string")
		}
		power := len(encoded) - i - 1
		result += index * int(math.Pow(float64(bc.Base), float64(power)))
	}
	return result, nil
}
