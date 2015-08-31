

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

)

var (
	program  gl.Program
	position gl.Attrib
	//offset   gl.Uniform
	color    gl.Uniform
	buf      gl.Buffer
	swBuf    gl.Buffer

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
	/*
	
	var vertexSrc, fragmentSrc string
	program := gl.CreateProgram()

	vertexShader := gl.CreateShader(gl.VERTEX_SHADER)
	gl.ShaderSource(vertexShader, vertexSrc)
	gl.CompileShader(vertexShader)

	if gl.GetShaderi(vertexShader, gl.COMPILE_STATUS) == 0 {
		defer gl.DeleteShader(vertexShader)
		log.Printf("shader compile: %s", gl.GetShaderInfoLog(vertexShader))
	}

	fragmentShader := gl.CreateShader(gl.FRAGMENT_SHADER)
	gl.ShaderSource(fragmentShader, fragmentSrc)
	gl.CompileShader(fragmentShader)

	if gl.GetShaderi(fragmentShader, gl.COMPILE_STATUS)	== 0 {
		defer gl.DeleteShader(fragmentShader)
		log.Printf("shader compile: %s", gl.GetShaderInfoLog(fragmentShader))
	}

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)

	gl.LinkProgram(program)

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	if gl.GetProgrami(program, gl.LINK_STATUS) == 0 {
		defer gl.DeleteProgram(program)
		log.Printf("program : %s", gl.GetProgramInfoLog(program))
	}
	
	*/
	var err error
	program, err = glutil.CreateProgram(vertexSrc, fragmentSrc)
	if err != nil {
		log.Printf("error creating GL program: %v", err)
		return
	}


	// Creating buffer
	buf = gl.CreateBuffer()

	// Starting work with buffer
	/*gl.BindBuffer(gl.ARRAY_BUFFER, buf)
	gl.BufferData(gl.ARRAY_BUFFER, triangleData, gl.STATIC_DRAW)*/
	

	swBuf = gl.CreateBuffer()

	gl.BindBuffer(gl.ARRAY_BUFFER, swBuf)
	gl.BufferData(gl.ARRAY_BUFFER, swastikaData, gl.STATIC_DRAW)
	


	position = gl.GetAttribLocation(program, "position")
	//offset   = gl.GetUniformLocation(program, "offset")
	color    = gl.GetUniformLocation(program, "color")
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
	gl.Uniform4f(color, 1.0, 0.0, 0.3, 1)
	// Opening buffer
	//gl.BindBuffer(gl.ARRAY_BUFFER, buf)

	/*gl.EnableVertexAttribArray(position)
	gl.VertexAttribPointer(position, 3, gl.FLOAT, false, 0, 0) 
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
	gl.DisableVertexAttribArray(position)*/

	gl.BindBuffer(gl.ARRAY_BUFFER, swBuf)

	gl.EnableVertexAttribArray(position)
	gl.VertexAttribPointer(position, 3, gl.FLOAT, false, 0, 0)
	gl.DrawArrays(gl.LINES, 0, 16)
	gl.DisableVertexAttribArray(position)


}

var triangleData = f32.Bytes(binary.LittleEndian,
	0.0, 0.4, 0.0, // top left
	0.0, 0.0, 0.0, // bottom left
	0.4, 0.0, 0.0, // bottom right
)

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
