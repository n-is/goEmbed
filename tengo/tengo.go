// Package tengo provides the helper functions for loading a tengo script in a
// sandbox environment, and running it multiple times.
package tengo

import (
	"errors"

	"github.com/n-is/tengo/script"
	"github.com/n-is/tengo/stdlib"
)

// TengoScript is for loading the script and make interacting with tengo scripts
// easier.
type TengoScript struct {
	script    *script.Script
	compiled  *script.Compiled
	outputVar string
}

// NewTengoScript generate a new TengoScript using the source code.
// The script is loaded and saved, so it can be run multiple times, without
// loading again.
//
// The parameter `outVar` determines the exported variable from the Tengo
// script that will be extracted by the golang.
//
// Currently supported libraries(libs) are:
// 	"math"
// 	"os"
// 	"text"
// 	"times"
// 	"rand"
// 	"fmt"
// 	"json"
// 	"base64"
// 	"hex"
func NewTengoScript(src []byte, outVar string, libs ...string) *TengoScript {
	t := &TengoScript{}

	s := script.New([]byte(src))

	libraries := make([]string, len(stdlib.SourceModules)+len(libs))
	libraries = append(libraries, libs...)
	for lib := range stdlib.SourceModules {
		libraries = append(libraries, lib)
	}
	s.SetImports(stdlib.GetModuleMap(libraries...))

	t.script = s
	t.outputVar = outVar

	return t
}

// SetGlobal sets a variable as a global variable in the tengo script.
// Running with same varName overrides the earlier value of variable.
// Returns error for unsupported types and if the tengo was not able
// to add the variable to the script.
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

// Compile compiles the loaded tengo script along with the provided global
// variables.
// Return error if the script could not be compiled.
func (t *TengoScript) Compile() error {

	c, err := t.script.Compile()
	if err != nil {
		return err
	}
	t.compiled = c
	return nil
}

// RunGetInt runs the tengo script with the configured variables.
// RunGetInt considers the result of the running of the script to be an int64.
// It extracts the integer, converts it to int64, and returns it.
func (t *TengoScript) RunGetInt() int64 {

	t.compiled.Run()
	output := t.compiled.Get(t.outputVar)
	// We expect the output of Tengo Function to be a int64
	return output.Int64()
}

// RunGetMap runs the tengo script with the configured variables.
// RunGetMap considers the result of the running of the script to be a
// map[string]interface{}.
// It extracts the map, converts it to map[string]interface{}, and returns it.
func (t *TengoScript) RunGetMap() map[string]interface{} {

	t.compiled.Run()
	output := t.compiled.Get(t.outputVar)
	// We expect the output of Tengo Function to be a map[string]interface{}
	return output.Map()
}

// RunGetString runs the tengo script with the configured variables.
// RunGetString considers the result of the running of the script to be a string.
// It extracts the value, converts it to string, and returns it.
func (t *TengoScript) RunGetString() string {

	t.compiled.Run()
	output := t.compiled.Get(t.outputVar)
	// We expect the output of Tengo Function to be a string
	return output.String()
}

// RunGetBool runs the tengo script with the configured variables.
// RunGetBool considers the result of the running of the script to be a bool.
// It extracts the value, converts it to bool, and returns it.
func (t *TengoScript) RunGetBool() bool {

	t.compiled.Run()
	output := t.compiled.Get(t.outputVar)
	// We expect the output of Tengo Function to be a bool
	return output.Bool()
}

// RunGetFloat runs the tengo script with the configured variables.
// RunGetFloat considers the result of the running of the script to be a float64.
// It extracts the value, converts it to float64, and returns it.
func (t *TengoScript) RunGetFloat() float64 {

	t.compiled.Run()
	output := t.compiled.Get(t.outputVar)
	// We expect the output of Tengo Function to be a float64
	return output.Float()
}
