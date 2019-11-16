package gotengo

import "testing"

func TestTengoConcat(t *testing.T) {
	type args struct {
		m map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Nums", args{m: map[string]interface{}{"first": "3", "last": "4"}}, "34"},
		{"Name", args{m: map[string]interface{}{"first": "Nischal", "last": "Nepal"}}, "NischalNepal"},
		{"Greet", args{m: map[string]interface{}{"first": "Hello", "last": "World!!"}}, "HelloWorld!!"},
	}
	l := setupTengoConcat()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if len(randString(10)) != 10 {
				t.Error("Length of generated strings incorrect")
			}
			if got := tengoConcat(l, tt.args.m); got != tt.want {
				t.Errorf("tengoConcat() = %v, want %v", got, tt.want)
			}
		})
	}
}

var resultConcat string

func benchmarkTengoConcat(first, last string, b *testing.B) {
	var r string
	l := setupTengoConcat()
	m := map[string]interface{}{"first": first, "last": last}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = tengoConcat(l, m)
	}
	resultConcat = r
}

func BenchmarkTengoConcatNums(b *testing.B)  { benchmarkTengoConcat("3", "4", b) }
func BenchmarkTengoConcatName(b *testing.B)  { benchmarkTengoConcat("Nischal", "Nepal", b) }
func BenchmarkTengoConcatGreet(b *testing.B) { benchmarkTengoConcat("Hello", "World!!", b) }
func BenchmarkTengoConcat10(b *testing.B) {
	first := randString(10)
	last := randString(10)
	benchmarkTengoConcat(first, last, b)
}
func BenchmarkTengoConcat100(b *testing.B) {
	first := randString(100)
	last := randString(100)
	benchmarkTengoConcat(first, last, b)
}
func BenchmarkTengoConcat1000(b *testing.B) {
	first := randString(1000)
	last := randString(1000)
	benchmarkTengoConcat(first, last, b)
}
func BenchmarkTengoConcat10000(b *testing.B) {
	first := randString(10000)
	last := randString(10000)
	benchmarkTengoConcat(first, last, b)
}
func BenchmarkTengoConcat100000(b *testing.B) {
	first := randString(100000)
	last := randString(100000)
	benchmarkTengoConcat(first, last, b)
}

func BenchmarkTengoConcat1000000(b *testing.B) {
	first := randString(1000000)
	last := randString(1000000)
	benchmarkTengoConcat(first, last, b)
}

func BenchmarkTengoConcat10000000(b *testing.B) {
	first := randString(10000000)
	last := randString(10000000)
	benchmarkTengoConcat(first, last, b)
}

func BenchmarkTengoConcat100000000(b *testing.B) {
	first := randString(100000000)
	last := randString(100000000)
	benchmarkTengoConcat(first, last, b)
}

func BenchmarkTengoConcat200000000(b *testing.B) {
	first := randString(200000000)
	last := randString(200000000)
	benchmarkTengoConcat(first, last, b)
}

func BenchmarkTengoConcat400000000(b *testing.B) {
	first := randString(400000000)
	last := randString(400000000)
	benchmarkTengoConcat(first, last, b)
}

func BenchmarkTengoConcat600000000(b *testing.B) {
	first := randString(600000000)
	last := randString(600000000)
	benchmarkTengoConcat(first, last, b)
}

func BenchmarkTengoConcat900000000(b *testing.B) {
	first := randString(900000000)
	last := randString(900000000)
	benchmarkTengoConcat(first, last, b)
}
