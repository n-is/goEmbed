package golua

import "testing"

func TestLuaConcat(t *testing.T) {
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
	l := SetupLuaConcat()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LuaConcat(l, tt.args.m); got != tt.want {
				t.Errorf("LuaConcat() = %v, want %v", got, tt.want)
			}
		})
	}
}

var resultConcat string

func benchmarkLuaConcat(first, last string, b *testing.B) {
	var r string
	l := SetupLuaConcat()
	m := map[string]interface{}{"first": first, "last": last}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = LuaConcat(l, m)
	}
	resultConcat = r
}

func BenchmarkLuaConcatNums(b *testing.B)  { benchmarkLuaConcat("3", "4", b) }
func BenchmarkLuaConcatName(b *testing.B)  { benchmarkLuaConcat("Nischal", "Nepal", b) }
func BenchmarkLuaConcatGreet(b *testing.B) { benchmarkLuaConcat("Hello", "World!!", b) }
func BenchmarkLuaConcat10(b *testing.B) {
	first := RandString(10)
	last := RandString(10)
	benchmarkLuaConcat(first, last, b)
}
func BenchmarkLuaConcat100(b *testing.B) {
	first := RandString(100)
	last := RandString(100)
	benchmarkLuaConcat(first, last, b)
}
func BenchmarkLuaConcat1000(b *testing.B) {
	first := RandString(1000)
	last := RandString(1000)
	benchmarkLuaConcat(first, last, b)
}
func BenchmarkLuaConcat10000(b *testing.B) {
	first := RandString(10000)
	last := RandString(10000)
	benchmarkLuaConcat(first, last, b)
}
func BenchmarkLuaConcat100000(b *testing.B) {
	first := RandString(100000)
	last := RandString(100000)
	benchmarkLuaConcat(first, last, b)
}

func BenchmarkLuaConcat1000000(b *testing.B) {
	first := RandString(1000000)
	last := RandString(1000000)
	benchmarkLuaConcat(first, last, b)
}

func BenchmarkLuaConcat10000000(b *testing.B) {
	first := RandString(10000000)
	last := RandString(10000000)
	benchmarkLuaConcat(first, last, b)
}

func BenchmarkLuaConcat100000000(b *testing.B) {
	first := RandString(100000000)
	last := RandString(100000000)
	benchmarkLuaConcat(first, last, b)
}

func BenchmarkLuaConcat200000000(b *testing.B) {
	first := RandString(200000000)
	last := RandString(200000000)
	benchmarkLuaConcat(first, last, b)
}

func BenchmarkLuaConcat400000000(b *testing.B) {
	first := RandString(400000000)
	last := RandString(400000000)
	benchmarkLuaConcat(first, last, b)
}

func BenchmarkLuaConcat600000000(b *testing.B) {
	first := RandString(600000000)
	last := RandString(600000000)
	benchmarkLuaConcat(first, last, b)
}

func BenchmarkLuaConcat900000000(b *testing.B) {
	first := RandString(900000000)
	last := RandString(900000000)
	benchmarkLuaConcat(first, last, b)
}
