package service

import (
	"game/pkg/life"
	"math/rand"
	"time"
)

// для хранения состояния
type LifeService struct {
	CurrentWorld *life.World
	NextWorld    *life.World
}

func New(height, width int) (*LifeService, error) {
	rand.NewSource(time.Now().UTC().UnixNano())

	currentWorld, err := life.NewWorld(height, width)
	if err != nil {
		return nil, err
	}
	// для упрощения примера хаотично заполним
	currentWorld.RandInit(7)

	newWorld, err := life.NewWorld(height, width)
	if err != nil {
		return nil, err
	}

	ls := LifeService{
		CurrentWorld: currentWorld,
		NextWorld:    newWorld,
	}

	return &ls, nil
}

// получение очередного состояния игры
func (ls *LifeService) NewState() *life.World {
	life.NextState(ls.CurrentWorld, ls.NextWorld)

	ls.CurrentWorld = ls.NextWorld

	return ls.CurrentWorld
}
