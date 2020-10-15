package systems

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

// Texture2D is able to store and configure a texture in OpenGL.
// It also hosts utility functions for easy management.
type Texture2D struct {
	ID uint32 // holds the ID of the texture object, used for all texture operations to reference to this particlar texture

	Height, Width int32 // texture image dimensions

	Internal_Format int32  // format of texture object
	Image_Format    uint32 // format of loaded image

	Wrap_S     int32 // wrapping mode on S axis
	Wrap_T     int32 // wrapping mode on T axis
	Filter_Min int32 // filtering mode if texture pixels < screen pixels
	Filter_Max int32 // filtering mode if texture pixels > screen pixels
}

func NewTexture2D() *Texture2D {
	myTex := Texture2D{Width: 0, Height: 0, Internal_Format: gl.RGB, Image_Format: gl.RGB, Wrap_S: gl.REPEAT, Wrap_T: gl.REPEAT, Filter_Min: gl.LINEAR, Filter_Max: gl.LINEAR}
	gl.GenTextures(1, &myTex.ID)
	return &myTex
}

// Generate generates texture from image data
func (tx *Texture2D) Generate(width, height int32, data interface{}) {
	tx.Width = width
	tx.Height = height

	gl.BindTexture(gl.TEXTURE_2D, tx.ID)
	gl.TexImage2D(gl.TEXTURE_2D, 0, tx.Internal_Format, tx.Width, tx.Height, 0, tx.Image_Format, gl.UNSIGNED_BYTE, gl.Ptr(data))

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, tx.Wrap_S)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, tx.Wrap_T)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, tx.Filter_Min)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, tx.Filter_Max)

	gl.BindTexture(gl.TEXTURE_2D, 0)
}

// Bind binds the texture as the current active GL_TEXTURE_2D texture object
func (tx *Texture2D) Bind() {
	gl.BindTexture(gl.TEXTURE_2D, tx.ID)
}
