package slice

import (
	"testing"
)

func Benchmark_setSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dest := make([]byte, 3)
		setSlice(dest)
		none(dest)
	}
}

func Benchmark_returnGenSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dest := returnGenSlice()
		none(dest)
	}
}

func Benchmark_returnSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dest := returnSlice()
		none(dest)
	}
}

// 確保した変数を消費するだけの関数.
func none([]byte) {}
