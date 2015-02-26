package level

import (
	"azul3d.org/gfx.v1"
	math "azul3d.org/lmath.v1"
	m "math"
)

const (
	radius float64 = 2
)

var (
	mesh = gfx.NewMesh()
	obj = gfx.NewObject()
)

func InitObject(shader *gfx.Shader) {
	mesh.Vertices = []gfx.Vec3{
		{0, 0, 0},
		{float32(xLength), 0, 0},
		{0, 0, float32(zLength)},
		
		{float32(xLength), 0, float32(zLength)},
		{float32(xLength), 0, 0},
		{0, 0, float32(zLength)}}
	
	mesh.Colors = []gfx.Color{
		{1, 0, 0, 1},
		{0, 1, 0, 1},
		{0, 0, 1, 1},
		
		{1, 0, 0, 1},
		{0, 1, 0, 1},
		{0, 0, 1, 1}}
	
	obj.Shader = shader
	obj.OcclusionTest = true
	obj.State.FaceCulling = gfx.NoFaceCulling
	obj.Meshes = []*gfx.Mesh{mesh}
}

func DrawLevel(l *Level, r gfx.Renderer) {
	for _, p := range l.platforms {
		x := radius * m.Cos(float64(p.rotation) * m.Pi / 180)
		y := radius * m.Sin(float64(p.rotation) * m.Pi / 180)
		p := math.Vec3{x, y, 0}
		obj.SetPos(obj.Pos().Add(p.MulScalar(3)))
	}
}

