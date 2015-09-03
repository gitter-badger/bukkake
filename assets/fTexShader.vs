#version 100
precision mediump float;
uniform vec4 color;
uniform sampler2D myTexture;

varying vec2 v_texCoords;

void main() {
	vec4 tee = texture2D(myTexture, v_texCoords);
	//if (tee.x == 0.0f){
	//	color = vec4(1.0f, 1.0f, 1.0f, 1.0f);
	//}

	gl_FragColor = tee;
}