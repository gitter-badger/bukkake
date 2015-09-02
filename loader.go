package main 

import (
	"golang.org/x/mobile/asset"
)

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