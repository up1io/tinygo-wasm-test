package main

import (
	"syscall/js"
	"time"
)

type JsFunc func(this js.Value, args []js.Value) interface{}

func main() {
	quit := make(chan bool)
	doc := js.Global().Get("document")
	if doc.IsNull() {
		panic("document is null")
	}
	doc.Call("addEventListener", "mousemove", js.FuncOf(handleMouseMove()))

	go func() {
		timer := time.Tick(100 * time.Millisecond)
		for now := range timer {
			println(now.String())
		}
	}()

	<-quit
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
