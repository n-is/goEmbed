package gotengo

import (
	"github.com/n-is/goEmbed/tengo"
)

func setupTengoNBody() *tengo.TengoScript {

	libs := []string{"math"}

	l := newTengoFromFile("../scripts/tengo/nbody.tengo", libs...)

	return l
}
func tengoNBody(l *tengo.TengoScript, n int64) float64 {

	l.SetGlobal("Input", n)
	l.Compile()

	output := l.RunGetFloat()

	return output
}
