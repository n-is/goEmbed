package tengo

import "fmt"

func ExampleNewTengoScript() {
	tr := NewTengoScript([]byte(`
	Output := 7
	`), "Output")

	tr.Compile()
	fmt.Println(tr.RunGetInt())

	// Output: 7
}

func ExampleTengoScript_SetGlobal() {
	tr := NewTengoScript([]byte(`
	Output := Input
	`), "Output")

	tr.SetGlobal("Input", "Hello World!!")
	tr.Compile()
	fmt.Println(tr.RunGetString())

	// Output: Hello World!!
}

func ExampleTengoScript_Compile() {
	tr := NewTengoScript([]byte(`
	Input++
	Output := Input
	`), "Output")

	tr.SetGlobal("Input", 3)

	tr.Compile()

	v1 := tr.RunGetInt()
	v2 := tr.RunGetInt()

	tr.Compile()

	v3 := tr.RunGetInt()

	fmt.Println(v1, v2, v3)

	// Output: 4 5 4
}

func ExampleTengoScript_RunGetInt() {
	tr := NewTengoScript([]byte(`
	Input++
	Output := Input
	`), "Output")

	tr.SetGlobal("Input", 3)

	tr.Compile()

	v1 := tr.RunGetInt()
	v2 := tr.RunGetInt()

	tr.Compile()

	v3 := tr.RunGetInt()

	fmt.Println(v1, v2, v3)

	// Output: 4 5 4
}

func ExampleTengoScript_RunGetMap() {
	tr := NewTengoScript([]byte(`
	Input++
	Output := {value: Input}
	`), "Output")

	tr.SetGlobal("Input", 3)

	tr.Compile()

	val := tr.RunGetMap()

	fmt.Println(val["value"])

	// Output: 4
}

func ExampleTengoScript_RunGetString() {
	tr := NewTengoScript([]byte(`
	Output := Input + " Minna"
	`), "Output")

	tr.SetGlobal("Input", "Nihao")

	tr.Compile()

	val := tr.RunGetString()

	fmt.Println(val)

	// Output: Nihao Minna
}

func ExampleTengoScript_RunGetBool() {
	tr := NewTengoScript([]byte(`
	Output := false
	`), "Output")

	tr.Compile()

	val := tr.RunGetBool()

	fmt.Println(val)

	// Output: false
}

func ExampleTengoScript_RunGetFloat() {
	tr := NewTengoScript([]byte(`
	math := import("math")

	Output := math.sqrt(2)
	`), "Output", "math")

	tr.Compile()

	val := tr.RunGetFloat()

	fmt.Printf("%0.3f\n", val)

	// Output: 1.414
}
