package gotengo

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
		{"TengoFib_0", args{n: 0}, 0},
		{"TengoFib_1", args{n: 1}, 1},
		{"TengoFib_2", args{n: 2}, 1},
		{"TengoFib_3", args{n: 3}, 2},
		{"TengoFib_4", args{n: 4}, 3},
		{"TengoFib_5", args{n: 5}, 5},
		{"TengoFib_6", args{n: 6}, 8},
		{"TengoFib_20", args{n: 20}, 6765},
		{"TengoFib_40", args{n: 40}, 102334155},
		{"TengoFib_50", args{n: 50}, 12586269025},
		{"TengoFib_65", args{n: 65}, 17167680177565},
		{"TengoFib_75", args{n: 75}, 2111485077978050},
		{"TengoFib_78", args{n: 78}, 8944394323791464},
		{"TengoFib_79", args{n: 79}, 14472334024676221},
		{"TengoFib_-1", args{n: -1}, -1},   // Error Case
		{"TengoFib_-2", args{n: -2}, -1},   // Error Case
		{"TengoFib_-20", args{n: -20}, -1}, // Error Case
	}
	l := setupTengoFib()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tengoFib(l, tt.args.n); got != tt.want {
				t.Errorf("tengoFib() = %v, want %v", got, tt.want)
			}
		})
	}
}

var resultFib int64

func benchmarkTengoFib(i int, b *testing.B) {
	var r int64
	t := setupTengoFib()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = tengoFib(t, i)
	}
	resultFib = r
}

func BenchmarkTengoFibM20(b *testing.B)        { benchmarkTengoFib(-20, b) }
func BenchmarkTengoFibM1(b *testing.B)         { benchmarkTengoFib(-1, b) }
func BenchmarkTengoFib0(b *testing.B)          { benchmarkTengoFib(0, b) }
func BenchmarkTengoFib1(b *testing.B)          { benchmarkTengoFib(1, b) }
func BenchmarkTengoFib2(b *testing.B)          { benchmarkTengoFib(2, b) }
func BenchmarkTengoFib3(b *testing.B)          { benchmarkTengoFib(3, b) }
func BenchmarkTengoFib10(b *testing.B)         { benchmarkTengoFib(10, b) }
func BenchmarkTengoFib20(b *testing.B)         { benchmarkTengoFib(20, b) }
func BenchmarkTengoFib40(b *testing.B)         { benchmarkTengoFib(40, b) }
func BenchmarkTengoFib60(b *testing.B)         { benchmarkTengoFib(60, b) }
func BenchmarkTengoFib80(b *testing.B)         { benchmarkTengoFib(80, b) }
func BenchmarkTengoFib100(b *testing.B)        { benchmarkTengoFib(100, b) }
func BenchmarkTengoFib1000(b *testing.B)       { benchmarkTengoFib(1000, b) }
func BenchmarkTengoFib10000(b *testing.B)      { benchmarkTengoFib(10000, b) }
func BenchmarkTengoFib100000(b *testing.B)     { benchmarkTengoFib(100000, b) }
func BenchmarkTengoFib1000000(b *testing.B)    { benchmarkTengoFib(1000000, b) }
func BenchmarkTengoFib10000000(b *testing.B)   { benchmarkTengoFib(10000000, b) }
func BenchmarkTengoFib100000000(b *testing.B)  { benchmarkTengoFib(100000000, b) }
func BenchmarkTengoFib1000000000(b *testing.B) { benchmarkTengoFib(1000000000, b) }
