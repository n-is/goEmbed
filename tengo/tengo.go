package tengo

import (
	"errors"

	"github.com/n-is/tengo/script"
	"github.com/n-is/tengo/stdlib"
)

type TengoScript struct {
	script   *script.Script
	compiled *script.Compiled
}

func NewTengoScript(src []byte, libs ...string) *TengoScript {
	t := &TengoScript{}

	s := script.New([]byte(src))

	libraries := make([]string, len(stdlib.SourceModules)+len(libs))
	libraries = append(libraries, libs...)
	for lib := range stdlib.SourceModules {
		libraries = append(libraries, lib)
	}
	s.SetImports(stdlib.GetModuleMap(libraries...))

	t.script = s

	return t
}

func (t *TengoScript) SetGlobal(varName string, val interface{}) error {
	var v interface{}
	switch val := val.(type) {
	case uint:
		v = int64(val)
	case uint16:
		v = int64(val)
	case uint32:
		v = int64(val)
	case uint64:
		return errors.New("Unsupported Variable Type")
	default:
		v = val
	}
	err := t.script.Add(varName, v)

	return err
}

func (t *TengoScript) Compile() {

	c, err := t.script.Compile()
	if err != nil {
		panic(err)
	}
	t.compiled = c
}

func (t *TengoScript) RunGetInt() int64 {

	t.compiled.Run()
	output := t.compiled.Get("Output")
	// We expect the output of Tengo Function to be a int64
	return output.Int64()
}

func (t *TengoScript) RunGetMap() map[string]interface{} {

	t.compiled.Run()
	output := t.compiled.Get("Output")
	// We expect the output of Tengo Function to be a map[string]interface{}
	return output.Map()
}

func (t *TengoScript) RunGetString() string {

	t.compiled.Run()
	output := t.compiled.Get("Output")
	// We expect the output of Tengo Function to be a string
	return output.String()
}

func (t *TengoScript) RunGetBool() bool {

	t.compiled.Run()
	output := t.compiled.Get("Output")
	// We expect the output of Tengo Function to be a string
	return output.Bool()
}

func (t *TengoScript) RunGetFloat() float64 {

	t.compiled.Run()
	output := t.compiled.Get("Output")
	// We expect the output of Tengo Function to be a string
	return output.Float()
}
