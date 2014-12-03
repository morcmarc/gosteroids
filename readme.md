Gosteroids
==========

Asteroids clone built in Go, a tribute to [rusteroids](https://github.com/benbrunton/rusteroids)

## Installing

Run-time dependencies:

- gl
- glew
- glfw3

All compile-time dependencies can be installed via [Godep](https://github.com/tools/godep):

```
$ godep restore
```

## Compiling

To compile run:

```
$ make install
```

This will create the executable binary in the root folder.

## Running

Make sure you're in the project root folder as the binary have to load in
the shader files.

```
$ ./gosteroids
```