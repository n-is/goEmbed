package golua

import (
	"math/rand"
	"time"

	"github.com/n-is/goEmbed/lua"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func setupLuaConcat() *lua.LuaScript {

	libs := []string{}

	l := newLuaFromFile("../scripts/lua/concat.lua", libs...)

	return l
}

func luaConcat(l *lua.LuaScript, m map[string]interface{}) string {

	l.SetGlobalMap("Input", m)

	output := l.RunGetString()

	return output
}

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
