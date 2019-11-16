package gotengo

import (
	"math"
	"testing"
)

func TestTengoNBody(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"10SecSim", args{n: 10}, -0.16907302171469984},
		{"500SecSim", args{n: 500}, -0.16902152766534373},
	}
	for _, tt := range tests {
		nb := setupTengoNBody()
		t.Run(tt.name, func(t *testing.T) {
			if got := tengoNBody(nb, tt.args.n); math.Abs(got-tt.want) > 1e-8 {
				t.Errorf("tengoNBody() = %v, want %v", got, tt.want)
			}
		})
	}
}

var resultNBody float64

func benchmarkTengoNBody(i int64, b *testing.B) {
	var r float64
	nb := setupTengoNBody()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = tengoNBody(nb, i)
	}
	resultNBody = r
}

func BenchmarkTengoNBodyM20(b *testing.B) { benchmarkTengoNBody(-20, b) }

func BenchmarkTengoNBody0(b *testing.B) { benchmarkTengoNBody(0, b) }

func BenchmarkTengoNBody10(b *testing.B) { benchmarkTengoNBody(10, b) }

func BenchmarkTengoNBody100(b *testing.B) { benchmarkTengoNBody(100, b) }

func BenchmarkTengoNBody1000(b *testing.B) { benchmarkTengoNBody(1000, b) }

func BenchmarkTengoNBody10000(b *testing.B) { benchmarkTengoNBody(10000, b) }

func BenchmarkTengoNBody100000(b *testing.B) { benchmarkTengoNBody(100000, b) }

func BenchmarkTengoNBody1000000(b *testing.B) { benchmarkTengoNBody(1000000, b) }

func BenchmarkTengoNBody10000000(b *testing.B) { benchmarkTengoNBody(10000000, b) }
