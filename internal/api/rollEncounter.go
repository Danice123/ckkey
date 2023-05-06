package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/Danice123/ckkey/package/ckp"
	"github.com/Danice123/ckkey/package/roll"
	"github.com/julienschmidt/httprouter"
)

var ckpData *ckp.EmiCalcData

func init() {
	var err error
	ckpData, err = ckp.ParseEmiData("ckp_data.json")
	if err != nil {
		panic(err)
	}
}

type RollEncounterResponse struct {
	Area string

	Walking   *TimeEncounters
	Surfing   []Encounter
	Fishing   map[string]TimeEncounters
	Headbutt  []Encounter
	RockSmash []Encounter
	Special   []SpecialEncounter
}

type FishingEncounters struct {
	Old   []Encounter
	Good  TimeEncounters
	Super TimeEncounters
}

type SpecialEncounter struct {
	Type string
	Pool []Encounter
}

type TimeEncounters struct {
	Day     []Encounter
	Night   []Encounter
	Morning []Encounter
}

type Encounter struct {
	Pokemon  string
	Level    int
	HealthDV int
	DVs      roll.DVs
}

func RollEncounter(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	trainerId, err := strconv.Atoi(ps.ByName("trainerId"))
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	if trainerId < 0 || trainerId > 65535 {
		writeError(w, http.StatusBadRequest, errors.New("trainer id invalid"))
		return
	}
	areaName := strings.ReplaceAll(strings.ToLower(ps.ByName("area")), " ", "")

	resp := BuildEncounterResponse(trainerId, areaName)
	if resp == nil {
		writeError(w, http.StatusBadRequest, errors.New("area invalid"))
		return
	}

	b, err := json.Marshal(resp)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	w.Write(b)
}

func BuildEncounterResponse(trainerId int, areaName string) *RollEncounterResponse {
	var area *ckp.EncounterArea
	for _, a := range ckpData.Encounters {
		if strings.ReplaceAll(strings.ToLower(a.Area), "-", "") == areaName {
			area = &a
			break
		}
	}
	if area == nil {
		return nil
	}

	var resp RollEncounterResponse
	resp.Area = area.Area
	if area.Walking != nil {
		ts := rollTimedStates(trainerId, area.Walking, areaName+"_walking")
		resp.Walking = &ts
	}

	if area.Surfing != nil {
		resp.Surfing = rollStats(
			trainerId,
			ckp.GenerateEncounterOrderFromEmi(trainerId, area.Surfing, areaName+"_surfing"),
		)
	}

	if area.FishingTable != "" {
		table := ckpData.Pools.FishingMap[area.FishingTable]
		resp.Fishing = map[string]TimeEncounters{}
		resp.Fishing["Old"] = rollTimedStates(trainerId, &table.Old, areaName+"_oldrod")
		resp.Fishing["Good"] = rollTimedStates(trainerId, &table.Good, areaName+"_goodrod")
		resp.Fishing["Super"] = rollTimedStates(trainerId, &table.Super, areaName+"_superrod")
	}

	if area.HeadbuttTable != "" {
		table := ckpData.Pools.HeadbuttMap[area.HeadbuttTable]
		resp.Headbutt = rollStats(
			trainerId,
			ckp.GenerateEncounterOrderFromEmi(trainerId, table.Table, areaName+"_headbutt"),
		)
	}

	if area.RockSmashTable != "" {
		table := ckpData.Pools.RockSmashMap[area.RockSmashTable]
		resp.RockSmash = rollStats(
			trainerId,
			ckp.GenerateEncounterOrderFromEmi(trainerId, table.Table, areaName+"_rocksmash"),
		)
	}

	if area.Special != nil {
		resp.Special = []SpecialEncounter{}
		for _, special := range area.Special {
			if special.Pool != nil {
				resp.Special = append(resp.Special, SpecialEncounter{
					Type: special.Type,
					Pool: rollStats(
						trainerId,
						ckp.GenerateEncounterOrderFromEmi(trainerId, special.Pool, areaName+"_"+special.Type),
					),
				})
			}
		}
	}

	return &resp
}

func rollTimedStates(trainerId int, table *ckp.EncounterTable, seed string) TimeEncounters {
	var enc TimeEncounters
	enc.Day = rollStats(
		trainerId,
		ckp.GenerateEncounterOrderFromEmi(trainerId, table.Day, seed+"_day"),
	)
	if len(table.Night) > 0 && !reflect.DeepEqual(table.Day, table.Night) {
		enc.Night = rollStats(
			trainerId,
			ckp.GenerateEncounterOrderFromEmi(trainerId, table.Night, seed+"_night"),
		)
	}
	if len(table.Morning) > 0 && !reflect.DeepEqual(table.Day, table.Morning) {
		enc.Morning = rollStats(
			trainerId,
			ckp.GenerateEncounterOrderFromEmi(trainerId, table.Morning, seed+"_morning"),
		)
	}
	return enc
}

func rollStats(trainerId int, table []ckp.Encounter) []Encounter {
	enc := []Encounter{}
	for _, e := range table {
		dvs := roll.CalcDVs(trainerId, ckpData.PokeMap[e.Pokemon], e.Level)
		enc = append(enc, Encounter{
			Pokemon:  e.Pokemon,
			Level:    e.Level,
			HealthDV: dvs.CalcHealth(),
			DVs:      dvs,
		})
	}
	return enc
}
