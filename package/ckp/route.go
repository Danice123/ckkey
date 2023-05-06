package ckp

import (
	"hash/fnv"
	"math/rand"
	"strconv"
)

func GenerateEncounterOrderFromEmi(trainerId int, encounters []Encounter, areaSeed string) []Encounter {
	r := rand.New(rand.NewSource(hashString(strconv.Itoa(trainerId) + areaSeed)))
	buckets := append([]Encounter{}, encounters...)

	var order []Encounter
Loop:
	for len(buckets) > 0 {
		var total int
		for i := 0; i < len(buckets); i++ {
			total += buckets[i].Chance
		}
		n := r.Intn(total)
		for i := 0; i < len(buckets); i++ {
			n -= buckets[i].Chance
			if n < 0 {
				order = append(order, buckets[i])
				buckets = removeFromSlice(buckets, i)
				continue Loop
			}
		}
	}
	return order
}

func removeFromSlice(slice []Encounter, s int) []Encounter {
	return append(slice[:s], slice[s+1:]...)
}

func hashString(s string) int64 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int64(h.Sum32())
}
