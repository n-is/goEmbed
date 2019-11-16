package golua

import (
	"testing"

	"github.com/n-is/goEmbed/lua"
)

func Test_luaTree(t *testing.T) {
	type args struct {
		l *lua.LuaScript
		n int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{"Test1", args{setupLuaTree(), 5}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := luaTree(tt.args.l, tt.args.n); got != tt.want {
				t.Errorf("luaTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

var resultTree int64

func benchmarkLuaTree(i int64, b *testing.B) {
	var r int64
	tree := setupLuaTree()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = luaTree(tree, i)
	}
	resultTree = r
}

func BenchmarkLuaTreeM20(b *testing.B) { benchmarkLuaTree(-20, b) }
func BenchmarkLuaTree0(b *testing.B)   { benchmarkLuaTree(0, b) }
func BenchmarkLuaTree2(b *testing.B)   { benchmarkLuaTree(2, b) }
func BenchmarkLuaTree4(b *testing.B)   { benchmarkLuaTree(4, b) }
func BenchmarkLuaTree8(b *testing.B)   { benchmarkLuaTree(8, b) }
func BenchmarkLuaTree10(b *testing.B)  { benchmarkLuaTree(10, b) }
func BenchmarkLuaTree15(b *testing.B)  { benchmarkLuaTree(15, b) }
func BenchmarkLuaTree20(b *testing.B)  { benchmarkLuaTree(20, b) }
