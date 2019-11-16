package gotengo

import (
	"math/rand"
	"time"

	"github.com/n-is/goEmbed/tengo"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func setupTengoConcat() *tengo.TengoScript {

	libs := []string{}

	l := newTengoFromFile("../scripts/tengo/concat.tengo", libs...)

	return l
}

func tengoConcat(t *tengo.TengoScript, m map[string]interface{}) string {

	t.SetGlobal("Input", m)
	t.Compile()

	output := t.RunGetString()

	return output
}

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
