package fib

import (
	"testing"
	"time"
)

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(30)
	}
}

func BenchmarkFibReset(b *testing.B) {
	// 如果在循环前需要一些耗时的准备工作，我们就需要重置性能测试时间计数
	time.Sleep(3 * time.Second)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fib(30)
	}
}

func BenchmarkFibStart(b *testing.B) {
	// 如果在循环前需要一些耗时的准备工作，我们就需要重置性能测试时间计数
	b.StopTimer()
	time.Sleep(4 * time.Second)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		fib(30)
	}
}
