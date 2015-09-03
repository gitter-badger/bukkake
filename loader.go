package main

import (
	"bytes"
	"golang.org/x/mobile/asset"
	"image"
	"image/png"
)

func loadImages(paths ...string) [][]byte {
	// loading up to 4 files
	sources := make([][]byte, 4)
	// allocating image holder
	var img image.Image
	// creating byte buffer
	buffer := new(bytes.Buffer)

	for index, path := range paths {
		// opening file from /assets/
		imgFile, e := asset.Open(path)
		def_check(e)
		defer imgFile.Close()

		// decoding image
		img, _, e = image.Decode(imgFile)
		def_check(e)
		// encoding to buffer
		e = png.Encode(buffer, img)
		// injecting []byte image
		sources[index] = buffer.Bytes()
	}
	return sources
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
