package glo

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v3.2-core/gl"
)

/*
	Program is a ref to a collection of shaders which have been compiled
	onto the GPU.

	This is a fence around pointers, but our API is to only produce
	valid ones (RAII-like, to help keep partially instantiated objects out
	of sight).
*/
type Program uint32

/*
	ProgramParameter enum: list of parameters that can be set or read from a program.
*/
type ProgramParameter uint32

const (
	ProgramPCompileStatus ProgramParameter = gl.LINK_STATUS
	ProgramPInfoLogLength ProgramParameter = gl.INFO_LOG_LENGTH
)

func CompileProgram(shaders ...Shader) Program {
	program := Program(gl.CreateProgram())
	for _, shader := range shaders {
		gl.AttachShader(uint32(program), uint32(shader))
	}
	gl.LinkProgram(uint32(program))
	// Get compile status and errors back out from this nasty API.
	status := program.GetParameter(ProgramPCompileStatus)
	if status == gl.FALSE {
		log := program.GetCompileLog()
		panic(fmt.Errorf("failed to compile program: %v", log))
	}
	return program
}

func (program Program) GetParameter(param ProgramParameter) int {
	var p int32
	gl.GetProgramiv(uint32(program), uint32(param), &p)
	return int(p)
}

func (program Program) GetCompileLog() string {
	l := program.GetParameter(ProgramPInfoLogLength)
	if l <= 0 {
		return ""
	}
	buf := make([]byte, l)
	ptr := gl.Ptr(buf)
	gl.GetShaderInfoLog(uint32(program), int32(l), nil, (*uint8)(ptr))
	return strings.TrimRight(string(buf), "\x00")
}
