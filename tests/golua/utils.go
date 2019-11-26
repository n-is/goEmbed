package golua

import (
	"io/ioutil"

	"github.com/n-is/goEmbed/lua"
)

// newLuaFromFile creates a LuaScript from a lua script in a file
func newLuaFromFile(fileName string) *lua.LuaScript {

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	// Add final line to obtain output of the script
	lastLine := []byte("\nreturn Test(Input)")
	data = append(data, lastLine...)

	return lua.NewLuaScript(data)
}
