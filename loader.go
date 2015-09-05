package main

import (
	"golang.org/x/mobile/asset"

	"image"
	"image/draw"
	_ "image/png"

	"golang.org/x/mobile/exp/gl/glutil"
	"golang.org/x/mobile/gl"
	"io/ioutil"
)

func loadImages(path string) *image.RGBA {
	// allocating image holder
	var img image.Image

	imgFile, e := asset.Open(path)
	def_check(e)
	defer imgFile.Close()

	// decoding image
	img, _, e = image.Decode(imgFile)
	def_check(e)
	// pure rgbaing
	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)

	return rgba
}

func loadSource(path string) ([]byte, error) {
	sourceFile, e := asset.Open(path)
	defer sourceFile.Close()
	def_check(e)
	return ioutil.ReadAll(sourceFile)
}

func createProgram(vShaderPath, fShaderPath string) gl.Program {
	// loading vertex shader
	vShader, e := loadSource(vShaderPath)
	def_check(e)
	// loading fragment shader
	fShader, e := loadSource(fShaderPath)
	def_check(e)
	// linking program
	program, e := glutil.CreateProgram(string(vShader), string(fShader))
	crash_check(e)
	return program
}
