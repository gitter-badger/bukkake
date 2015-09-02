#version 100
uniform mat3 rotationMatrix;
uniform float resIndex;
uniform vec3 translation;

attribute vec4 position;
void main() {
	vec3 pos = rotationMatrix * position.xyz + translation;
	gl_Position = vec4(pos.x, pos.y*resIndex, pos.z, 1.0);
}