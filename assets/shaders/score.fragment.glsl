#version 330

in vec2 texpos;

uniform sampler2D tex;
uniform vec4 color;

out vec4  fragColor;

void main(void) {
  fragColor = texture(tex, texpos) * color;
}