package golua

import (
	"errors"
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
		nb := setupNBodyDemo()
		t.Run(tt.name, func(t *testing.T) {
			if got := runNBodyDemo(nb, tt.args.n); math.Abs(got-tt.want) > 1e-8 {
				t.Errorf("NBodyDemo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_luaNBody(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
		{"10SecSim", args{n: 10}, float64(-0.16907302171469984)},
		{"500SecSim", args{n: 500}, float64(-0.16902152766534373)},
		{"M10SecSim", args{n: -10}, int64(0)},
	}
	for _, tt := range tests {
		nb := setupLuaNBody()
		t.Run(tt.name, func(t *testing.T) {
			v, err := getFloat64(tt.want)
			if err == nil {
				if got := luaNBody(nb, tt.args.n); math.Abs(got-v) > 1e-8 {
					t.Errorf("luaNBody() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func getFloat64(v interface{}) (float64, error) {
	switch v := v.(type) {
	case int64:
		return float64(v), nil
	case float64:
		return v, nil
	case float32:
		return float64(v), nil
	}
	return 0, errors.New("Unsupported Type")
}

var resultNBody float64

func benchmarkNBody(i int64, b *testing.B) {
	var r float64
	nb := setupNBodyDemo()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = runNBodyDemo(nb, i)
	}
	resultNBody = r
}

func BenchmarkNBody10(b *testing.B)       { benchmarkNBody(10, b) }
func BenchmarkNBody50000000(b *testing.B) { benchmarkNBody(50000000, b) }

func benchmarkLuaNBody(i int64, b *testing.B) {
	var r float64
	nb := setupLuaNBody()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = luaNBody(nb, i)
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
