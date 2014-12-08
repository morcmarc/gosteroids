package graphics

import (
	// "fmt"
	"image"
	"image/png"
	"os"
	// "reflect"

	"github.com/go-gl/gl"
)

type Font struct {
	TextureSource *image.NRGBA
	Texture       gl.Texture
	Width         int
	Height        int
	Columns       int
	Rows          int
	GlyphW        float32
	GlyphH        float32
	vbo           gl.Buffer
	vao           gl.VertexArray
	program       gl.Program
	color         []float32
	vertices      []float32
}

func NewFont(filename string, cols, rows, width, height int) *Font {
	f := &Font{
		Width:    width,
		Height:   height,
		Columns:  cols,
		Rows:     rows,
		color:    []float32{1, 1, 1, 1},
		vertices: []float32{},
	}

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	img, err := png.Decode(file)
	if err != nil {
		panic(err)
	}

	rgbaImg, ok := img.(*image.NRGBA)
	if !ok {
		panic("Invalid image type")
	}

	imgWidth, imgHeight := img.Bounds().Dx(), img.Bounds().Dy()

	f.GlyphW = ((float32(imgWidth) / float32(cols)) / float32(width)) * 2
	f.GlyphH = ((float32(imgHeight) / float32(rows)) / float32(height)) * 2
	f.TextureSource = rgbaImg

	vs, _ := LoadShader("assets/shaders/textured.vertex.glsl", VertexShader)
	fs, _ := LoadShader("assets/shaders/textured.fragment.glsl", FragmentShader)

	f.program = NewProgram(vs, fs)

	f.vao = gl.GenVertexArray()
	f.vao.Bind()

	f.vbo = gl.GenBuffer()
	f.vbo.Bind(gl.ARRAY_BUFFER)
	f.vbo.Unbind(gl.ARRAY_BUFFER)

	textureUniform := f.program.GetUniformLocation("tex")

	// tex, _ := CreateTexture(file)
	gl.ActiveTexture(gl.TEXTURE0)
	tex := gl.GenTexture()
	tex.Bind(gl.TEXTURE_2D)
	textureUniform.Uniform1i(0)

	/* We require 1 byte alignment when uploading texture data */
	gl.PixelStorei(gl.UNPACK_ALIGNMENT, 1)
	/* Clamping to edges is important to prevent artifacts when scaling */
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.CLAMP_TO_EDGE)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.CLAMP_TO_EDGE)
	/* Linear filtering usually looks best for text */
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexImage2D(
		gl.TEXTURE_2D,
		0, gl.RGBA,
		rgbaImg.Bounds().Dx(),
		rgbaImg.Bounds().Dy(),
		0,
		gl.RGBA,
		gl.UNSIGNED_BYTE,
		rgbaImg.Pix)

	f.vao.Unbind()

	f.Texture = tex

	return f
}

func (f *Font) Printf(s string, x, y float32) {
	var sw float32 = 1.0 / float32(f.Columns)
	var th float32 = 1.0 / float32(f.Rows)

	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	for i, c := range s {
		f.program.Use()
		defer f.program.Unuse()

		colorUniform := f.program.GetUniformLocation("color")
		colorUniform.Uniform4fv(1, f.color)

		f.vbo.Bind(gl.ARRAY_BUFFER)
		f.vao.Bind()

		column := int(c) % f.Columns
		row := (int(c) / f.Columns)
		u := float32(column) * sw
		v := float32(f.Rows-row+1) * th
		xx := x + float32(i)*f.GlyphW
		yy := y

		f.vertices = []float32{
			xx, yy - f.GlyphH, u, v,
			xx, yy, u, v + th,
			xx + f.GlyphW, yy - f.GlyphH, u + sw, v,

			xx + f.GlyphW, yy - f.GlyphH, u + sw, v,
			xx + f.GlyphW, yy, u + sw, v + th,
			xx, yy, u, v + th,
		}

		// Invert V because we're using a compressed texture
		for i := 3; i < len(f.vertices); i += 4 {
			f.vertices[i] = 1.0 - f.vertices[i]
		}

		positionAttrib := f.program.GetAttribLocation("position")
		positionAttrib.AttribPointer(4, gl.FLOAT, false, 0, nil)
		positionAttrib.EnableArray()

		gl.BufferData(gl.ARRAY_BUFFER, len(f.vertices)*4, f.vertices, gl.STATIC_DRAW)
		gl.DrawArrays(gl.TRIANGLES, 0, len(f.vertices))

		f.vao.Unbind()
		f.vbo.Unbind(gl.ARRAY_BUFFER)
	}

	f.program.Unuse()
	gl.Disable(gl.BLEND)
}

func (f *Font) Delete() {
	f.vao.Delete()
	f.vbo.Delete()
	f.Texture.Delete()
	f.program.Delete()
}
