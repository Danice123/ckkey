package internal

import (
	"errors"
	"hash/fnv"
	"math/rand"
	"strconv"
)

var walkingValues = map[string]int{
	"a": 30,
	"b": 30,
	"c": 20,
	"d": 10,
	"e": 5,
	"f": 4,
	"g": 1,
}

var surfingValues = map[string]int{
	"a": 60,
	"b": 30,
	"c": 10,
}

var headbuttValues = map[string]int{
	"a": 50,
	"b": 15,
	"c": 15,
	"d": 10,
	"e": 5,
	"f": 5,
}

var oldRodValues = map[string]int{
	"a": 75,
	"b": 15,
	"c": 15,
}

var goodRodValues = map[string]int{
	"a": 35,
	"b": 35,
	"c": 20,
	"d": 10,
}

var superRodValues = map[string]int{
	"a": 40,
	"b": 30,
	"c": 20,
	"d": 10,
}

var rockSmashValues = map[string]int{
	"a": 90,
	"b": 10,
}

var oddEggValues = map[string]int{
	"a": 19,
	"b": 19,
	"c": 16,
	"d": 14,
	"e": 12,
	"f": 11,
	"g": 9,
}

func GenerateEncounterOrder(trainerId int, area string, eType string, time string) ([]string, error) {
	switch eType {
	case "walking":
		r := rand.New(rand.NewSource(hashString(strconv.Itoa(trainerId) + area + "_walking_" + time)))
		return calcBuckets(r, walkingValues), nil
	case "surfing":
		r := rand.New(rand.NewSource(hashString(strconv.Itoa(trainerId) + area + "_surfing")))
		return calcBuckets(r, surfingValues), nil
	case "headbutt":
		r := rand.New(rand.NewSource(hashString(strconv.Itoa(trainerId) + area + "_headbutt")))
		return calcBuckets(r, headbuttValues), nil
	case "old rod":
		r := rand.New(rand.NewSource(hashString(strconv.Itoa(trainerId) + area + "_oldrod")))
		return calcBuckets(r, oldRodValues), nil
	case "good rod":
		if time == "morning" {
			time = "day"
		}
		r := rand.New(rand.NewSource(hashString(strconv.Itoa(trainerId) + area + "_goodrod_" + time)))
		return calcBuckets(r, goodRodValues), nil
	case "super rod":
		if time == "morning" {
			time = "day"
		}
		r := rand.New(rand.NewSource(hashString(strconv.Itoa(trainerId) + area + "_superrod_" + time)))
		return calcBuckets(r, superRodValues), nil
	case "rock smash":
		r := rand.New(rand.NewSource(hashString(strconv.Itoa(trainerId) + area + "_rocksmash")))
		return calcBuckets(r, rockSmashValues), nil
	case "odd egg":
		r := rand.New(rand.NewSource(hashString(strconv.Itoa(trainerId))))
		return calcBuckets(r, oddEggValues), nil
	default:
		return nil, errors.New("bad encounter type name")
	}
}

func calcBuckets(r *rand.Rand, valueMap map[string]int) []string {
	var buckets []string
	for k := range valueMap {
		buckets = append(buckets, k)
	}

	var order []string
Loop:
	for len(buckets) > 0 {
		var total int
		for i := 0; i < len(buckets); i++ {
			total += valueMap[buckets[i]]
		}
		n := r.Intn(total)
		for i := 0; i < len(buckets); i++ {
			n -= valueMap[buckets[i]]
			if n < 0 {
				order = append(order, buckets[i])
				buckets = removeFromSlice(buckets, i)
				continue Loop
			}
		}
	}
	return order
}

func removeFromSlice(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func hashString(s string) int64 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int64(h.Sum32())
}
