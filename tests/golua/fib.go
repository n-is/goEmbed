package golua

import "github.com/n-is/goEmbed/lua"

func setupLuaFib() *lua.LuaScript {

	libs := []string{}

	l := newLuaFromFile("../scripts/lua/fib.lua", libs...)

	return l
}

func luaFib(l *lua.LuaScript, n int) int64 {

	l.SetGlobalNumber("Input", n)

	output := l.RunGetNumber()

	if out, ok := output.(int64); ok {
		return out
	}
	// panic("Output is not int64")
	return int64(output.(float64))
}

func fib(n int) int64 {

	if n < 0 {
		return -1
	}

	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	n0, n1 := int64(0), int64(1)

	for i := 2; i <= n; i++ {
		tmp := n0 + n1
		n0 = n1
		n1 = tmp
	}

	return n1
}
