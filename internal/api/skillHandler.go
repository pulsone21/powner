package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/pulsone21/powner/internal/entities"
)

type skillRequest struct {
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Type        entities.SkillType `json:"type"`
	Importance  int                `json:"importance"`
}

func getSkills(w http.ResponseWriter, r *http.Request) (any, *responseError) {
	mem, err := entities.GetSkills(db)
	if err != nil {
		return nil, newRespErr(500, err)
	}

	if len(*mem) == 0 {
		return nil, nil
	}

	return mem, nil
}

func createSkill(w http.ResponseWriter, r *http.Request) (any, *responseError) {
	var skillReq skillRequest
	err := json.NewDecoder(r.Body).Decode(&skillReq)
	if err != nil {
		return nil, newRespErr(400, err)
	}

	s := *entities.NewSkill(skillReq.Name, skillReq.Description, skillReq.Type, skillReq.Importance)

	log.Println(s)

	mem, err := entities.CreateSkill(db, s)
	if err != nil {
		return nil, newRespErr(500, err)
	}

	return mem, nil
}

func getSkillById(w http.ResponseWriter, r *http.Request) (any, *responseError) {
	strId := r.PathValue("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return nil, newRespErr(400, fmt.Errorf("id is not an uint: %v", strId))
	}

	mem, err := entities.GetSkillById(db, uint(id))
	if err != nil {
		return nil, newRespErr(500, err)
	}

	if mem == nil {
		return nil, nil
	}

	return mem, nil
}

func deleteSkill(w http.ResponseWriter, r *http.Request) (any, *responseError) {
	strId := r.PathValue("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return nil, newRespErr(400, fmt.Errorf("id is not an uint: %v", strId))
	}

	err = entities.DeleteSkill(db, uint(id))
	if err != nil {
		return nil, newRespErr(500, err)
	}
	return "Done", nil
}

func updateSkill(w http.ResponseWriter, r *http.Request) (any, *responseError) {
	var skillReq skillRequest
	err := json.NewDecoder(r.Body).Decode(&skillReq)
	if err != nil {
		return nil, newRespErr(400, err)
	}

	strId := r.PathValue("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return nil, newRespErr(400, fmt.Errorf("id is not an uint: %v", strId))
	}

	nT := &entities.Skill{
		Name:        skillReq.Name,
		Description: skillReq.Description,
		Importance:  skillReq.Importance,
		Type:        int(skillReq.Type),
	}

	nT.ID = uint(id)

	err = entities.UpdateSkill(db, *nT)
	if err != nil {
		return nil, newRespErr(500, err)
	}

	return nT, nil
}
