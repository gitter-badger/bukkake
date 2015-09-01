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

	"log"
	

)

var (
	program  		gl.Program
	position 		gl.Attrib
	color    		gl.Uniform
	matrixId 		gl.Uniform
	resolutionId	gl.Uniform

	swasBuffer 		gl.Buffer
	cubeBuffer 		gl.Buffer
	indicesBuffer	gl.Buffer

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
	var err error
	program, err = glutil.CreateProgram(array[0], array[1])
	if err != nil {
		log.Printf("error creating GL program: %v", err)
		return
	}

	cubeBuffer = gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, cubeBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, cubeData, gl.STATIC_DRAW)

	indicesBuffer = gl.CreateBuffer()
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, indicesBuffer)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, cubeIndices, gl.STATIC_DRAW)

	swasBuffer = gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, swasBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, swastikaData, gl.STATIC_DRAW)

	position = gl.GetAttribLocation(program, "position")
	color    = gl.GetUniformLocation(program, "color")
	matrixId = gl.GetUniformLocation(program, "rotationMatrix")
	resolutionId = gl.GetUniformLocation(program, "resIndex")

}

func onStop() {
	gl.DeleteProgram(program)
	gl.DeleteBuffer(swasBuffer)
	gl.DeleteBuffer(cubeBuffer)
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
	gl.Uniform4f(color, 1.0, 1.0, 1.0, 1)
	gl.UniformMatrix3fv(matrixId, rotationMatrix)
	gl.Uniform1f(resolutionId, resIndex)

	gl.BindBuffer(gl.ARRAY_BUFFER, swasBuffer)

	gl.EnableVertexAttribArray(position)
	gl.VertexAttribPointer(position, 3, gl.FLOAT, false, 0, 0)
	gl.DrawArrays(gl.LINES, 0, 16)
	gl.DisableVertexAttribArray(position)


	gl.UseProgram(program)
	// setting color
	gl.Uniform4f(color, 0.0, 1.0, 1.0, 1)
	gl.Uniform1f(resolutionId, resIndex)
	gl.UniformMatrix3fv(matrixId, rotationMatrix)

	gl.BindBuffer(gl.ARRAY_BUFFER, cubeBuffer)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, indicesBuffer)

	gl.EnableVertexAttribArray(position)
	gl.VertexAttribPointer(position, 3, gl.FLOAT, false, 0, 0)
	gl.DrawElements(gl.LINES, 48, gl.BYTE, 0)
	gl.DisableVertexAttribArray(position)

	if spin == true{
		alpha += 0.1
	}

	if alpha >= 360 {
		alpha = 0.0
	}

}
