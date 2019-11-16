package gotengo

import (
	"io/ioutil"

	"github.com/n-is/goEmbed/tengo"
)

// newTengoFromFile creates a TengoScript from a tengo script in a file
func newTengoFromFile(fileName string, libs ...string) *tengo.TengoScript {

	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	// Add final line to obtain output of the script
	lastLine := []byte("\nOutput := Test(Input)")
	data = append(data, lastLine...)

	return tengo.NewTengoScript(data, "Output", libs...)
}
