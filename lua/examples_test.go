package lua

import (
	"fmt"

	golua "github.com/n-is/gopher-lua"
)

func ExampleNewLuaScript() {
	ls := NewLuaScript([]byte(`
	-- os.exit()
	y = math.half(12)
	return y
	`))

	ls1 := NewLuaScript([]byte(`
	val = math.sqrt(2)
	val = math.abs(val)
	return val
	`))

	availableFuncs := []string{"sqrt", "abs"}

	done := make(chan bool, 2)

	go func() {
		ls.AddLib("os", "all")

		f := func(L *golua.LState) int {
			num := float64(L.CheckNumber(1))
			L.Push(golua.LNumber(num / 2))
			return 1
		}
		ls.AddFunc("math", "half", f)

		ls.OpenLibs()
		done <- true
	}()
	go func() {
		ls1.AddLib("math", availableFuncs...)
		ls1.OpenLibs()
		done <- true
	}()

	<-done
	<-done

	val1 := ls1.RunGetNumber()

	val := ls.RunGetNumber()

	fmt.Printf("%0.3f, %d\n", val1, val)
	// Output: 1.414, 6
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

	ls.AddLib("math", availableFuncs...)
	ls.OpenLibs()

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
