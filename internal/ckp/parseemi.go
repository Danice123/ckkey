package ckp

import (
	"encoding/json"
	"os"
)

type EmiCalcData struct {
	Encounters []EncounterArea `json:"encounters"`
	Pools      EncounterPools  `json:"encounter_pools"`
	Pokemon    []Pokemon       `json:"pokemon"`

	PokeMap map[string]int
}

type EncounterArea struct {
	Area string `json:"area"`

	FishingTable   string `json:"fishing"`
	HeadbuttTable  string `json:"headbutt"`
	RockSmashTable string `json:"rock"`

	Walking *EncounterTable    `json:"normal"`
	Surfing []Encounter        `json:"surf"`
	Special []EncounterSpecial `json:"special"`
}

type EncounterPools struct {
	Fishing      []EncounterFishing `json:"fishing"`
	FishingMap   map[string]EncounterFishing
	Headbutt     []EncounterHeadbutt `json:"headbutt"`
	HeadbuttMap  map[string]EncounterHeadbutt
	RockSmash    []EncounterRockSmash `json:"rock"`
	RockSmashMap map[string]EncounterRockSmash
}

type EncounterFishing struct {
	Area  string         `json:"area"`
	Good  EncounterTable `json:"good"`
	Old   EncounterTable `json:"old"`
	Super EncounterTable `json:"super"`
}

type EncounterHeadbutt struct {
	Area  string      `json:"area"`
	Table []Encounter `json:"headbutt"`
}

type EncounterRockSmash struct {
	Area  string      `json:"area"`
	Table []Encounter `json:"rock"`
}

type EncounterSpecial struct {
	Pool []Encounter `json:"pool"`
	Type string      `json:"type"`
}

type EncounterTable struct {
	Day     []Encounter `json:"day"`
	Night   []Encounter `json:"night"`
	Morning []Encounter `json:"morning"`
}

type Encounter struct {
	Pokemon string `json:"pokemon"`
	Chance  int    `json:"chance"`
	Level   int    `json:"level"`
	Extra   string `json:"extra"`
}

type Pokemon struct {
	Name    string `json:"name"`
	Pokedex int    `json:"pokedex"`
}

func ParseEmiData(file string) (*EmiCalcData, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var data EmiCalcData
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, err
	}

	data.PokeMap = map[string]int{}
	for _, p := range data.Pokemon {
		data.PokeMap[p.Name] = p.Pokedex
	}

	data.Pools.FishingMap = map[string]EncounterFishing{}
	for _, e := range data.Pools.Fishing {
		data.Pools.FishingMap[e.Area] = e
	}

	data.Pools.HeadbuttMap = map[string]EncounterHeadbutt{}
	for _, e := range data.Pools.Headbutt {
		data.Pools.HeadbuttMap[e.Area] = e
	}

	data.Pools.RockSmashMap = map[string]EncounterRockSmash{}
	for _, e := range data.Pools.RockSmash {
		data.Pools.RockSmashMap[e.Area] = e
	}

	return &data, nil
}
