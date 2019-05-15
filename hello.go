package main

import (
	"fmt"
	"sync"
	"syscall/js"
	"time"
	"wasmx/pomodoro"
)

func main() {
	fmt.Println("Hello, WebAssembly!")
	var wg sync.WaitGroup
	wg.Add(1)

	pomodoro := pomodoro.Pomodoro{}
	pomodoro.Register(setValue)

	var cb js.Func
	cb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Println("button clicked")
		pomodoro.Tick()
		//cb.Release() // release the function if the button will not be clicked again
		return nil
	})
	js.Global().Get("document").Call("getElementById", "btn0").Call("addEventListener", "click", cb)
	wg.Wait()
}

func setValue(s string) {
	js.Global().Get("document").Call("getElementById", "id0").Set("value", s)
}

func test() {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
			setValue(t.String())
		}
	}()
}
