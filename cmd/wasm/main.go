package main

import (
	"syscall/js"
)

type JsFunc func(this js.Value, args []js.Value) interface{}

func main() {
	quit := make(chan bool)
	doc := js.Global().Get("document")
	if doc.IsNull() {
		panic("document is null")
	}
	doc.Call("addEventListener", "mousemove", js.FuncOf(handleMouseMove()))
	js.Global().Set("update", js.FuncOf(handleUpdate()))

	<-quit
}

func handleUpdate() JsFunc {
	return func(_ js.Value, _ []js.Value) interface{} {
		println("update program")
		return nil
	}
}

func handleMouseMove() JsFunc {
	return func(_ js.Value, args []js.Value) interface{} {
		event := args[0]
		x := event.Get("clientX").Int()
		y := event.Get("clientY").Int()
		println(x, y)
		return nil
	}
}
