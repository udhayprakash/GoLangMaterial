package benchmarking

import "testing"

// Usage: go test -bench=.

var bool_0 bool

func Benchmark_CopyBool(b *testing.B) {
	var bool_1 bool
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bool_0 = bool_1
	}
}

var byte_0 byte

func Benchmark_CopyByte(b *testing.B) {
	var byte_1 byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		byte_0 = byte_1
	}
}

var int16_0 byte

func Benchmark_CopyInt16(b *testing.B) {
	var int16_1 byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		int16_0 = int16_1
	}
}

var int32_0 byte

func Benchmark_CopyInt32(b *testing.B) {
	var int32_1 byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		int32_0 = int32_1
	}
}

var int64_0 byte

func Benchmark_CopyInt64(b *testing.B) {
	var int64_1 byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		int64_0 = int64_1
	}
}

var pointer_0 = new(int)

func Benchmark_CopyPointer(b *testing.B) {
	var pointer_1 = new(int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pointer_0 = pointer_1
	}
}

var string_0 string

func Benchmark_CopyString(b *testing.B) {
	var string_1 = "abc"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		string_0 = string_1
	}
}

var slice_0 []int

func Benchmark_CopySlice(b *testing.B) {
	var slice_1 []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		slice_0 = slice_1
	}
}

var array2_0 [2]int

func Benchmark_CopyArray_2_elements(b *testing.B) {
	var array2_1 [2]int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		array2_0 = array2_1
	}
}

var array3_0 [3]int

func Benchmark_CopyArray_3_elements(b *testing.B) {
	var array3_1 [3]int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		array3_0 = array3_1
	}
}

var array4_0 [4]int

func Benchmark_CopyArray_4_elements(b *testing.B) {
	var array4_1 [4]int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		array4_0 = array4_1
	}
}

var array5_0 [5]int

func Benchmark_CopyArray_5_elements(b *testing.B) {
	var array5_1 [5]int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		array5_0 = array5_1
	}
}

var array6_0 [6]int

func Benchmark_CopyArray_6_elements(b *testing.B) {
	var array6_1 [6]int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		array6_0 = array6_1
	}
}

var array7_0 [7]int

func Benchmark_CopyArray_7_elements(b *testing.B) {
	var array7_1 [7]int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		array7_0 = array7_1
	}
}

var array8_0 [8]int

func Benchmark_CopyArray_8_elements(b *testing.B) {
	var array8_1 [8]int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		array8_0 = array8_1
	}
}

var array9_0 [9]int

func Benchmark_CopyArray_9_elements(b *testing.B) {
	var array9_1 [9]int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		array9_0 = array9_1
	}
}

var array10_0 [10]int

func Benchmark_CopyArray_10_elements(b *testing.B) {
	var array10_1 [10]int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		array10_0 = array10_1
	}
}

var array11_0 [11]int

func Benchmark_CopyArray_11_elements(b *testing.B) {
	var array11_1 [11]int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		array11_0 = array11_1
	}
}

var array12_0 [12]int

func Benchmark_CopyArray_12_elements(b *testing.B) {
	var array12_1 [12]int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		array12_0 = array12_1
	}
}

var array13_0 [13]int

func Benchmark_CopyArray_13_elements(b *testing.B) {
	var array13_1 [13]int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		array13_0 = array13_1
	}
}

var struct2_0 struct{ a, b int }

func Benchmark_CopyStruct_2_fields(b *testing.B) {
	var struct2_1 struct{ a, b int }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		struct2_0 = struct2_1
	}
}

var struct3_0 struct{ a, b, c int }

func Benchmark_CopyStruct_3_fields(b *testing.B) {
	var struct3_1 struct{ a, b, c int }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		struct3_0 = struct3_1
	}
}

var struct4_0 struct{ a, b, c, d int }

func Benchmark_CopyStruct_4_fields(b *testing.B) {
	var struct4_1 struct{ a, b, c, d int }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		struct4_0 = struct4_1
	}
}

var struct5_0 struct{ a, b, c, d, e int }

func Benchmark_CopyStruct_5_fields(b *testing.B) {
	var struct5_1 struct{ a, b, c, d, e int }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		struct5_0 = struct5_1
	}
}

var struct6_0 struct{ a, b, c, d, e, f int }

func Benchmark_CopyStruct_6_fields(b *testing.B) {
	var struct6_1 struct{ a, b, c, d, e, f int }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		struct6_0 = struct6_1
	}
}

var struct7_0 struct{ a, b, c, d, e, f, g int }

func Benchmark_CopyStruct_7_fields(b *testing.B) {
	var struct7_1 struct{ a, b, c, d, e, f, g int }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		struct7_0 = struct7_1
	}
}

var struct8_0 struct{ a, b, c, d, e, f, g, h int }

func Benchmark_CopyStruct_8_fields(b *testing.B) {
	var struct8_1 struct{ a, b, c, d, e, f, g, h int }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		struct8_0 = struct8_1
	}
}

var struct9_0 struct{ a, b, c, d, e, f, g, h, i int }

func Benchmark_CopyStruct_9_fields(b *testing.B) {
	var struct9_1 struct{ a, b, c, d, e, f, g, h, i int }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		struct9_0 = struct9_1
	}
}

var struct10_0 struct{ a, b, c, d, e, f, g, h, i, j int }

func Benchmark_CopyStruct_10_fields(b *testing.B) {
	var struct10_1 struct{ a, b, c, d, e, f, g, h, i, j int }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		struct10_0 = struct10_1
	}
}

var struct11_0 struct{ a, b, c, d, e, f, g, h, i, j, k int }

func Benchmark_CopyStruct_11_fields(b *testing.B) {
	var struct11_1 struct{ a, b, c, d, e, f, g, h, i, j, k int }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		struct11_0 = struct11_1
	}
}

var struct12_0 struct{ a, b, c, d, e, f, g, h, i, j, k, l int }

func Benchmark_CopyStruct_12_fields(b *testing.B) {
	var struct12_1 struct{ a, b, c, d, e, f, g, h, i, j, k, l int }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		struct12_0 = struct12_1
	}
}

var struct13_0 struct{ a, b, c, d, e, f, g, h, i, j, k, l, m int }

func Benchmark_CopyStruct_13_fields(b *testing.B) {
	var struct13_1 struct{ a, b, c, d, e, f, g, h, i, j, k, l, m int }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		struct13_0 = struct13_1
	}
}

func main() {

}
