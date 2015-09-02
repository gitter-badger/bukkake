package main 

import (
	"golang.org/x/mobile/exp/f32"
	"encoding/binary"
)

var swastikaData = f32.Bytes(binary.LittleEndian,
	0.0, -0.5, 0.0,     0.0, 0.5, 0.0,
	-0.5, -0.5, 0.0,    0.0, -0.5, 0.0,
	0.0, 0.5, 0.0,      0.5, 0.5, 0.0,


	-0.5, 0.5, 0.0,     -0.5, 0.0, 0.0,
	-0.5, 0.0, 0.0,     0.5, 0.0, 0.0,
	0.5, 0.0, 0.0,      0.5, -0.5, 0.0,
	
)

var quadData = f32.Bytes(binary.LittleEndian,
	-0.3, -0.3, 0.0,
	0.3, -0.3, 0.0,
	0.3, 0.3, 0.0,

	0.3, 0.3, 0.0,
	-0.3, 0.3, 0.0,
	-0.3, -0.3, 0.0,
)

var quadTexData = f32.Bytes(binary.LittleEndian,
	0.0, 0.0, 	0.0, 1.0, 	1.0, 1.0,
	1.0, 1.0,	0.0, 1.0,	0.0, 0.0,
)



