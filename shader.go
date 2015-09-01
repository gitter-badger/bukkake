package main 

import (
	"golang.org/x/mobile/asset"
)

func loadShaders(paths ...string) []string {
	// reading files up to 1KB
	buffer := make([]byte, 1024)
	// reading 2 shaders
	shaderSources := make([]string, 2)

	for index, path := range paths {
		file, e := asset.Open(path)
		def_check(e)

		n, e := file.Read(buffer)
		def_check(e)

		shaderSources[index] = string(buffer[:n])
	}
	return shaderSources
}