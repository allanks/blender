package level

import (
	"player"
)

const(
	xLength, zLength int32 = 10,100
)

type Platform struct {
	rotation, zPos int32
}

type Level struct {
	platforms []Platform
	playerStartXPos int32
}

func CreateLevel() (*Level) {
	
	l := new(Level)
	
	
	
	return l
}

func (l *Level) CreatePlayer() (*player.Player) {
	p := player.CreatePlayer(l.playerStartXPos)
	return p
}

func (p *Platform) isPlayerOnPlatform(user *player.Player) bool {
	if p.isCurrent() && p.isPlayerOnPlatform(user){
		return true
	}
	
	return false
}

func (p *Platform) isCurrent() bool {
	return p.zPos < zLength && p.zPos > 0
}

func (p *Platform) userCollision(user *player.Player) bool {
	return user.Rotation < (p.rotation + xLength) && user.Rotation > p.rotation
}

func (l *Level) Update() {
	for _, p := range l.platforms {
		p.zPos = p.zPos - 1
	}
}

func (l *Level) IsPlayerOnPlatform(user *player.Player) bool {
	for _,p := range l.platforms {
		if p.isPlayerOnPlatform(user) {
			return true
		}
	}
	return false
}

