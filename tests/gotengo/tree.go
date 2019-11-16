package gotengo

import "github.com/n-is/goEmbed/tengo"

func setupTengoTree() *tengo.TengoScript {

	libs := []string{"math"}

	l := newTengoFromFile("../scripts/tengo/tree.tengo", libs...)

	return l
}

func tengoTree(l *tengo.TengoScript, n int64) int64 {

	l.SetGlobal("Input", n)
	l.Compile()

	output := l.RunGetInt()

	return output
}
