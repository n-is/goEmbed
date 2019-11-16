package golua

import (
	"math"
	"testing"
)

func TestNBodyDemo(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"MinuteSim", args{n: 50000000}, -0.169059907},
		// {"10SecSim", args{n: 10}, -0.169059907},
	}
	for _, tt := range tests {
		nb := SetupNBodyDemo()
		t.Run(tt.name, func(t *testing.T) {
			if got := RunNBodyDemo(nb, tt.args.n); math.Abs(got-tt.want) > 1e-8 {
				t.Errorf("NBodyDemo() = %v, want %v", got, tt.want)
			}
		})
	}
}

var resultNBody float64

func benchmarkNBody(i int64, b *testing.B) {
	var r float64
	nb := SetupNBodyDemo()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = RunNBodyDemo(nb, i)
	}
	resultNBody = r
}

func BenchmarkNBody10(b *testing.B)       { benchmarkNBody(10, b) }
func BenchmarkNBody50000000(b *testing.B) { benchmarkNBody(50000000, b) }

func benchmarkLuaNBody(i int64, b *testing.B) {
	var r float64
	nb := SetupLuaNBody()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = LuaNBody(nb, i)
	}
	resultNBody = r
}

func BenchmarkLuaNBodyM20(b *testing.B) { benchmarkLuaNBody(-20, b) }

func BenchmarkLuaNBody0(b *testing.B) { benchmarkLuaNBody(0, b) }

func BenchmarkLuaNBody10(b *testing.B) { benchmarkLuaNBody(10, b) }

func BenchmarkLuaNBody100(b *testing.B) { benchmarkLuaNBody(100, b) }

func BenchmarkLuaNBody1000(b *testing.B) { benchmarkLuaNBody(1000, b) }

func BenchmarkLuaNBody10000(b *testing.B) { benchmarkLuaNBody(10000, b) }

func BenchmarkLuaNBody100000(b *testing.B) { benchmarkLuaNBody(100000, b) }

func BenchmarkLuaNBody1000000(b *testing.B) { benchmarkLuaNBody(1000000, b) }

func BenchmarkLuaNBody10000000(b *testing.B) { benchmarkLuaNBody(10000000, b) }
