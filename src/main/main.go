package main

import (
	"player"
	"level"
	"time"
	"image"
	"fmt"
	"azul3d.org/gfx.v1"
	"azul3d.org/gfx/window.v2"
	"azul3d.org/keyboard.v1"
	math "azul3d.org/lmath.v1"
)

var (
	user *player.Player
	lev *level.Level
	camFOV, camFar, camNear float64
)

var glslVert = []byte(`
#version 120

attribute vec3 Vertex;
attribute vec4 Color;

uniform mat4 MVP;

varying vec4 frontColor;

void main()
{
	frontColor = Color;
	gl_Position = MVP * vec4(Vertex, 1.0);
}
`)

var glslFrag = []byte(`
#version 120

varying vec4 frontColor;

void main()
{
	gl_FragColor = frontColor;
}
`)

func setup(){
	fmt.Println("Creating level")
	lev = level.CreateLevel()
	fmt.Println("Creating player")
	user = lev.CreatePlayer()
}

func eventHandler(w window.Window, r gfx.Renderer, camera *gfx.Camera) {
	// Create an event mask for the events we are interested in.
	evMask := window.FramebufferResizedEvents
	evMask |= window.KeyboardStateEvents

	// Create a channel of events.
	events := make(chan window.Event, 256)
	
	// Channels for animation
	xMove := make(chan int32)
	go moveUser(xMove)
	
	xMove <- player.STATIONARY

	// Have the window notify our channel whenever events occur.
	w.Notify(events, evMask)

	fmt.Println("Starting Event Loop")
	for e := range events {
		switch ev := e.(type) {
		case window.FramebufferResized:
			// Update the camera's projection matrix for the new width and
			// height.
			camera.Lock()
			camera.SetPersp(r.Bounds(), camFOV, camNear, camFar)
			camera.Unlock()

		case keyboard.StateEvent :
			switch ev.Key {
			case keyboard.A:
				if ev.State == keyboard.Down {
					xMove <- player.LEFT
				} else {
					xMove <- player.STATIONARY
				}
			case keyboard.D:
				if ev.State == keyboard.Down {
					xMove <- player.RIGHT
				} else {
					xMove <- player.STATIONARY
				}
			}
		}
	}
}

func levelUpdateLoop() {
	for {
		lev.Update()
		time.Sleep(16 * time.Millisecond)
	}
}

func gfxLoop(w window.Window, r gfx.Renderer) {
	fmt.Println("Starting Graphics Loop")
	// Setup a camera to use a perspective projection.
	camera := gfx.NewCamera()
	camFOV = 75.0
	camNear = 0.0001
	camFar = 1000.0
	camera.SetPersp(r.Bounds(), camFOV, camNear, camFar)

	// Move the camera -100 on the Y axis (back two units away from the triangle
	// object).
	camera.SetPos(math.Vec3{0, -10, 0})

	// Create a simple shader.
	shader := gfx.NewShader("SimpleShader")
	shader.GLSLVert = glslVert
	shader.GLSLFrag = glslFrag
	
	fmt.Println("Starting Event Handler")
	go eventHandler(w,r,camera)
	go levelUpdateLoop();
	
	level.InitObject(shader)
	player.InitObject(shader)
	
	fmt.Println("Starting Draw Loop")
	// Render loop
	tick := time.Tick(16 * time.Millisecond)
	for{
		r.Clear(image.Rect(0, 0, 0, 0), gfx.Color{1, 1, 1, 1})
		r.ClearDepth(image.Rect(0, 0, 0, 0), 1.0)
		level.DrawLevel(lev, r)
		player.DrawPlayer(user, r, camera)
		r.Render()
		<- tick
	}
}

func moveUser(xDir chan int32) {
	fmt.Println("Starting Move User Listener")
	xMove := int32(0)
	
	// Spawn listener for movement
	go func() {
		for {
			xMove =  <- xDir
		}
	}()
	
	for {
		//fmt.Println("Moving User")
		user.MoveX(xMove)
		time.Sleep(16 * time.Millisecond)
	}
}

func main() {
	setup()
	window.Run(gfxLoop, nil)
}

