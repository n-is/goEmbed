package gotengo

import (
	"github.com/n-is/goEmbed/tengo"
)

func setupTengoFib() *tengo.TengoScript {

	libs := []string{}

	t := newTengoFromFile("../scripts/tengo/fib.tengo", libs...)

	return t
}

func tengoFib(t *tengo.TengoScript, n int) int64 {

	t.SetGlobal("Input", n)
	t.Compile()

	output := t.RunGetInt()

	return output
}
