// Package lua provides the helper functions for loading a lua script in a
// sandbox environment, and running it multiple times.

package lua

import (
	"reflect"
	"testing"

	lua "github.com/n-is/gopher-lua"
)

func TestLuaScript_RunGetString(t *testing.T) {
	ls := NewLuaScript([]byte(`
	value = "Hello "..Input
	return value
	`))
	tests := []struct {
		name  string
		l     *LuaScript
		input string
		want  string
	}{
		// TODO: Add test cases.
		{"Test1", ls, "Mercury", "Hello Mercury"},
		{"Test1", ls, "venus", "Hello venus"},
		{"Test1", ls, "EaRtH", "Hello EaRtH"},
		{"Test1", ls, "maRS", "Hello maRS"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ls.SetGlobalString("Input", tt.input)
			if got := tt.l.RunGetString(); got != tt.want {
				t.Errorf("LuaScript.RunGetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLuaScript_RunGetNumber(t *testing.T) {
	ls := NewLuaScript([]byte(`
	value = Input + 4.3
	return value
	`))
	tests := []struct {
		name  string
		l     *LuaScript
		input float64
		want  float64
	}{
		// TODO: Add test cases.
		{"Test1", ls, 10, 14.3},
		{"Test1", ls, 9, 13.3},
		{"Test1", ls, -12, -7.7},
		{"Test1", ls, 0, 4.3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ls.SetGlobalNumber("Input", tt.input)
			if got := tt.l.RunGetNumber(); got != tt.want {
				t.Errorf("LuaScript.RunGetNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLuaScript_RunGetMap(t *testing.T) {
	tests := []struct {
		name  string
		l     *LuaScript
		input map[string]interface{}
		want  map[string]interface{}
	}{
		// TODO: Add test cases.
		{"Test1", NewLuaScript([]byte(`
				Input.Greet.Chinese = "Nihao"
				return Input
				`)),
			map[string]interface{}{
				"Greet": map[string]interface{}{"Nepali": "Namaste"}},
			map[string]interface{}{
				"Greet": map[string]interface{}{"Nepali": "Namaste", "Chinese": "Nihao"}},
		},

		{"Test2", NewLuaScript([]byte(`
				Input.Greet.Chinese = "Nihao"
				return Input
				`)),
			map[string]interface{}{
				"Greet": map[string]interface{}{"Mars": "$2#$W#*&", "Chinese": "MoshiMoshi"}},
			map[string]interface{}{
				"Greet": map[string]interface{}{"Mars": "$2#$W#*&", "Chinese": "Nihao"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.SetGlobalMap("Input", tt.input)
			if got := tt.l.RunGetMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LuaScript.RunGetMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_luaValueOf(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name string
		args args
		want lua.LValue
	}{
		// TODO: Add test cases.
		{"TestFloat64", args{float64(12.3)}, lua.LNumber(12.3)},
		{"TestFloat32", args{float32(float64(12))}, lua.LNumber(12)},
		{"TestString", args{"Hello"}, lua.LString("Hello")},
		{"TestInt64", args{int64(-12)}, lua.LNumber(-12)},
		{"TestInt32", args{int32(-12)}, lua.LNumber(-12)},
		{"TestInt", args{int(-12)}, lua.LNumber(-12)},
		{"TestBoolTrue", args{true}, lua.LBool(true)},
		{"TestBoolFalse", args{false}, lua.LBool(false)},
		{"TestNil", args{nil}, lua.LNil},
		{"TestLValue", args{lua.LBool(false)}, lua.LBool(false)},
		{"TestOthers", args{map[string]interface{}{}}, lua.LNil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := luaValueOf(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("luaValueOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseLuaNumber(t *testing.T) {
	type args struct {
		val lua.LNumber
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
		{"TestInt", args{12}, int64(12)},
		{"TestInt", args{-12.3}, float64(-12.3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseLuaNumber(tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseLuaNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
