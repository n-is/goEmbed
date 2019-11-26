package golua

import "github.com/n-is/goEmbed/lua"

func setupLuaNBody() *lua.LuaScript {

	l := newLuaFromFile("../scripts/lua/nbody.lua")

	return l
}

func luaNBody(l *lua.LuaScript, n int64) float64 {

	l.SetGlobalNumber("Input", n)

	output := l.RunGetNumber()

	if out, ok := output.(float64); ok {
		return out
	}
	// panic("Output is not float64")
	return float64(output.(int64))
}
