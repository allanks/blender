package player

import (
)

const(
	LEFT,RIGHT,STATIONARY int32 = -1,1,0
)

type Player struct {
	Rotation int32
}

func CreatePlayer(startPos int32) (*Player) {
	p := new(Player)
	p.Rotation = startPos
	return p
}

func (user *Player) MoveX(direction int32) {
	user.Rotation = (user.Rotation + direction) % 360
}