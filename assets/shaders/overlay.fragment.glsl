#version 330

// in vec2 texpos;

// uniform sampler2D tex;
// uniform vec4 color;

out vec4 fragColor;

void main(void) {
  // vec3 texel = texture(tex, texpos).rgb;
  // fragColor = vec4(texel.x,texel.y,texel.z, 1.0);
  // fragColor = max(texel.r,max(texel.g,texel.b));
  fragColor = vec(0.0, 0.0, 0.0, 1.0)
}