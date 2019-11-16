// Package lua provides the helper functions for loading a lua script in a
// sandbox environment, and running it multiple times.
package lua

import (
	"bytes"
	"math"

	lua "github.com/n-is/gopher-lua"
)

// LuaScript is for loading the script and make interacting with lua scripts
// easier.
type LuaScript struct {
	state        *lua.LState
	loadedScript *lua.LFunction
}

// NewLuaScript generate a new LuaScript using the source code.
// The script is loaded and saved, so it can be run multiple times, without
// loading again.
//
// Currently supported libraries(libs) are:
// 	"table"
// 	"io"
// 	"os"
// 	"string"
// 	"math"
// 	"debug"
// 	"channel"
// 	"coroutine"
func NewLuaScript(src []byte, libs ...string) *LuaScript {
	L := lua.NewState(lua.Options{
		RegistrySize:        1024 * 20,
		CallStackSize:       1024,
		IncludeGoStackTrace: true,
		SkipOpenLibs:        true,
	})

	// Load Base Library (Minimum library required to run basic scripts)
	L.OpenLibs("", "package")

	// Load the remaining libraries
	L.OpenLibs(libs...)

	source := bytes.NewBuffer(src)

	// Load the source code
	loadedScript, err := L.Load(source, "?")

	if err != nil {
		panic(err)
	}

	return &LuaScript{state: L, loadedScript: loadedScript}
}

// SetGlobalString sets a string as a global variable in the lua script.
// Running with same varName overrides the earlier value of variable.
func (l *LuaScript) SetGlobalString(varName, str string) {

	l.state.SetGlobal(varName, lua.LString(str))
}

// SetGlobalBool sets a bool as a global variable in the lua script.
// Running with same varName overrides the earlier value of variable.
func (l *LuaScript) SetGlobalBool(varName string, b bool) {

	l.state.SetGlobal(varName, lua.LBool(b))
}

// SetGlobalNumber sets a number as a global variable in the lua script.
// The number can be either a int, int32, int64, float32, float64.
// Running with same varName overrides the earlier value of variable.
func (l *LuaScript) SetGlobalNumber(varName string, num interface{}) {

	switch n := num.(type) {
	case int, int32, int64, float32, float64:
		l.state.SetGlobal(varName, luaValueOf(n))
	}
}

// SetGlobalMap sets a map as a global variable in the lua script.
// The map is represented by a table in the lua script.
// Running with same varName overrides the earlier value of variable.
func (l *LuaScript) SetGlobalMap(varName string, m map[string]interface{}) {

	tbl := createTable(l.state, m)
	l.state.SetGlobal(varName, tbl)
}

// RunGetString runs the lua script with the configured variables.
// RunGetString considers the result of the running of the script to be a string.
// It extracts the string and returns it.
func (l *LuaScript) RunGetString() string {

	l.state.Push(l.loadedScript)
	l.state.Call(0, lua.MultRet)

	output := l.state.ToString(-1)
	l.state.Pop(1)

	return output
}

// RunGetNumber runs the lua script with the configured variables.
// RunGetNumber considers the result of the running of the script to be a number.
// It extracts the number, converts it to int64 or float64, and returns it.
func (l *LuaScript) RunGetNumber() interface{} {

	l.state.Push(l.loadedScript)
	l.state.Call(0, lua.MultRet)

	output := l.state.ToNumber(-1)
	l.state.Pop(1)

	return parseLuaNumber(output)
}

// RunGetBool runs the lua script with the configured variables.
// RunGetBool considers the result of the running of the script to be a boolean.
// It extracts the boolean value and returns it.
func (l *LuaScript) RunGetBool() bool {

	l.state.Push(l.loadedScript)
	l.state.Call(0, lua.MultRet)

	output := l.state.ToBool(-1)
	l.state.Pop(1)

	return output
}

// RunGetMap runs the lua script with the configured variables.
// RunGetMap considers the result of the running of the script to be a table.
// that holds a map[string]interface{}. It extracts the map and returns it.
func (l *LuaScript) RunGetMap() map[string]interface{} {

	l.state.Push(l.loadedScript)
	l.state.Call(0, lua.MultRet)

	tbl := l.state.ToTable(-1)
	l.state.Pop(1)
	output := tableToGoMap(tbl)

	return output
}

// createTable creates a lua table from the provided map and return it.
func createTable(state *lua.LState, m map[string]interface{}) *lua.LTable {
	tbl := state.NewTable()
	for k, v := range m {
		switch v := v.(type) {
		case map[string]interface{}:
			t := createTable(state, v)
			state.RawSetString(tbl, k, t)
		default:
			state.RawSetString(tbl, k, luaValueOf(v))
		}

	}

	return tbl
}

// luaValueOf gives the corresponding lua value of given go input variable.
func luaValueOf(value interface{}) lua.LValue {
	switch value := value.(type) {
	case float64:
		return lua.LNumber(value)
	case float32:
		return lua.LNumber(float64(value))
	case string:
		return lua.LString(value)
	case int64:
		return lua.LNumber(value)
	case int32:
		return lua.LNumber(int64(value))
	case int:
		return lua.LNumber(int64(value))
	case bool:
		return lua.LBool(value)
	case lua.LValue:
		return value
	case nil:
		return lua.LNil
	}
	return lua.LNil
}

// goValueOf gives the corresponding go value of given lua input variable.
func goValueOf(val lua.LValue) interface{} {
	switch val := val.(type) {
	case lua.LString:
		return lua.LVAsString(val)
	case lua.LNumber:
		return parseLuaNumber(val)
	case lua.LBool:
		return bool(val)
	case *lua.LTable:
		return tableToGoMap(val)
	}

	return nil
}

// tableToGoMap converts a Lua Table to a map if it represents a map
func tableToGoMap(table *lua.LTable) map[string]interface{} {
	if length := table.Len(); length == 0 { // map
		gomap := make(map[string]interface{})
		table.ForEach(func(key, val lua.LValue) {
			k := goValueOf(key)

			if k, ok := k.(string); ok {
				v := goValueOf(val)
				gomap[k] = v
			}
		})
		return gomap
	}

	return nil
}

// parseLuaNumber converts a Lua Number to int64 or float64 depending upon the
// tolerance(1e-10)
func parseLuaNumber(val lua.LNumber) interface{} {
	num := float64(val)
	nearNum := math.Round(num)

	tol := float64(1e-10)

	if diff := math.Abs(num - nearNum); diff < tol {
		val := int64(nearNum)
		return val
	}

	return num
}
