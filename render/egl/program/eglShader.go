package eglProgram

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v3.2-core/gl"
)

/*
	Shader is a ref to a segment of code compiled onto the GPU.
	String them up into `Program`s, then tell the GPU to use that program
	when rendering.
*/
type Shader uint32

/*
	ShaderType enum: different types run at different stages in the pipeline.
	You must specify the type of shader when compiling it.
*/
type ShaderType uint32

const (
	ShaderTVertex   ShaderType = gl.VERTEX_SHADER
	ShaderTFragment ShaderType = gl.FRAGMENT_SHADER
	ShaderTGeometry ShaderType = gl.GEOMETRY_SHADER
)

func CompileShader(t ShaderType, src string) Shader {
	s := Shader(gl.CreateShader(uint32(t)))
	src_cstrs, src_cstrsFreeFn := gl.Strs(src + "\x00")
	// If for some reason you were doing a ton of different shaders,
	//  you could see this as mutable.  But we're not exposing that.  why.
	gl.ShaderSource(uint32(s), 1, src_cstrs, nil)
	src_cstrsFreeFn()
	gl.CompileShader(uint32(s))
	// Get compile status and errors back out from this nasty API.
	status := s.getParameter(gl.COMPILE_STATUS)
	if status == gl.FALSE {
		log := s.getCompileLog()
		panic(fmt.Errorf("failed to compile shader: %v", log))
	}
	return s
}

// Helper method for reading int params from shader.
// Used in other what-when-wrong detection.
func (s Shader) getParameter(param uint32) int {
	var p int32
	gl.GetShaderiv(uint32(s), param, &p)
	return int(p)
}

func (s Shader) getCompileLog() string {
	l := s.getParameter(gl.INFO_LOG_LENGTH)
	if l <= 0 {
		return ""
	}
	buf := make([]byte, l)
	ptr := gl.Ptr(buf)
	gl.GetShaderInfoLog(uint32(s), int32(l), nil, (*uint8)(ptr))
	return strings.TrimRight(string(buf), "\x00")
}
