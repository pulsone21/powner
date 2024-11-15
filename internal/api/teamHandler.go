package api

//
//import (
//	"encoding/json"
//	"fmt"
//	"net/http"
//	"strconv"
//	"strings"
//
//	"github.com/pulsone21/powner/internal/entities"
//	"github.com/pulsone21/powner/internal/ui/charts"
//	"github.com/pulsone21/powner/internal/ui/partials"
//	"github.com/pulsone21/powner/internal/ui/subpage"
//)
//
//type teamRequest struct {
//	Name        string `json:"name"`
//	Description string `json:"description"`
//}
//
//func getTeams(w http.ResponseWriter, r *http.Request) *response {
//	team, err := entities.GetTeams(db)
//	if err != nil {
//		return internalError(err)
//	}
//
//	if len(*team) == 0 {
//		return emptyResp()
//	}
//
//	mem := getMemberFromQuery(r.URL.Query().Get("memID"))
//	if mem != nil {
//		return success(team, partials.TeamList(*team, mem.GetID(), "No teams found"))
//	}
//	return success(team, partials.TeamNavbarList(*team))
//}
//
//func createTeam(w http.ResponseWriter, r *http.Request) *response {
//	teamReq, err := decodeRequest[teamRequest](r)
//	if err != nil {
//		return badRequest(err)
//	}
//
//	mem, err := entities.CreateTeam(db, *entities.NewTeam(teamReq.Name, teamReq.Description))
//	if err != nil {
//		return internalError(err)
//	}
//
//	w.Header().Add("HX-Trigger", "newTeam")
//	return success(mem, nil)
//}
//
//func getTeamById(w http.ResponseWriter, r *http.Request) *response {
//	strId := r.PathValue("id")
//	id, err := strconv.Atoi(strId)
//	if err != nil {
//		return idNotValid(strId)
//	}
//
//	mem, err := entities.GetTeamById(db, uint(id))
//	if err != nil {
//		return internalError(err)
//	}
//
//	if mem == nil {
//		return emptyResp()
//	}
//
//	return success(mem, subpage.Team(*mem))
//}
//
//func deleteTeam(w http.ResponseWriter, r *http.Request) *response {
//	strId := r.PathValue("id")
//	id, err := strconv.Atoi(strId)
//	if err != nil {
//		return idNotValid(strId)
//	}
//	err = entities.DeleteTeam(db, uint(id))
//	if err != nil {
//		return internalError(err)
//	}
//
//	return success("Done", nil)
//}
//
//func updateTeam(w http.ResponseWriter, r *http.Request) *response {
//	var teamReq teamRequest
//	err := json.NewDecoder(r.Body).Decode(&teamReq)
//	if err != nil {
//		return newResponse(nil, nil, 400, err)
//	}
//
//	strId := r.PathValue("id")
//	id, err := strconv.Atoi(strId)
//	if err != nil {
//		return idNotValid(strId)
//	}
//
//	nT := entities.Team{
//		Name:        teamReq.Name,
//		Description: teamReq.Description,
//	}
//	nT.ID = uint(id)
//
//	err = entities.UpdateTeam(db, nT)
//	if err != nil {
//		return internalError(err)
//	}
//
//	return success(nT, nil)
//}
//
//func addMember(w http.ResponseWriter, r *http.Request) *response {
//	teamId := r.PathValue("id")
//	id, err := strconv.Atoi(teamId)
//	if err != nil {
//		return idNotValid(teamId)
//	}
//
//	strId := r.PathValue("mem_id")
//	mem_id, err := strconv.Atoi(strId)
//	if err != nil {
//		return idNotValid(strId)
//	}
//
//	err = entities.AddMemberToTeam(db, uint(id), uint(mem_id))
//	if err != nil {
//		return badRequest(err)
//	}
//
//	w.Header().Add("HX-Trigger", "teamMemberChange")
//	fmt.Println(r.URL.Path)
//	// if we request out of the member api we want to return the Team html
//	if strings.HasPrefix(r.URL.Path, "/member/") {
//		fmt.Println("Asking from Member API, returning Team Items")
//		t, err := entities.GetTeamById(db, uint(id))
//		if err != nil {
//			return badRequest(err)
//		}
//
//		return success("Done", partials.TeamListItem(*t, fmt.Sprint(mem_id), true))
//	}
//
//	mem, err := entities.GetMemberById(db, uint(mem_id))
//	if err != nil {
//		return badRequest(err)
//	}
//
//	return success("Done", partials.MemberListItem(*mem, teamId, true))
//}
//
//func removeMember(w http.ResponseWriter, r *http.Request) *response {
//	teamId := r.PathValue("id")
//	id, err := strconv.Atoi(teamId)
//	if err != nil {
//		return idNotValid(teamId)
//	}
//
//	strId := r.PathValue("mem_id")
//	mem_id, err := strconv.Atoi(strId)
//	if err != nil {
//		return idNotValid(strId)
//	}
//
//	err = entities.RemoveMemberFromTeam(db, uint(id), uint(mem_id))
//	if err != nil {
//		return newResponse(nil, nil, 400, err)
//	}
//
//	w.Header().Add("HX-Trigger", "teamMemberChange")
//	// return success("Done", partials.MemberListItem(*mem, teamId, false))
//	fmt.Println(r.URL.Path)
//	// if we request out of the member api we want to return the Team html
//	if strings.HasPrefix(r.URL.Path, "/member/") {
//		fmt.Println("Asking from Member API, returning Team Items")
//		t, err := entities.GetTeamById(db, uint(id))
//		if err != nil {
//			return badRequest(err)
//		}
//
//		return success("Done", partials.TeamListItem(*t, fmt.Sprint(mem_id), false))
//	}
//
//	mem, err := entities.GetMemberById(db, uint(mem_id))
//	if err != nil {
//		return badRequest(err)
//	}
//
//	return success("Done", partials.MemberListItem(*mem, teamId, false))
//}
//
//func addSkill(w http.ResponseWriter, r *http.Request) *response {
//	team_id := r.PathValue("id")
//	id, err := strconv.Atoi(team_id)
//	if err != nil {
//		return idNotValid(team_id)
//	}
//
//	strId := r.PathValue("skill_id")
//	skill_id, err := strconv.Atoi(strId)
//	if err != nil {
//		return idNotValid(strId)
//	}
//
//	err = entities.AddSkillToTeam(db, uint(id), uint(skill_id))
//	fmt.Println(err)
//	if err != nil {
//		return newResponse(nil, nil, 400, err)
//	}
//
//	s, err := entities.GetSkillById(db, uint(skill_id))
//	if err != nil {
//		return newResponse(nil, nil, 400, err)
//	}
//
//	w.Header().Add("HX-Trigger", "skillChange")
//	return success("Done", partials.SkillListItem(*s, team_id, "team", true))
//}
//
//func removeSkill(w http.ResponseWriter, r *http.Request) *response {
//	team_id := r.PathValue("id")
//	id, err := strconv.Atoi(team_id)
//	if err != nil {
//		return idNotValid(team_id)
//	}
//
//	strId := r.PathValue("skill_id")
//	skill_id, err := strconv.Atoi(strId)
//	if err != nil {
//		return idNotValid(strId)
//	}
//
//	err = entities.RemoveSkillFromTeam(db, uint(id), uint(skill_id))
//	fmt.Println(err)
//	if err != nil {
//		return newResponse(nil, nil, 400, err)
//	}
//
//	s, err := entities.GetSkillById(db, uint(skill_id))
//	if err != nil {
//		return newResponse(nil, nil, 400, err)
//	}
//
//	w.Header().Add("HX-Trigger", "skillChange")
//	return success("Done", partials.SkillListItem(*s, team_id, "team", false))
//}
//
//func getMemberByTeam(w http.ResponseWriter, r *http.Request) *response {
//	strId := r.PathValue("id")
//	id, err := strconv.Atoi(strId)
//	if err != nil {
//		return idNotValid(strId)
//	}
//
//	t, err := entities.GetTeamById(db, uint(id))
//	if err != nil {
//		return newResponse(nil, nil, 400, err)
//	}
//
//	return success(t.Members, partials.MemberList(t.Members, t, "No members on the team"))
//}
//
//func getSkillsByTeam(w http.ResponseWriter, r *http.Request) *response {
//	strId := r.PathValue("id")
//	id, err := strconv.Atoi(strId)
//	if err != nil {
//		return idNotValid(strId)
//	}
//
//	t, err := entities.GetTeamById(db, uint(id))
//	if err != nil {
//		return newResponse(nil, nil, 400, err)
//	}
//
//	return success(t.Skills, partials.SkillList(t.Skills, t, "No skills on the team"))
//}
//
//func getDiagrams(w http.ResponseWriter, r *http.Request) *response {
//	strId := r.PathValue("id")
//	id, err := strconv.Atoi(strId)
//	if err != nil {
//		return idNotValid(strId)
//	}
//
//	t, err := entities.GetTeamById(db, uint(id))
//	if err != nil {
//		return newResponse(nil, nil, 400, err)
//	}
//
//	return success(t.Skills, charts.DiagramList(*t))
//}
