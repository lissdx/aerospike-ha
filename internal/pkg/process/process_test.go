package process

import (
	"testing"
)

func BenchmarkProcessName_String(b *testing.B) {
	nameable := NewNameable("A_PROCESS_NAME")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = nameable.Name().String()
	}
}

func BenchmarkProcessNameCast_String(b *testing.B) {
	nameable := NewNameable("A_PROCESS_NAME")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = string(nameable.Name())
	}
}

func BenchmarkProcessNameWithFuncCast_String(b *testing.B) {
	nameable := newNameableWithCastForTest("A_PROCESS_NAME")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = nameable.Name().String()
	}
}

func BenchmarkProcessNameCastWithFuncCast_String(b *testing.B) {
	nameable := newNameableWithCastForTest("A_PROCESS_NAME")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = string(nameable.Name())
	}
}

func BenchmarkProcessNameWithFuncCast2_String(b *testing.B) {
	nameable := newNameableWithCastForTest2("A_PROCESS_NAME")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = nameable.Name().String()
	}
}

func BenchmarkProcessNameCastWithFuncCast2_String(b *testing.B) {
	nameable := newNameableWithCastForTest2("A_PROCESS_NAME")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = string(nameable.Name())
	}
}

func Benchmark_String(b *testing.B) {
	s := "A_PROCESS_NAME"
	//nameable := NewNameable("A_PROCESS_NAME")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s
	}
}
