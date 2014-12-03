#version 330

in vec2 vrtx;

void main()
{
    float x = vrtx[0];
    float y = vrtx[1];
    gl_Position = vec4(x, y, 0.9, 1.0);
}