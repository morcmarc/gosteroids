package shaders

const SpaceshipVertex string = `#version 330
in vec2 position;

void main()
{
    gl_Position = vec4(position, 0.0, 1.0);
}
`
const SpaceshipFragment string = `#version 330
out vec4 outColor;

void main()
{
    outColor = vec4(1.0, 1.0, 1.0, 1.0);
}
`
