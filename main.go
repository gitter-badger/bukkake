package main

import (
	"golang.org/x/mobile/app"

	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/event/size"
	"golang.org/x/mobile/event/touch"

	"golang.org/x/mobile/exp/f32"
	"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/gl"
)

var (
	program      gl.Program
	position     gl.Attrib
	color        gl.Uniform
	matrixId     gl.Uniform
	resolutionId gl.Uniform

	swasBuffer gl.Buffer

	quadBuffer    gl.Buffer
	quadTexBuffer gl.Buffer

	textureId gl.Texture

	alpha    float32 = 0.0
	resIndex float32
	spin     bool

	texProgram    gl.Program
	position2     gl.Attrib
	textureCoords gl.Attrib
	matrixId2     gl.Uniform
	resolutionId2 gl.Uniform
	color2        gl.Uniform
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
				resIndex = float32(sz.WidthPx) / float32(sz.HeightPx)
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
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.ONE, gl.ONE_MINUS_SRC_ALPHA)

	array := loadShaders("vShader.vs", "fShader.vs")
	array2 := loadShaders("vTexShader.vs", "fTexShader.vs")
	pic := loadImages("495.png")

	textureId = gl.CreateTexture()
	gl.BindTexture(gl.TEXTURE_2D, textureId)

	gl.TexImage2D(gl.TEXTURE_2D, 0, 256, 256, gl.RGBA, gl.UNSIGNED_BYTE, pic[0])

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)

	var e error
	program, e = glutil.CreateProgram(array[0], array[1])
	crash_check(e)

	texProgram, e = glutil.CreateProgram(array2[0], array2[1])
	crash_check(e)

	quadBuffer = gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, quadBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, quadData, gl.STATIC_DRAW)

	quadTexBuffer = gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, quadTexBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, quadTexData, gl.STATIC_DRAW)

	swasBuffer = gl.CreateBuffer()
	gl.BindBuffer(gl.ARRAY_BUFFER, swasBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, swastikaData, gl.STATIC_DRAW)

	position = gl.GetAttribLocation(program, "position")
	color = gl.GetUniformLocation(program, "color")
	matrixId = gl.GetUniformLocation(program, "rotationMatrix")
	resolutionId = gl.GetUniformLocation(program, "resIndex")

	position2 = gl.GetAttribLocation(texProgram, "position")
	textureCoords = gl.GetAttribLocation(texProgram, "texCoords")
	matrixId2 = gl.GetUniformLocation(texProgram, "rotationMatrix")
	resolutionId2 = gl.GetUniformLocation(texProgram, "resIndex")
	color2 = gl.GetUniformLocation(texProgram, "color")

}

func onStop() {
	gl.DeleteProgram(program)
	gl.DeleteBuffer(swasBuffer)
	//gl.DeleteBuffer(quadBuffer)
}

func onPaint(sz size.Event) {
	// Setting background
	gl.ClearColor(0.2, 0.0, 0.2, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT)

	var rotationMatrix = []float32{
		f32.Cos(-alpha), -f32.Sin(-alpha), 0.0, // top left
		f32.Sin(-alpha), f32.Cos(-alpha), 0.0, // bottom left
		0.0, 0.0, 1.0, // bottom right

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

	gl.UseProgram(texProgram)
	// setting color
	gl.Uniform4f(color2, rgb(130), rgb(50), rgb(80), 1)
	gl.Uniform1f(resolutionId2, resIndex)
	gl.UniformMatrix3fv(matrixId2, rotationMatrix)

	gl.BindBuffer(gl.ARRAY_BUFFER, quadBuffer)
	gl.EnableVertexAttribArray(position2)
	gl.VertexAttribPointer(position2, 3, gl.FLOAT, false, 0, 0)

	gl.BindBuffer(gl.ARRAY_BUFFER, quadTexBuffer)
	gl.EnableVertexAttribArray(textureCoords)
	gl.VertexAttribPointer(textureCoords, 2, gl.FLOAT, false, 0, 0)

	gl.Uniform1i(gl.GetUniformLocation(texProgram, "myTexture"), 0)
	gl.ActiveTexture(gl.TEXTURE0)
	gl.BindTexture(gl.TEXTURE_2D, textureId)

	gl.DrawArrays(gl.TRIANGLES, 0, 6)
	gl.DisableVertexAttribArray(position2)
	gl.DisableVertexAttribArray(textureCoords)

	if spin == true {
		alpha += 0.1
	}

	if alpha >= 360 {
		alpha = 0.0
	}

}
