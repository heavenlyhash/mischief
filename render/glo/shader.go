package glo

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v3.2-core/gl"
)

/*
	Shader is a ref to a segment of code compiled onto the GPU.
	String them up into `Program`s, then tell the GPU to use that program
	when rendering.

	This is a fence around pointers, but our API is to only produce
	valid ones (RAII-like, to help keep partially instantiated objects out
	of sight).
*/
type Shader uint32

/*
	ShaderType enum: different types run	at different stages in the pipeline.
*/
type ShaderType uint32

const (
	ShaderTVertex   ShaderType = gl.VERTEX_SHADER
	ShaderTFragment ShaderType = gl.FRAGMENT_SHADER
	ShaderTGeometry ShaderType = gl.GEOMETRY_SHADER
)

/*
	ShaderParameter enum: list of parameters that can be set or read from a shader.
*/
type ShaderParameter uint32

const (
	ShaderPCompileStatus ShaderParameter = gl.COMPILE_STATUS
	ShaderPInfoLogLength ShaderParameter = gl.INFO_LOG_LENGTH
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
	status := s.GetParameter(ShaderPCompileStatus)
	if status == gl.FALSE {
		log := s.GetCompileLog()
		panic(fmt.Errorf("failed to compile shader: %v", log))
	}
	return s
}

func (s Shader) GetParameter(param ShaderParameter) int {
	var p int32
	gl.GetShaderiv(uint32(s), uint32(param), &p)
	return int(p)
}

func (s Shader) GetCompileLog() string {
	l := s.GetParameter(ShaderPInfoLogLength)
	if l <= 0 {
		return ""
	}
	buf := make([]byte, l)
	ptr := gl.Ptr(buf)
	gl.GetShaderInfoLog(uint32(s), int32(l), nil, (*uint8)(ptr))
	return strings.TrimRight(string(buf), "\x00")
}
