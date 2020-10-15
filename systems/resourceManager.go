package systems

import (
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-gl/gl/v4.1-core/gl"
)

// ResourceManager is a class that hosts several
// functions to load Textures and Shaders. Each loaded texture
// and/or shader is also stored for future reference by string
// handles. All functions and resources are static and no
// public constructor is defined.
type ReasourceManager struct {
	Shaders  map[string]*Shader
	Textures map[string]*Texture2D
}

func (rm *ReasourceManager) LoadShader(VertexShaderFile, FragmentShaderFile, name string) *Shader {
	rm.Shaders[name] = loadShaderFromFile(VertexShaderFile, FragmentShaderFile)
	return rm.Shaders[name]
}

func (rm *ReasourceManager) GetShader(name string) *Shader {
	return rm.Shaders[name]
}

func (rm *ReasourceManager) LoadTexture(filepath, name string, alpha bool) *Texture2D {
	rm.Textures[name] = loadTextureFromFile(filepath, alpha)
	return rm.Textures[name]
}

func (rm *ReasourceManager) GetTexture(name string) *Texture2D {
	return rm.Textures[name]
}

func (rm *ReasourceManager) Clear() {
	for _, shader := range rm.Shaders {
		gl.DeleteProgram(shader.ID)
	}
	for _, texture := range rm.Shaders {
		gl.DeleteTextures(1, &texture.ID)
	}
}

//Currently only supports PNG's
func loadTextureFromFile(filepath string, alpha bool) *Texture2D {
	myTexture := NewTexture2D()
	if alpha {
		myTexture.Internal_Format = gl.RGBA
		myTexture.Image_Format = gl.RGBA
	}

	infile, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer infile.Close()

	img, err := png.Decode(infile)
	if err != nil {
		log.Fatal(err.Error())
	}

	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y

	pixels := make([]byte, width*height*4)
	bindex := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			pixels[bindex] = byte(r / 256)
			bindex++
			pixels[bindex] = byte(g / 256)
			bindex++
			pixels[bindex] = byte(b / 256)
			bindex++
			pixels[bindex] = byte(a / 256)
			bindex++
		}
	}

	myTexture.Generate(int32(width), int32(height), pixels)

	return myTexture
}

func loadShaderFromFile(VertexShaderFile, FragmentShaderFile string) *Shader {
	var vertexCode, fragmentCode string
	vertexFile, err := ioutil.ReadFile(VertexShaderFile)
	if err != nil {
		log.Fatal(err.Error())
	}
	fragmentFile, err := ioutil.ReadFile(FragmentShaderFile)
	if err != nil {
		log.Fatal(err.Error())
	}
	vertexCode = string(vertexFile)
	fragmentCode = string(fragmentFile)

	myShader := NewShader()
	myShader.Compile(vertexCode, fragmentCode)

	return myShader

}
