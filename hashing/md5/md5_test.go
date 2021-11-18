package md5

import (
	"encoding/hex"
	"testing"
)

func TestMd5(t *testing.T) {
	testCases := []struct {
		in       string
		expected string
	}{
		{"", "d41d8cd98f00b204e9800998ecf8427e"},
		{"a", "0cc175b9c0f1b6a831c399e269772661"},
		{"The quick brown fox jumps over the lazy dog", "9e107d9d372bb6826bd81d3542a419d6"},
		{"The quick brown fox jumps over the lazy dog.", "e4d909c290d0fb1ca068ffaddf22cbd0"},
	}
	for _, tc := range testCases {
		res := Md5([]byte(tc.in))
		result := hex.EncodeToString(res[:])
		if result != tc.expected {
			t.Fatalf("Md5(%s) = %s, expected %s", tc.in, result, tc.expected)
		}
	}
}
