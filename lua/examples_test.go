package lua

import "fmt"

func ExampleNewLuaScript() {
	ls := NewLuaScript([]byte(`
	n = math.abs(-12)
	return n
	`))

	availableFuncs := []string{"pow", "abs"}

	ls1 := NewLuaScript([]byte(`
	val = math.sqrt(2)
	val = math.abs(val)
	return val
	`))

	availableFuncs = []string{"sqrt", "abs"}

	done := make(chan bool, 2)

	go func() {
		ls.OpenLib("math", availableFuncs...)
		done <- true
	}()
	go func() {
		ls1.OpenLib("math", availableFuncs...)
		done <- true
	}()

	<-done
	<-done

	val1 := ls1.RunGetNumber()

	val := ls.RunGetNumber()

	fmt.Printf("%0.3f, %d\n", val1, val)
	// Output: 1.414, 12
}

func ExampleLuaScript_RunGetString() {
	ls := NewLuaScript([]byte(`
	name = "Lua Script"
	return name
	`))

	val := ls.RunGetString()
	fmt.Println(val)
	// Output: Lua Script
}

func ExampleLuaScript_RunGetNumber() {
	ls := NewLuaScript([]byte(`
	pallindrome = 1234321
	return pallindrome
	`))

	val := ls.RunGetNumber()
	fmt.Println(val)
	// Output: 1234321
}

func ExampleLuaScript_RunGetBool() {
	ls := NewLuaScript([]byte(`
	tired = true
	return tired
	`))

	val := ls.RunGetBool()
	fmt.Println(val)
	// Output: true
}

func ExampleLuaScript_RunGetMap() {
	ls := NewLuaScript([]byte(`
	tbl = {}
	tbl["type"] = "LuaScript"
	tbl["method"] = "RunGetMap"
	return tbl
	`))

	val := ls.RunGetMap()
	fmt.Println(val["type"], val["method"])
	// Output: LuaScript RunGetMap
}

func ExampleLuaScript_SetGlobalBool() {
	ls := NewLuaScript([]byte(`
	return Input
	`))

	ls.SetGlobalBool("Input", false)
	val := ls.RunGetBool()
	fmt.Println(val)
	// Output: false
}

func ExampleLuaScript_SetGlobalNumber() {
	ls := NewLuaScript([]byte(`
	return math.floor(Input)
	`))

	availableFuncs := []string{"floor", "pow", "abs"}

	ls.OpenLib("math", availableFuncs...)

	ls.SetGlobalNumber("Input", 12.1)
	val := ls.RunGetNumber()
	fmt.Println(val)
	// Output: 12
}

func ExampleLuaScript_SetGlobalString() {
	ls := NewLuaScript([]byte(`
	return "Input: "..Input
	`))

	ls.SetGlobalString("Input", "Hola Amigo")
	ls.SetGlobalString("Input", "Ohayo Minna")
	val := ls.RunGetString()
	fmt.Println(val)
	// Output: Input: Ohayo Minna
}

func ExampleLuaScript_SetGlobalMap() {
	ls := NewLuaScript([]byte(`
	Input["value"] = Input["value"] + 32
	return Input
	`))

	m := map[string]interface{}{"value": 10}
	ls.SetGlobalMap("Input", m)
	val := ls.RunGetMap()
	fmt.Println(val["value"])
	// Output: 42
}
