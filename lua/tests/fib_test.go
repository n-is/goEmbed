package tests

import "testing"

func TestLuaFib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{"LuaFib_0", args{n: 0}, 0},
		{"LuaFib_1", args{n: 1}, 1},
		{"LuaFib_2", args{n: 2}, 1},
		{"LuaFib_3", args{n: 3}, 2},
		{"LuaFib_4", args{n: 4}, 3},
		{"LuaFib_5", args{n: 5}, 5},
		{"LuaFib_6", args{n: 6}, 8},
		{"LuaFib_20", args{n: 20}, 6765},
		{"LuaFib_40", args{n: 40}, 102334155},
		{"LuaFib_50", args{n: 50}, 12586269025},
		{"LuaFib_65", args{n: 65}, 17167680177565},
		{"LuaFib_75", args{n: 75}, 2111485077978050},
		{"LuaFib_78", args{n: 78}, 8944394323791464},
		// Output was 14472334024676220 (due to rounding error)
		{"LuaFib_79", args{n: 79}, 14472334024676221},
		{"LuaFib_-1", args{n: -1}, -1},   // Error Case
		{"LuaFib_-2", args{n: -2}, -1},   // Error Case
		{"LuaFib_-20", args{n: -20}, -1}, // Error Case
	}
	l := SetupLuaFib()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LuaFib(l, tt.args.n); got != tt.want {
				t.Errorf("LuaFib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{"Fib_0", args{n: 0}, 0},
		{"Fib_1", args{n: 1}, 1},
		{"Fib_2", args{n: 2}, 1},
		{"Fib_3", args{n: 3}, 2},
		{"Fib_4", args{n: 4}, 3},
		{"Fib_5", args{n: 5}, 5},
		{"Fib_6", args{n: 6}, 8},
		{"Fib_20", args{n: 20}, 6765},
		{"Fib_40", args{n: 40}, 102334155},
		{"Fib_50", args{n: 50}, 12586269025},
		{"Fib_65", args{n: 65}, 17167680177565},
		{"Fib_75", args{n: 75}, 2111485077978050},
		{"Fib_79", args{n: 79}, 14472334024676221},
		{"Fib_-1", args{n: -1}, -1},   // Error Case
		{"Fib_-2", args{n: -2}, -1},   // Error Case
		{"Fib_-20", args{n: -20}, -1}, // Error Case
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fib(tt.args.n); got != tt.want {
				t.Errorf("Fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

var resultFib int64

func benchmarkLuaFib(i int, b *testing.B) {
	var r int64
	l := SetupLuaFib()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = LuaFib(l, i)
	}
	resultFib = r
}

func BenchmarkLuaFibM20(b *testing.B)        { benchmarkLuaFib(-20, b) }
func BenchmarkLuaFibM1(b *testing.B)         { benchmarkLuaFib(-1, b) }
func BenchmarkLuaFib0(b *testing.B)          { benchmarkLuaFib(0, b) }
func BenchmarkLuaFib1(b *testing.B)          { benchmarkLuaFib(1, b) }
func BenchmarkLuaFib2(b *testing.B)          { benchmarkLuaFib(2, b) }
func BenchmarkLuaFib3(b *testing.B)          { benchmarkLuaFib(3, b) }
func BenchmarkLuaFib10(b *testing.B)         { benchmarkLuaFib(10, b) }
func BenchmarkLuaFib20(b *testing.B)         { benchmarkLuaFib(20, b) }
func BenchmarkLuaFib40(b *testing.B)         { benchmarkLuaFib(40, b) }
func BenchmarkLuaFib60(b *testing.B)         { benchmarkLuaFib(60, b) }
func BenchmarkLuaFib80(b *testing.B)         { benchmarkLuaFib(80, b) }
func BenchmarkLuaFib100(b *testing.B)        { benchmarkLuaFib(100, b) }
func BenchmarkLuaFib1000(b *testing.B)       { benchmarkLuaFib(1000, b) }
func BenchmarkLuaFib10000(b *testing.B)      { benchmarkLuaFib(10000, b) }
func BenchmarkLuaFib100000(b *testing.B)     { benchmarkLuaFib(100000, b) }
func BenchmarkLuaFib1000000(b *testing.B)    { benchmarkLuaFib(1000000, b) }
func BenchmarkLuaFib10000000(b *testing.B)   { benchmarkLuaFib(10000000, b) }
func BenchmarkLuaFib100000000(b *testing.B)  { benchmarkLuaFib(100000000, b) }
func BenchmarkLuaFib1000000000(b *testing.B) { benchmarkLuaFib(1000000000, b) }

func benchmarkFib(i int, b *testing.B) {
	var r int64
	for n := 0; n < b.N; n++ {
		r = Fib(i)
	}
	resultFib = r
}

func BenchmarkFibM20(b *testing.B)        { benchmarkFib(-20, b) }
func BenchmarkFibM1(b *testing.B)         { benchmarkFib(-1, b) }
func BenchmarkFib0(b *testing.B)          { benchmarkFib(0, b) }
func BenchmarkFib1(b *testing.B)          { benchmarkFib(1, b) }
func BenchmarkFib2(b *testing.B)          { benchmarkFib(2, b) }
func BenchmarkFib3(b *testing.B)          { benchmarkFib(3, b) }
func BenchmarkFib10(b *testing.B)         { benchmarkFib(10, b) }
func BenchmarkFib20(b *testing.B)         { benchmarkFib(20, b) }
func BenchmarkFib40(b *testing.B)         { benchmarkFib(40, b) }
func BenchmarkFib60(b *testing.B)         { benchmarkFib(60, b) }
func BenchmarkFib80(b *testing.B)         { benchmarkFib(80, b) }
func BenchmarkFib100(b *testing.B)        { benchmarkFib(100, b) }
func BenchmarkFib1000(b *testing.B)       { benchmarkFib(1000, b) }
func BenchmarkFib10000(b *testing.B)      { benchmarkFib(10000, b) }
func BenchmarkFib100000(b *testing.B)     { benchmarkFib(100000, b) }
func BenchmarkFib1000000(b *testing.B)    { benchmarkFib(1000000, b) }
func BenchmarkFib10000000(b *testing.B)   { benchmarkFib(10000000, b) }
func BenchmarkFib100000000(b *testing.B)  { benchmarkFib(100000000, b) }
func BenchmarkFib1000000000(b *testing.B) { benchmarkFib(1000000000, b) }
