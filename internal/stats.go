package internal

import "math/rand"

func calcAttack(trainerId int, monId int, level int) int {
	adjustedId := 10000 + trainerId
	withMonData := adjustedId * monId * level
	n := rand.NewSource(int64(withMonData)).Int63()

	return int(n % 16)
}

func calcDefense(trainerId int, monId int, level int) int {
	adjustedId := 10000 + trainerId
	withMonData := adjustedId * monId * level
	n := rand.NewSource(int64(withMonData)).Int63()

	return int((n >> 2) % 16)
}

func calcSpeed(trainerId int, monId int, level int) int {
	adjustedId := 10000 + trainerId
	withMonData := adjustedId * monId * level
	n := rand.NewSource(int64(withMonData)).Int63()

	return int((n >> 4) % 16)
}

func calcSpecial(trainerId int, monId int, level int) int {
	adjustedId := 10000 + trainerId
	withMonData := adjustedId * monId * level
	n := rand.NewSource(int64(withMonData)).Int63()

	return int((n >> 6) % 16)
}

func calcHealth(att int, def int, spe int, spc int) int {
	return (att%2)*8 + (def%2)*4 + (spe%2)*2 + (spc%2)*1
}
