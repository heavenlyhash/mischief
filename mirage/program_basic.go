package mirage

import (
	"github.com/go-gl/gl/v3.2-core/gl"
	"github.com/go-gl/mathgl/mgl32"

	"exultant.us/mischief/render/glo"
	"exultant.us/mischief/render/prag"
)

type BasicProgram struct {
	prog               glo.Program
	uniform_projection uint32
	uniform_camera     uint32
	uniform_model      uint32
	uniform_tex        uint32
	attrib_vert        uint32
	attrib_textCoord   uint32

	vao, vbo uint32
}

var (
	_ prag.Program               = &BasicProgram{}
	_ prag.ProgramWithCamera     = &BasicProgram{}
	_ prag.ProgramWithProjection = &BasicProgram{}
)

func NewBasicProgram() prag.Program {
	prog := glo.CompileProgram(
		glo.CompileShader(glo.ShaderTVertex, shaderSource_vertex),
		glo.CompileShader(glo.ShaderTFragment, shaderSource_fragment),
	)
	return &BasicProgram{
		prog:               prog,
		uniform_projection: uint32(gl.GetUniformLocation(uint32(prog), gl.Str("projection\x00"))),
		uniform_camera:     uint32(gl.GetUniformLocation(uint32(prog), gl.Str("camera\x00"))),
		uniform_model:      uint32(gl.GetUniformLocation(uint32(prog), gl.Str("model\x00"))),
		uniform_tex:        uint32(gl.GetUniformLocation(uint32(prog), gl.Str("tex\x00"))),
		attrib_vert:        uint32(gl.GetAttribLocation(uint32(prog), gl.Str("vert\x00"))),
		attrib_textCoord:   uint32(gl.GetAttribLocation(uint32(prog), gl.Str("vertTexCoord\x00"))),
	}
}

var shaderSource_vertex = `
#version 330

uniform mat4 projection;
uniform mat4 camera;
uniform mat4 model;

in vec3 vert;
in vec2 vertTexCoord;

out vec2 fragTexCoord;

void main() {
    fragTexCoord = vertTexCoord;
    gl_Position = projection * camera * model * vec4(vert, 1);
}
` + "\x00"

var shaderSource_fragment = `
#version 330

uniform sampler2D tex;

in vec2 fragTexCoord;

out vec4 outputColor;

void main() {
    outputColor = texture(tex, fragTexCoord);
}
` + "\x00"

func (p *BasicProgram) SetModel(mtrx mgl32.Mat4) {
	gl.UniformMatrix4fv(int32(p.uniform_model), 1, false, &mtrx[0])
}

func (p *BasicProgram) SetLook(mtrx mgl32.Mat4) {
	gl.UniformMatrix4fv(int32(p.uniform_camera), 1, false, &mtrx[0])
}

func (p *BasicProgram) SetProjection(mtrx mgl32.Mat4) {
	gl.UniformMatrix4fv(int32(p.uniform_projection), 1, false, &mtrx[0])
}

func (p *BasicProgram) Preload() {
	// jk, didn't need this phase after all
}

func (p *BasicProgram) Arrange() {
	// We're in charge now.
	gl.UseProgram(uint32(p.prog))

	// Ask GL for handles to arrays.  Bind them.
	gl.GenVertexArrays(1, &p.vao)
	gl.GenBuffers(1, &p.vbo)

	// ?
	gl.Uniform1i(int32(p.uniform_tex), 0) // ?

	// Turn around and tell GL that we're going to use that array.
	gl.BindVertexArray(p.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, p.vbo)

	// Enable the 'attributes' to the GL program.
	// They're not enabled by default because GL loves to be redundant.
	// Also we need to be able to refer to these 'attributes' so that
	//  we can turn around and tell GL where to get their values from.
	gl.EnableVertexAttribArray(p.attrib_vert)
	gl.EnableVertexAttribArray(p.attrib_textCoord)

	// Tell GL how to interpret locations in our arrays onto meanings in the program.
	// - name the attribute to assign from slot
	// - specify how many there are (e.g. 3 for a vec3)
	// - specify the type
	// - "normalized", always false afaict
	// - "stride" length
	// - offset per stride
	gl.VertexAttribPointer(p.attrib_vert, 3, gl.FLOAT, false, 5*4, gl.PtrOffset(0))
	gl.VertexAttribPointer(p.attrib_textCoord, 2, gl.FLOAT, false, 5*4, gl.PtrOffset(3*4))
}

func (p *BasicProgram) Drop() {
	gl.DeleteBuffers(1, &p.vbo)
	gl.DeleteVertexArrays(1, &p.vao)
}
