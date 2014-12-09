#version 330

in vec2 texpos;

uniform sampler2D tex;
// uniform vec4 color;

out vec4 fragColor;

void main(void) {
  vec3 texel = texture(tex, texpos).rgb;
  fragColor = max(texel.r,max(texel.g,texel.b));
}