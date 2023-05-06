package roll

import (
	"fmt"
	"math/rand"
)

type DVs struct {
	Attack  int
	Defense int
	Speed   int
	Special int
}

func CalcDVs(trainerId int, dexId int, level int) DVs {
	adjustedId := 10000 + int64(trainerId)
	withMonData := adjustedId * int64(dexId) * int64(level)
	r := rand.NewSource(withMonData)

	return DVs{
		Attack:  int(r.Int63() % 16),
		Defense: int(r.Int63() % 16),
		Speed:   int(r.Int63() % 16),
		Special: int(r.Int63() % 16),
	}

}

func (e DVs) CalcHealth() int {
	return (e.Attack%2)*8 + (e.Defense%2)*4 + (e.Speed%2)*2 + (e.Special%2)*1
}

func (e DVs) Print() {
	fmt.Printf("Health: %d\nAttack: %d\nDefense: %d\nSpeed: %d\nSpecial: %d\n",
		e.CalcHealth(),
		e.Attack,
		e.Defense,
		e.Speed,
		e.Special)
}

func (e DVs) PrintHex() {
	fmt.Printf("%x%x%x%x\n", e.Attack, e.Defense, e.Speed, e.Special)
}
