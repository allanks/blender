package player

import (
	"azul3d.org/gfx.v1"
	math "azul3d.org/lmath.v1"
	m "math"
	"fmt"
)

const (
	LENGTH,HEIGHT int = 10,10
	radius float64 = 2
)

var (
	CENTER math.Vec3 = math.Vec3{0,0,0}
	mesh = gfx.NewMesh()
	obj = gfx.NewObject()
)

func InitObject(shader *gfx.Shader) {
	mesh.Vertices = []gfx.Vec3{
		// Right Top Face
		{0, 4, 0},
		{0, 0, 0},
		{-1, -1, -1},
		
		// Left Top Face
		{0, 4, 0},
		{0, 0, 0},
		{1, -1, -1}}
	
	mesh.Colors = []gfx.Color{
		// Right Top Face
		gfx.Color{1, 0, 0, 1},
		gfx.Color{0, 1, 0, 1},
		gfx.Color{0, 0, 1, 1},
		
		// Left Top Face
		gfx.Color{1, 0, 0, 1},
		gfx.Color{0, 1, 0, 1},
		gfx.Color{0, 0, 1, 1}}
	
	obj.Shader = shader
	obj.OcclusionTest = true
	obj.State.FaceCulling = gfx.NoFaceCulling
	obj.Meshes = []*gfx.Mesh{mesh}
	fmt.Println(obj.WriteRed)
	fmt.Println(obj.WriteGreen)
	fmt.Println(obj.WriteBlue)
	fmt.Println(obj.WriteAlpha)
}

func DrawPlayer(user *Player, r gfx.Renderer, camera *gfx.Camera) {
	x := radius * m.Cos(float64(user.Rotation) * m.Pi / 180)
	y := radius * m.Sin(float64(user.Rotation) * m.Pi / 180)
	p := math.Vec3{x, 0, y}
	obj.SetPos(CENTER.Add(p.MulScalar(3)))
	r.Draw(r.Bounds(), obj, camera)
}

