
package main

import (
	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"

	"golang.org/x/mobile/gl"
)

func main() {
	app.Main(func(a app.App) {
		var sz size.Event
		for e := range a.Events() {
			switch e := app.Filter(e).(type) {
			case size.Event:
				sz = e
			case paint.Event:
				onDraw(sz)
				a.EndPaint(e)
			}
		}
	})
}

func onDraw(sz size.Event) {
	gl.ClearColor(1, 0, 0.6, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}