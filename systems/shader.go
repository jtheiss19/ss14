package systems

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

// Shader is a general purpsoe shader object. Compiles from file, generates
// compile/link-time error messages and hosts several utility
// functions for easy management.
type Shader struct {
	ID uint32
}

func NewShader() *Shader {
	myID := gl.CreateProgram()
	return &Shader{ID: myID}
}

// compiles the shader from given source code
func (sh *Shader) Compile(vertexShaderSource, fragmentShaderSource string) {
	var sVertex, sFragment uint32
	var err error

	sVertex, err = compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	sFragment, err = compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	gl.AttachShader(sh.ID, sVertex)
	gl.AttachShader(sh.ID, sFragment)
	gl.LinkProgram(sh.ID)
	gl.DeleteShader(sVertex)
	gl.DeleteShader(sFragment)
}

func (sh *Shader) Use() {
	gl.UseProgram(sh.ID)
}
func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
