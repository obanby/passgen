package randstr_test

import (
	"bytes"
	"strings"
	"testing"
)

func BenchmarkStringConcatWithBuilder(b *testing.B) {
	str := strings.Builder{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		str.Write([]byte{byte(i)})
	}
}

func BenchmarkStringConcatWithBuffer(b *testing.B) {
	str := &bytes.Buffer{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		str.Write([]byte{byte(i)})
	}
}

func BenchmarkStringConcat(b *testing.B) {
	str := ""
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		str += string(byte(i))
	}
}
