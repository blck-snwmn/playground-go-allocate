package main

import "testing"

func Benchmark_setMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dest := make(map[int]string)
		setMap(dest)
		none(dest)
	}
}

func Benchmark_returnGenMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dest := returnGenMap()
		none(dest)
	}
}

func Benchmark_returnGeneratedMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dest := returnGeneratedMap()
		none(dest)
	}
}

// 確保した変数を消費するだけの関数.
func none(map[int]string) {}
