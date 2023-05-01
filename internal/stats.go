package internal

import (
	"errors"
	"fmt"
	"math/rand"
)

type Encounter struct {
	Attack  int
	Defense int
	Speed   int
	Special int
}

func CalcEncounter(trainerId int, dexId int, level int) (Encounter, error) {
	if trainerId < 0 || trainerId > 65535 {
		return Encounter{}, errors.New("trainer ID invalid")
	}
	if dexId < 1 || dexId > 251 {
		return Encounter{}, errors.New("dex ID invalid")
	}
	if level < 1 {
		return Encounter{}, errors.New("level invalid")
	}
	adjustedId := 10000 + int64(trainerId)
	withMonData := adjustedId * int64(dexId) * int64(level)
	r := rand.NewSource(withMonData)

	return Encounter{
		Attack:  int(r.Int63() % 16),
		Defense: int(r.Int63() % 16),
		Speed:   int(r.Int63() % 16),
		Special: int(r.Int63() % 16),
	}, nil

}

func (e Encounter) CalcHealth() int {
	return (e.Attack%2)*8 + (e.Defense%2)*4 + (e.Speed%2)*2 + (e.Special%2)*1
}

func (e Encounter) Print() {
	fmt.Printf("Health: %d\nAttack: %d\nDefense: %d\nSpeed: %d\nSpecial: %d\n",
		e.CalcHealth(),
		e.Attack,
		e.Defense,
		e.Speed,
		e.Special)
}
