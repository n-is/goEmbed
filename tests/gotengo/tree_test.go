package gotengo

import (
	"testing"

	"github.com/n-is/goEmbed/tengo"
)

func Test_tengoTree(t *testing.T) {
	type args struct {
		l *tengo.TengoScript
		n int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{"Test1", args{setupTengoTree(), 5}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tengoTree(tt.args.l, tt.args.n); got != tt.want {
				t.Errorf("tengoTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

var resultTree int64

func benchmarkTengoTree(i int64, b *testing.B) {
	var r int64
	tree := setupTengoTree()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		r = tengoTree(tree, i)
	}
	resultTree = r
}

func BenchmarkTengoTreeM20(b *testing.B) { benchmarkTengoTree(-20, b) }
func BenchmarkTengoTree0(b *testing.B)   { benchmarkTengoTree(0, b) }
func BenchmarkTengoTree2(b *testing.B)   { benchmarkTengoTree(2, b) }
func BenchmarkTengoTree4(b *testing.B)   { benchmarkTengoTree(4, b) }
func BenchmarkTengoTree8(b *testing.B)   { benchmarkTengoTree(8, b) }
func BenchmarkTengoTree10(b *testing.B)  { benchmarkTengoTree(10, b) }
func BenchmarkTengoTree15(b *testing.B)  { benchmarkTengoTree(15, b) }
func BenchmarkTengoTree20(b *testing.B)  { benchmarkTengoTree(20, b) }
