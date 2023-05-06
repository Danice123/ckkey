package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Danice123/ckkey/package/roll"
	"github.com/julienschmidt/httprouter"
)

type RollDVResponse struct {
	DVs      roll.DVs
	HealthDV int
	HexCode  string
}

func RollDVs(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	trainerId, err := strconv.Atoi(ps.ByName("trainerId"))
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	if trainerId < 0 || trainerId > 65535 {
		writeError(w, http.StatusBadRequest, errors.New("trainer ID invalid"))
		return
	}
	dexId, err := strconv.Atoi(ps.ByName("dexId"))
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	if dexId < 1 || dexId > 251 {
		writeError(w, http.StatusBadRequest, errors.New("dex ID invalid"))
		return
	}
	query := req.URL.Query()
	level, err := strconv.Atoi(query.Get("level"))
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	if level < 1 {
		writeError(w, http.StatusBadRequest, errors.New("level invalid"))
		return
	}

	enc := roll.CalcDVs(trainerId, dexId, level)
	resp := RollDVResponse{
		DVs:      enc,
		HealthDV: enc.CalcHealth(),
		HexCode:  fmt.Sprintf("%x%x%x%x", enc.Attack, enc.Defense, enc.Speed, enc.Special),
	}

	b, err := json.Marshal(resp)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}
	w.Write(b)
}

func writeError(w http.ResponseWriter, code int, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(err.Error()))
}
