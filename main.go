

package main

import (
	"encoding/binary"

	"golang.org/x/mobile/app"

	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/event/lifecycle"

	"golang.org/x/mobile/gl"
	"golang.org/x/mobile/exp/f32"
	"golang.org/x/mobile/exp/gl/glutil"

	"log"
	//"math"

)

var (
	program  gl.Program
	position gl.Attrib
	//offset   gl.Uniform
	color    gl.Uniform
	matrixId    gl.Uniform
	buf      gl.Buffer
	swBuf    gl.Buffer
	rotationMatrix []float32

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
			case paint.Event:
				onPaint(sz)
				a.EndPaint(e)
			}
		}
	})
}

func onStart() {
	var err error
	program, err = glutil.CreateProgram(vertexSrc, fragmentSrc)
	if err != nil {
		log.Printf("error creating GL program: %v", err)
		return
	}

	// Creating buffer
	buf = gl.CreateBuffer()

	swBuf = gl.CreateBuffer()

	gl.BindBuffer(gl.ARRAY_BUFFER, swBuf)
	gl.BufferData(gl.ARRAY_BUFFER, swastikaData, gl.STATIC_DRAW)

	position = gl.GetAttribLocation(program, "position")
	color    = gl.GetUniformLocation(program, "color")
	matrixId = gl.GetUniformLocation(program, "rotationMatrix")

	rotationMatrix = []float32 {
		0.866, 	-0.5, 	0.0, // top left
		0.5, 	0.866, 	0.0, // bottom left
		0.0, 	0.0,    1.0, // bottom right
		
	}
}

func onStop() {
	gl.DeleteProgram(program)
	gl.DeleteBuffer(buf)
}

func onPaint(sz size.Event) {
	// Setting background
	gl.ClearColor(0.2, 0.0, 0.0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)


	gl.UseProgram(program)
	// setting color
	gl.Uniform4f(color, 1.0, 1.0, 1.0, 1)

	gl.UniformMatrix3fv(matrixId, rotationMatrix)

	gl.BindBuffer(gl.ARRAY_BUFFER, swBuf)

	gl.EnableVertexAttribArray(position)
	gl.VertexAttribPointer(position, 3, gl.FLOAT, false, 0, 0)
	gl.DrawArrays(gl.LINES, 0, 16)
	gl.DisableVertexAttribArray(position)
}

var swastikaData = f32.Bytes(binary.LittleEndian,
	0.0, -0.3, 0.0,     0.0, 0.3, 0.0,
	-0.5, -0.3, 0.0,    0.0, -0.3, 0.0,
	0.0, 0.3, 0.0,      0.5, 0.3, 0.0,


	-0.5, 0.3, 0.0,     -0.5, 0.0, 0.0,
	-0.5, 0.0, 0.0,     0.5, 0.0, 0.0,
	0.5, 0.0, 0.0,      0.5, -0.3, 0.0,
	
)

const vertexSrc= `#version 100
//uniform vec2 offset;
uniform mat3 rotationMatrix;

attribute vec4 position;
void main() {
	// offset comes in with x/y values between 0 and 1.
	// position bounds are -1 to 1.
	//vec4 offset4 = vec4(2.0*offset.x-1.0, 1.0-2.0*offset.y, 0, 0);
	gl_Position = position;
}`

const fragmentSrc = `#version 100
precision mediump float;
uniform vec4 color;
void main() {
	gl_FragColor = color;
}`
