package api

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/pulsone21/powner/internal/entities"
)

type teamRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func getTeams(w http.ResponseWriter, r *http.Request) *response {
	mem, err := entities.GetTeams(db)
	if err != nil {
		return internalError(err)
	}

	if len(*mem) == 0 {
		return emptyResp()
	}

	return success(mem, nil)
}

func createTeam(w http.ResponseWriter, r *http.Request) *response {
	var teamReq teamRequest
	err := json.NewDecoder(r.Body).Decode(&teamReq)
	if err != nil {
		return newResponse(nil, nil, 400, err)
	}

	mem, err := entities.CreateTeam(db, *entities.NewTeam(teamReq.Name, teamReq.Description))
	if err != nil {
		return internalError(err)
	}

	return success(mem, nil)
}

func getTeamById(w http.ResponseWriter, r *http.Request) *response {
	strId := r.PathValue("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return idNotValid(strId)
	}

	mem, err := entities.GetTeamById(db, uint(id))
	if err != nil {
		return internalError(err)
	}

	if mem == nil {
		return emptyResp()
	}

	return success(mem, nil)
}

func deleteTeam(w http.ResponseWriter, r *http.Request) *response {
	strId := r.PathValue("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return idNotValid(strId)
	}
	err = entities.DeleteTeam(db, uint(id))
	if err != nil {
		return internalError(err)
	}

	return success("Done", nil)
}

func updateTeam(w http.ResponseWriter, r *http.Request) *response {
	var teamReq teamRequest
	err := json.NewDecoder(r.Body).Decode(&teamReq)
	if err != nil {
		return newResponse(nil, nil, 400, err)
	}

	strId := r.PathValue("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return idNotValid(strId)
	}

	nT := entities.Team{
		Name:        teamReq.Name,
		Description: teamReq.Description,
	}
	nT.ID = uint(id)

	err = entities.UpdateTeam(db, nT)
	if err != nil {
		return internalError(err)
	}

	return success(nT, nil)
}

func addMember(w http.ResponseWriter, r *http.Request) *response {
	strId := r.PathValue("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return idNotValid(strId)
	}

	strId = r.PathValue("mem_id")
	mem_id, err := strconv.Atoi(strId)
	if err != nil {
		return idNotValid(strId)
	}

	err = entities.AddMemberToTeam(db, uint(id), uint(mem_id))
	if err != nil {
		return newResponse(nil, nil, 400, err)
	}

	return success("Done", nil)
}

func removeMember(w http.ResponseWriter, r *http.Request) *response {
	strId := r.PathValue("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return idNotValid(strId)
	}

	strId = r.PathValue("mem_id")
	mem_id, err := strconv.Atoi(strId)
	if err != nil {
		return idNotValid(strId)
	}

	err = entities.RemoveMemberFromTeam(db, uint(id), uint(mem_id))
	if err != nil {
		return newResponse(nil, nil, 400, err)
	}

	return success("Done", nil)
}

func addSkill(w http.ResponseWriter, r *http.Request) *response {
	strId := r.PathValue("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return idNotValid(strId)
	}

	strId = r.PathValue("skill_id")
	skill_id, err := strconv.Atoi(strId)
	if err != nil {
		return idNotValid(strId)
	}

	err = entities.AddSkillToTeam(db, uint(id), uint(skill_id))
	fmt.Println(err)
	if err != nil {
		return newResponse(nil, nil, 400, err)
	}

	slog.Info(fmt.Sprintf("Add this point we should have added the skill with id: %b to team id: %b", skill_id, id))
	return success("Done", nil)
}

func removeSkill(w http.ResponseWriter, r *http.Request) *response {
	strId := r.PathValue("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return idNotValid(strId)
	}

	strId = r.PathValue("skill_id")
	skill_id, err := strconv.Atoi(strId)
	if err != nil {
		return idNotValid(strId)
	}

	err = entities.RemoveSkillFromTeam(db, uint(id), uint(skill_id))
	fmt.Println(err)
	if err != nil {
		return newResponse(nil, nil, 400, err)
	}

	return success("Done", nil)
}
