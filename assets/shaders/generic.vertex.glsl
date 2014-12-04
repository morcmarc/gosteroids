#version 330

in vec2 vrtx;
uniform vec3 position;

void main()
{
    float x = vrtx[0];
    float y = vrtx[1];
    float x_pos = position[0];
    float y_pos = position[1];
    float angle = position[2];
    float xx = (x * cos(angle) + y * sin(angle)) + x_pos;
    float yy = (-x * sin(angle) + y * cos(angle)) + y_pos;
    gl_Position = vec4(xx, yy, 0.0, 1.0);
}