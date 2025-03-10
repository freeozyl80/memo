package main

import "github.com/google/wire"

type Monster struct {
	Name string
}

func NewMonster() Monster {
	return Monster{Name: "kitty"}
}

type Player struct {
	Name string
}

func NewPlayer(name string) Player {
	return Player{Name: name}
}

type Mission struct {
	Player  Player
	Monster Monster
}

func NewMission(p Player, m Monster) Mission {
	return Mission{p, m}
}

func InitMission(name string) Mission {
	wire.Build(NewMonster, NewPlayer, NewMission)
	return Mission{}
}
