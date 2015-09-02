package main

import (
	"golang.org/x/mobile/app"

	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/touch"

	"golang.org/x/mobile/gl"
	"golang.org/x/mobile/exp/f32"
	"golang.org/x/mobile/exp/gl/glutil"

)

var (
	program  		gl.Program
	position 		gl.Attrib
	texture			gl.Attrib
	color    		gl.Uniform
	matrixId 		gl.Uniform
	resolutionId	gl.Uniform

	swasBuffer 		gl.Buffer
	quadBuffer 		gl.Buffer
	quadTexBuffer	gl.Buffer

	alpha    		float32 = 0.0
	resIndex		float32
	spin			bool
)

func main() {
	app.Main(func(a app.App) {
		var sz size.Event
		for e := range a.Events() {
			switch e := app.Filter(e).(type) {
			case lifecycle.Event:
				switch e.Crosses(lifecycle.StageVisible) {
					case lifecycle.CrossOn:
						onStart()
					case lifecycle.CrossOff:
						onStop()
				}
			case size.Event:
				sz = e
				resIndex = float32(sz.WidthPx)/float32(sz.HeightPx)
			case paint.Event:
				onPaint(sz)
				a.EndPaint(e)
			case touch.Event:
				eventType := e.Type.String()
				if eventType == "begin" {
					spin = !spin
				}
			}
		}
	})
}

func onStart() {
	array := loadShaders("vShader.vs", "fShader.vs")

	var e error
	program, e = glutil.CreateProgram(array[0], array[1])
	crash_check(e)

	quadBuffer = gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, quadBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, quadData, gl.STATIC_DRAW)

	quadTexBuffer = gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, quadTexBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, quadTexData, gl.STATIC_DRAW)

	guadBuffer2 = gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, quadBuffer2)
	gl.BufferData(gl.ARRAY_BUFFER, guadData )

	swasBuffer = gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, swasBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, swastikaData, gl.STATIC_DRAW)

	position = gl.GetAttribLocation(program, "position")
	texture  = gl.GetAttribLocation(program, "texCoords")
	color    = gl.GetUniformLocation(program, "color")
	matrixId = gl.GetUniformLocation(program, "rotationMatrix")
	resolutionId = gl.GetUniformLocation(program, "resIndex")

}

func onStop() {
	gl.DeleteProgram(program)
	gl.DeleteBuffer(swasBuffer)
	gl.DeleteBuffer(quadBuffer)
}

func onPaint(sz size.Event) {
	// Setting background
	gl.ClearColor(0.2, 0.0, 0.2, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	var rotationMatrix = []float32 {
			f32.Cos(-alpha), 	-f32.Sin(-alpha), 	0.0, // top left
			f32.Sin(-alpha), 	f32.Cos(-alpha), 	0.0, // bottom left
			0.0, 	0.0,    1.0, // bottom right
		
	}

	gl.UseProgram(program)
	// setting color
	gl.Uniform4f(color, rgb(255), rgb(255), rgb(255), 1)
	gl.UniformMatrix3fv(matrixId, rotationMatrix)
	gl.Uniform1f(resolutionId, resIndex)

	gl.BindBuffer(gl.ARRAY_BUFFER, swasBuffer)

	gl.EnableVertexAttribArray(position)
	gl.VertexAttribPointer(position, 3, gl.FLOAT, false, 0, 0)
	gl.DrawArrays(gl.LINES, 0, 16)
	gl.DisableVertexAttribArray(position)


	gl.UseProgram(program)
	// setting color
	gl.Uniform4f(color, rgb(130), rgb(50), rgb(80), 1)
	gl.Uniform1f(resolutionId, resIndex)
	gl.UniformMatrix3fv(matrixId, rotationMatrix)

	gl.BindBuffer(gl.ARRAY_BUFFER, quadBuffer)
	//gl.BindBuffer(gl.ARRAY_BUFFER, quadTexBuffer)

	gl.EnableVertexAttribArray(position)
	//gl.EnableVertexAttribArray(texture)
	gl.VertexAttribPointer(position, 3, gl.FLOAT, false, 0, 0)
	gl.DrawArrays(gl.TRIANGLES, 0, 6)
	gl.DisableVertexAttribArray(position)
	//gl.DisableVertexAttribArray(texture)

	if spin == true{
		alpha += 0.1
	}

	if alpha >= 360 {
		alpha = 0.0
	}

}
