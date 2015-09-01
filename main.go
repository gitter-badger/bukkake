
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
	program  		gl.Program
	position 		gl.Attrib
	//offset   		gl.Uniform
	color    		gl.Uniform
	matrixId 		gl.Uniform
	resolutionId	gl.Uniform
	buf      		gl.Buffer
	swBuf    		gl.Buffer
	alpha    		float32 = 0.0
	resolution 		size.Event
	//rotationMatrix []float32

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
	gl.BindBuffer(gl.ARRAY_BUFFER, buf)
	gl.BufferData(gl.ARRAY_BUFFER, triangleData, gl.STATIC_DRAW)

	swBuf = gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, swBuf)
	gl.BufferData(gl.ARRAY_BUFFER, swastikaData, gl.STATIC_DRAW)

	position = gl.GetAttribLocation(program, "position")
	color    = gl.GetUniformLocation(program, "color")
	matrixId = gl.GetUniformLocation(program, "rotationMatrix")
	//resolutionId = gl.GetUniformLocation(program, "resIndex")chee

}

func onStop() {
	gl.DeleteProgram(program)
	gl.DeleteBuffer(buf)
}

func onPaint(sz size.Event) {
	// Setting background
	gl.ClearColor(0.2, 0.0, 0.0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	var rotationMatrix = []float32 {
			f32.Cos(alpha), 	-f32.Sin(alpha), 	0.0, // top left
			f32.Sin(alpha), 	f32.Cos(alpha), 	0.0, // bottom left
			0.0, 	0.0,    1.0, // bottom right
		
	}

	gl.UseProgram(program)
	// setting color
	gl.Uniform4f(color, 1.0, 1.0, 1.0, 1)
	gl.UniformMatrix3fv(matrixId, rotationMatrix)
	//gl.Uniform1f(resolutionId, resIndex)

	gl.BindBuffer(gl.ARRAY_BUFFER, swBuf)

	gl.EnableVertexAttribArray(position)
	gl.VertexAttribPointer(position, 3, gl.FLOAT, false, 0, 0)
	gl.DrawArrays(gl.LINES, 0, 16)
	gl.DisableVertexAttribArray(position)


	gl.UseProgram(program)
	// setting color
	gl.Uniform4f(color, 0.0, 1.0, 1.0, 1)

	gl.UniformMatrix3fv(matrixId, rotationMatrix)

	gl.BindBuffer(gl.ARRAY_BUFFER, buf)

	gl.EnableVertexAttribArray(position)
	gl.VertexAttribPointer(position, 3, gl.FLOAT, false, 0, 0)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
	gl.DisableVertexAttribArray(position)

	alpha += 0.01


}

var swastikaData = f32.Bytes(binary.LittleEndian,
	0.0, -0.5, 0.0,     0.0, 0.5, 0.0,
	-0.5, -0.5, 0.0,    0.0, -0.5, 0.0,
	0.0, 0.5, 0.0,      0.5, 0.5, 0.0,


	-0.5, 0.5, 0.0,     -0.5, 0.0, 0.0,
	-0.5, 0.0, 0.0,     0.5, 0.0, 0.0,
	0.5, 0.0, 0.0,      0.5, -0.5, 0.0,
	
)

//var resIndex = float32(resolution.WidthPx/resolution.HeightPx)

var triangleData = f32.Bytes(binary.LittleEndian,
	0.0, 0.4, 0.0, // top left
	0.0, 0.0, 0.0, // bottom left
	0.4, 0.0, 0.0, // bottom right
)

const vertexSrc= `#version 120
//uniform vec2 offset;
uniform mat3 rotationMatrix;
//uniform float resIndex;

attribute vec4 position;
void main() {
	// offset comes in with x/y values between 0 and 1.
	// position bounds are -1 to 1.
	//vec4 offset4 = vec4(2.0*offset.x-1.0, 1.0-2.0*offset.y, 0, 0);
	vec3 pos = rotationMatrix * position.xyz;
	gl_Position = vec4(pos.xyz, 1.0);
}`

const fragmentSrc = `#version 120
precision mediump float;
uniform vec4 color;
void main() {
	gl_FragColor = color;
}`