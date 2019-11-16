package golua

import "github.com/n-is/goEmbed/lua"

func setupLuaTree() *lua.LuaScript {

	libs := []string{}

	l := newLuaFromFile("../scripts/lua/tree.lua", libs...)

	return l
}

func luaTree(l *lua.LuaScript, n int64) int64 {

	l.SetGlobalNumber("Input", n)

	output := l.RunGetNumber()
	// panic("Output is not float64")
	return output.(int64)
}
