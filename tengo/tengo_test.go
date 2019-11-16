package tengo

import (
	"testing"
)

func TestTengoScript_RunGetInt(t *testing.T) {
	tr := NewTengoScript([]byte(`
	fmt := import("fmt")

	fmt.println(Input)
	Output := Input
	`), "Output", "fmt")
	tests := []struct {
		name  string
		tr    *TengoScript
		input interface{}
		want  int64
	}{
		// TODO: Add test cases.
		{"TestInt64", tr, int64(12), 12},
		{"TestUint", tr, uint(13), 13},
		{"TestUint16", tr, uint16(14), 14},
		{"TestUint32", tr, uint32(15), 15},
		{"TestUnsupportedType", tr, uint64(15), 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.tr.SetGlobal("Input", int64(1))
			err = tt.tr.SetGlobal("Input", tt.input)
			if err != nil {
				_, ok1 := tt.input.(uint64)
				ok := ok1
				if !ok {
					t.Error("No Error Shown for Unsupported Types")
				}
			}
			tt.tr.Compile()
			if got := tt.tr.RunGetInt(); got != tt.want {
				t.Errorf("TengoScript.RunGetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestTengoScript_RunGetMap(t *testing.T) {
// 	tr := NewTengoScript([]byte(`
// 	fmt := import("fmt")
// 	val := 42

// 	Tengo := func(m) {
// 		m["Greet"] = "Hello"

// 		if m["value"] == undefined {
// 			m["value"] = 0
// 		}
// 		m["value"] += val

// 		return m
// 	}

// 	Output := Tengo(Input)
// 	`), "fmt")
// 	tests := []struct {
// 		name  string
// 		tr    *TengoScript
// 		input map[string]interface{}
// 		want  map[string]interface{}
// 	}{
// 		// TODO: Add test cases.
// 		{"Test1", tr, map[string]interface{}{},
// 			map[string]interface{}{"Greet": "Hello", "value": int64(42)}},
// 		{"Test2", tr, map[string]interface{}{"Greet": "Nihao"},
// 			map[string]interface{}{"Greet": "Hello", "value": int64(42)}},
// 		{"Test3", tr, map[string]interface{}{"Greet": "Ohayo", "value": 10},
// 			map[string]interface{}{"Greet": "Hello", "value": int64(52)}},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.tr.SetGlobal("Input", tt.input)
// 			tt.tr.Compile()
// 			if got := tt.tr.RunGetMap(); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("TengoScript.RunGetMap() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestTengoScript_RunGetString(t *testing.T) {
// 	tr := NewTengoScript([]byte(`

// 	Output := Input1 + Input2
// 	`))
// 	tests := []struct {
// 		name   string
// 		tr     *TengoScript
// 		input1 string
// 		input2 string
// 		want   string
// 	}{
// 		// TODO: Add test cases.
// 		{"Test1", tr, "Hello", "Earth", "HelloEarth"},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.tr.SetGlobal("Input1", tt.input1)
// 			tt.tr.SetGlobal("Input2", tt.input2)
// 			tt.tr.Compile()
// 			if got := tt.tr.RunGetString(); got != tt.want {
// 				t.Errorf("TengoScript.RunGetString() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestTengoScript_RunGetBool(t *testing.T) {
// 	tr := NewTengoScript([]byte(`
// 	Equals := func(x, y) {
// 		return x == y
// 	}

// 	Output := Equals(Input1, Input2)
// 	`))
// 	tests := []struct {
// 		name string
// 		tr   *TengoScript
// 		in1  interface{}
// 		in2  interface{}
// 		want bool
// 	}{
// 		// TODO: Add test cases.
// 		{"TestStringsTrue", tr, "Name", "Name", true},
// 		{"TestStringsFalse", tr, "Name", "Hero", false},
// 		{"TestNums1", tr, float64(12), int64(12), false},
// 		{"TestNums2", tr, int64(12), int64(12), true},
// 		{"TestNums3", tr, -12.3, -12.3, true},
// 		{"TestBools1", tr, false, false, true},
// 		{"TestBools2", tr, false, true, false},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.tr.SetGlobal("Input1", tt.in1)
// 			tt.tr.SetGlobal("Input2", tt.in2)
// 			tt.tr.Compile()
// 			if got := tt.tr.RunGetBool(); got != tt.want {
// 				t.Errorf("TengoScript.RunGetBool() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestTengoScript_RunGetFloat(t *testing.T) {
// 	tr := NewTengoScript([]byte(`
// 	math := import("math")

// 	Output := math.abs(Input)
// 	`), "math")
// 	tests := []struct {
// 		name  string
// 		tr    *TengoScript
// 		input float64
// 		want  float64
// 	}{
// 		// TODO: Add test cases.
// 		{"Test1", tr, float64(-12.34), 12.34},
// 		{"Test2", tr, float64(1.234), 1.234},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.tr.SetGlobal("Input", tt.input)
// 			tt.tr.Compile()
// 			if got := tt.tr.RunGetFloat(); got != tt.want {
// 				t.Errorf("TengoScript.RunGetFloat() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
