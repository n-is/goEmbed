package tests

import (
	"io/ioutil"

	"github.com/n-is/lua-helper/lua"
)

func NewLuaFromFile(fileName string, libs ...string) *lua.LuaScript {

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	// Add final line to obtain output of the script
	lastLine := []byte("\nreturn Test(Input)")
	data = append(data, lastLine...)

	return lua.NewLuaScript(data, libs...)
}
