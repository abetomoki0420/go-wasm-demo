package main

import (
	"syscall/js"
	"wasm/src/math"
)

func add() js.Func {
	sumFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		// Converts js.Value to float64
		num1 := args[0].Float()
		num2 := args[1].Float()

		result := math.AddCore(num1, num2)

		return result
	})

	return sumFunc
}

func sum() js.Func {
	sumFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		floatArgs := make([]float64, args[0].Get("length").Int())
		for i := 0; i < args[0].Get("length").Int(); i++ {
			floatArgs[i] = args[0].Index(i).Float()
		}

		result := math.SumCore(floatArgs)

		return result
	})

	return sumFunc
}

func main() {
	js.Global().Set("add", add())
	js.Global().Set("sum", sum())
	<-make(chan bool)
}
