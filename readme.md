Gosteroids
==========

Asteroids clone built in Go, a tribute to [rusteroids](https://github.com/benbrunton/rusteroids)

## Current state

![Gosteroids](https://raw.github.com/morcmarc/gosteroids/master/gosteroids.gif)

## Installing

Run-time dependencies:

*Note: on Mac, make sure you install `libogg` and `libvorbis` first, otherwise
homebrew will compile `sdl` without Ogg support and you'll miss the tune :(*

- gl
- glfw3
- ogg
- vorbis
- sdl2
- sdl2_mixer

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

Less resource-heavy:

```
$ ./gosteroids -bg=0
```

### Controls

- Throttle (UP)
- Break (DOWN)
- Turn left (LEFT)
- Turn right (RIGHT)
- Fire (SPACE)

### Misc Controls

- Restart (R)
- Mute music (M)
- Volume down (-)
- Volume up (=)
- Quit (ESC)

### Usage

```
Usage of ./gosteroids:
  -bg=1: Background quality (0: low, 1: med, 2: high)
  -h=512: Height
  -w=512: Width
```