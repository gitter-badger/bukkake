package main

import (
	"bytes"
	"golang.org/x/mobile/asset"

	"image"
	"image/draw"
	"image/png"
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

func loadSources(paths ...string) [][]byte {
	// reading files up to 1KB
	buffer := make([]byte, 1024)
	// reading 4 files at the time
	sources := make([][]byte, 2)

	for index, path := range paths {
		file, e := asset.Open(path)
		def_check(e)

		n, e := file.Read(buffer)
		def_check(e)

		sources[index] = buffer[:n]
	}
	return sources
}

func loadShaders(paths ...string) []string {
	// reading files up to 1KB
	buffer := make([]byte, 1024)
	// reading 2 shaders
	sources := make([]string, 2)

	for index, path := range paths {
		file, e := asset.Open(path)
		def_check(e)

		n, e := file.Read(buffer)
		def_check(e)

		sources[index] = string(buffer[:n])
	}
	return sources
}

func testImage() []byte {
	var img image.Image
	buffer := new(bytes.Buffer)

	imgFile, e := asset.Open("pee.png")
	def_check(e)

	img, _, e = image.Decode(imgFile)
	e = png.Encode(buffer, img)
	return buffer.Bytes()

	//return img
}

func onePixel() []byte {
	pixel := []byte{255, 255, 255, 255}
	return pixel
}
