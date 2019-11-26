package lua

import (
	lua "github.com/n-is/gopher-lua"
)

// openLib opens a library with configured functions
func openLib(ls *lua.LState, libName string) {
	var libFunc lua.LGFunction
	switch libName {
	case lua.LoadLibName:
		libFunc = lua.OpenPackage
	case lua.BaseLibName:
		libFunc = lua.OpenBase
	// Do not support Other Functions for opening through this method
	default:
		// fmt.Printf("%s Library Not Available", libName)
	}
	ls.Push(ls.NewFunction(libFunc))
	ls.Push(lua.LString(libName))
	ls.Call(1, 0)
}

// openLibraries opens libraries with all their configured functions
func openLibraries(ls *lua.LState, libNames ...string) {
	for _, libName := range libNames {
		openLib(ls, libName)
	}
}
