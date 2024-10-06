# base-convert

## 说明 
baseconvert 适合用于 短URL生成 的场景，通过 Base62 和 Base54 编码缩短原始 URL，使生成的短链更简洁易读。这类短链常用于 URL 缩短服务、内容分发网络（CDN）优化等场景，并且广泛应用于社交媒体分享、广告跟踪等平台

## 例子
```go
package main

import (
	"fmt"
	"log"
	"github.com/nuominmin/base-convert"
)

func main() {
	// Base62 的字符表
	base62Alphabet := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	base62Codec, err := baseconvert.NewBaseNCodec(base62Alphabet)
	if err != nil {
		log.Fatal(err)
	}

	// 编码示例
	number := 12345
	encoded, err := base62Codec.Encode(number)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Encoded %d to base62: %s\n", number, encoded)

	// 解码示例
	decoded, err := base62Codec.Decode(encoded)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Decoded %s to base10: %d\n", encoded, decoded)

	// Base52 的字符表
	base52Alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	base52Codec, err := baseconvert.NewBaseNCodec(base52Alphabet)
	if err != nil {
		log.Fatal(err)
	}

	// 编码示例
	encodedBase52, err := base52Codec.Encode(number)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Encoded %d to base52: %s\n", number, encodedBase52)

	// 解码示例
	decodedBase52, err := base52Codec.Decode(encodedBase52)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Decoded %s to base10: %d\n", encodedBase52, decodedBase52)
}

```