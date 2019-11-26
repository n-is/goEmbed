package golua

import "github.com/n-is/goEmbed/lua"

func setupLuaFib() *lua.LuaScript {

	l := newLuaFromFile("../scripts/lua/fib.lua")

	return l
}

func luaFib(l *lua.LuaScript, n int) int64 {

	l.SetGlobalNumber("Input", n)

	output := l.RunGetNumber()

	return output.(int64)
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
