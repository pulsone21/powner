package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/pulsone21/powner/internal/entities"
)

type memberRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func getMembers(w http.ResponseWriter, r *http.Request) (any, *responseError) {
	mem, err := entities.GetMembers(db)
	if err != nil {
		return nil, newRespErr(500, err)
	}

	if len(*mem) == 0 {
		return nil, nil
	}

	return mem, nil
}

func createMember(w http.ResponseWriter, r *http.Request) (any, *responseError) {
	var memReq memberRequest
	err := json.NewDecoder(r.Body).Decode(&memReq)
	if err != nil {
		return nil, newRespErr(400, err)
	}

	log.Println(memReq)

	mem, err := entities.CreateMember(db, *entities.NewMember(memReq.Name, memReq.Age))
	if err != nil {
		return nil, newRespErr(500, err)
	}

	return mem, nil
}

func getMemberById(w http.ResponseWriter, r *http.Request) (any, *responseError) {
	strId := r.PathValue("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return nil, newRespErr(400, fmt.Errorf("id is not an uint: %v", strId))
	}
	mem, err := entities.GetMemberById(db, uint(id))
	if err != nil {
		return nil, newRespErr(500, err)
	}

	if mem == nil {
		return nil, nil
	}

	return mem, nil
}

func deleteMember(w http.ResponseWriter, r *http.Request) (any, *responseError) {
	strId := r.PathValue("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return nil, newRespErr(400, fmt.Errorf("id is not an uint: %v", strId))
	}
	err = entities.DeleteMember(db, uint(id))
	if err != nil {
		return nil, newRespErr(500, err)
	}
	return "Done", nil
}

func updateMember(w http.ResponseWriter, r *http.Request) (any, *responseError) {
	var memReq memberRequest
	err := json.NewDecoder(r.Body).Decode(&memReq)
	if err != nil {
		return nil, newRespErr(400, err)
	}

	strId := r.PathValue("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return nil, newRespErr(400, fmt.Errorf("id is not an uint: %v", strId))
	}

	newMem := &entities.Member{
		Name: memReq.Name,
		Age:  memReq.Age,
	}

	newMem.ID = uint(id)
	err = entities.UpdateMember(db, *newMem)
	if err != nil {
		return nil, newRespErr(500, err)
	}

	return newMem, nil
}

func try_find_member(id uint) (*entities.Member, *responseError) {
	mem, err := entities.GetMemberById(db, id)
	if err != nil {
		return nil, newRespErr(500, err)
	}

	if mem == nil {
		return nil, newRespErr(400, fmt.Errorf("no member found with id: %v", id))
	}

	return mem, nil
}
