package golua

import "testing"

func benchmarkLuaTree(i int64, b *testing.B) {
	var r float64
	tree := SetupLuaTree()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = LuaTree(tree, i)
	}
	resultNBody = r
}

func BenchmarkLuaTreeM20(b *testing.B) { benchmarkLuaTree(-20, b) }
func BenchmarkLuaTree0(b *testing.B)   { benchmarkLuaTree(0, b) }
func BenchmarkLuaTree2(b *testing.B)   { benchmarkLuaTree(2, b) }
func BenchmarkLuaTree4(b *testing.B)   { benchmarkLuaTree(4, b) }
func BenchmarkLuaTree8(b *testing.B)   { benchmarkLuaTree(8, b) }
func BenchmarkLuaTree10(b *testing.B)  { benchmarkLuaTree(10, b) }
func BenchmarkLuaTree15(b *testing.B)  { benchmarkLuaTree(15, b) }
func BenchmarkLuaTree20(b *testing.B)  { benchmarkLuaTree(20, b) }
