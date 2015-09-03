#version 100
uniform mat3 rotationMatrix;
uniform float resIndex;
uniform vec3 translation;

attribute vec4 position;
attribute vec2 texCoords;

varying vec2 v_texCoords;
void main() {
	v_texCoords = texCoords;
	vec3 pos = rotationMatrix * position.xyz + translation;
	gl_Position = vec4(pos.x, pos.y*resIndex, pos.z, 1.0);
}