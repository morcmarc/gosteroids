#version 330

in vec4 position;

out vec2 texpos;

void main() {
  gl_Position = vec4(position.x, position.y, 0.0, 1.0);
  texpos = position.zw;
}