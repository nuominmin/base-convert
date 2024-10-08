package baseconvert

import "testing"

// TestBaseNCodec_EncodeDecode 测试 Encode 和 Decode 方法
func TestBaseNCodec_EncodeDecode(t *testing.T) {
	codec, _ := NewBaseNCodec("0123456789ABCDEF")

	testCases := []struct {
		number   uint64
		expected string
	}{
		{0, "0"},
		{1, "1"},
		{10, "A"},
		{15, "F"},
		{16, "10"},
		{255, "FF"},
	}

	for _, tc := range testCases {
		encoded, err := codec.Encode(tc.number)
		if err != nil {
			t.Fatalf("Failed to encode %d: %v", tc.number, err)
		}
		if encoded != tc.expected {
			t.Errorf("Expected encoded %d to be %s, but got %s", tc.number, tc.expected, encoded)
		}

		decoded, err := codec.Decode(encoded)
		if err != nil {
			t.Fatalf("Failed to decode %s: %v", encoded, err)
		}
		if decoded != tc.number {
			t.Errorf("Expected decoded %s to be %d, but got %d", encoded, tc.number, decoded)
		}
	}
}

// TestBaseNCodec_EncodeStringDecodeString 测试 EncodeString 和 DecodeString 方法
func TestBaseNCodec_EncodeStringDecodeString(t *testing.T) {
	codec, _ := NewBaseNCodec("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

	testCases := []struct {
		input    string
		expected string
	}{
		{"hello", "hello"},
		{"Base62", "Base62"},
		{"12345", "12345"},
		{"", ""},
	}

	for _, tc := range testCases {
		encoded, err := codec.EncodeString(tc.input)
		if err != nil {
			t.Fatalf("Failed to encode string %s: %v", tc.input, err)
		}

		decoded, err := codec.DecodeString(encoded)
		if err != nil {
			t.Fatalf("Failed to decode string %s: %v", encoded, err)
		}

		if decoded != tc.expected {
			t.Errorf("Expected decoded string to be %s, but got %s", tc.expected, decoded)
		}
	}
}

// TestBaseNCodec_InvalidCharacter 测试 Decode 方法中的无效字符处理
func TestBaseNCodec_InvalidCharacter(t *testing.T) {
	codec, _ := NewBaseNCodec("0123456789ABCDEF")

	_, err := codec.Decode("Z")
	if err == nil {
		t.Error("Expected error for invalid character 'Z', but got none")
	}
}

// TestBaseNCodec_EncodeString_LengthMismatch 测试 DecodeString 中的编码长度错误
func TestBaseNCodec_EncodeString_LengthMismatch(t *testing.T) {
	codec, _ := NewBaseNCodec("0123456789ABCDEF")

	_, err := codec.DecodeString("0A")
	if err == nil {
		t.Error("Expected error for mismatched encoded string length, but got none")
	}

	_, err = codec.DecodeString("123")
	if err == nil {
		t.Error("Expected error for mismatched encoded string length, but got none")
	}
}
