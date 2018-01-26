package benchtest

import (
	"fmt"
	"testing"
)

func BenchmarkStringMapInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sMapInsert(fmt.Sprintf("%d", i))
	}
}

func BenchmarkStringMapGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sMapGet(fmt.Sprintf("%d", i))
	}
}

func BenchmarkMetroMapInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		metroMapInsert(fmt.Sprintf("%d", i))
	}
}

func BenchmarkMetroMapGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		metroMapGet(fmt.Sprintf("%d", i))
	}
}

func BenchmarkIntMapInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iMapInsert(i)
	}
}

func BenchmarkIntMapGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iMapGet(i)
	}
}

func BenchmarkInt32MapInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i32MapInsert(int32(i))
	}
}

func BenchmarkInt32MapGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i32MapGet(int32(i))
	}
}

func BenchmarkInt64MapInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i64MapInsert(int64(i))
	}
}

func BenchmarkInt64MapGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		i64MapGet(int64(i))
	}
}

func BenchmarkFloat32MapInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f32MapInsert(float32(i))
	}
}

func BenchmarkFloat32MapGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f32MapGet(float32(i))
	}
}

func BenchmarkFloat64MapInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f64MapInsert(float64(i))
	}
}

func BenchmarkFloat64MapGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f64MapGet(float64(i))
	}
}

func BenchmarkUint64MapInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ui64MapInsert(uint64(i))
	}
}

func BenchmarkUint64MapGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ui64MapGet(uint64(i))
	}
}

func BenchmarkUint32MapInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ui32MapInsert(uint32(i))
	}
}

func BenchmarkUint32MapGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ui32MapGet(uint32(i))
	}
}
