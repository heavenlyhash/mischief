package render

import "exultant.us/mischief/render/glo"

func NewProgram(vertexShaderSource, fragmentShaderSource string) glo.Program {
	return glo.CompileProgram(
		glo.CompileShader(glo.ShaderTVertex, vertexShaderSource),
		glo.CompileShader(glo.ShaderTFragment, fragmentShaderSource),
	)
}
