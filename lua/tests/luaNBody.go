package tests

import "github.com/user/lua"

func SetupLuaNBody() *lua.LuaScript {

	libs := []string{}

	l := NewLuaFromFile("scripts/lua/nbody.lua", libs...)

	return l
}

func LuaNBody(l *lua.LuaScript, n int64) float64 {

	l.SetGlobalNumber("Input", n)

	output := l.RunGetNumber()

	if out, ok := output.(float64); ok {
		return out
	} else {
		// panic("Output is not float64")
		return float64(output.(int64))
	}
}
