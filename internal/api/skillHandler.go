package api

import (
	"encoding/json"
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

func getSkills(w http.ResponseWriter, r *http.Request) *response {
	mem, err := entities.GetSkills(db)
	if err != nil {
		return internalError(err)
	}

	if len(*mem) == 0 {
		return emptyResp()
	}

	return success(mem, nil)
}

func createSkill(w http.ResponseWriter, r *http.Request) *response {
	var skillReq skillRequest
	err := json.NewDecoder(r.Body).Decode(&skillReq)
	if err != nil {
		return badRequest(err)
	}

	s := *entities.NewSkill(skillReq.Name, skillReq.Description, skillReq.Type, skillReq.Importance)

	log.Println(s)

	mem, err := entities.CreateSkill(db, s)
	if err != nil {
		return internalError(err)
	}

	return success(mem, nil)
}

func getSkillById(w http.ResponseWriter, r *http.Request) *response {
	strId := r.PathValue("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return idNotValid(strId)
	}

	mem, err := entities.GetSkillById(db, uint(id))
	if err != nil {
		return internalError(err)
	}

	if mem == nil {
		return emptyResp()
	}

	return success(mem, nil)
}

func deleteSkill(w http.ResponseWriter, r *http.Request) *response {
	strId := r.PathValue("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return idNotValid(strId)
	}

	err = entities.DeleteSkill(db, uint(id))
	if err != nil {
		return internalError(err)
	}
	return success("Done", nil)
}

func updateSkill(w http.ResponseWriter, r *http.Request) *response {
	var skillReq skillRequest
	err := json.NewDecoder(r.Body).Decode(&skillReq)
	if err != nil {
		return badRequest(err)
	}

	strId := r.PathValue("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return idNotValid(strId)
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
		return internalError(err)
	}

	return success(nT, nil)
}
