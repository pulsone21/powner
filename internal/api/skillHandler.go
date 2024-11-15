package api

//
//import (
//	"encoding/json"
//	"fmt"
//	"net/http"
//	"strconv"
//
//	"github.com/pulsone21/powner/internal/entities"
//	"github.com/pulsone21/powner/internal/ui/partials"
//)
//
//type skillRequest struct {
//	Name        string             `json:"name"`
//	Description string             `json:"description"`
//	Type        entities.SkillType `json:"type"`
//	Importance  int                `json:"importance"`
//}
//
//func getSkills(w http.ResponseWriter, r *http.Request) *response {
//	skills, err := entities.GetSkills(db)
//	if err != nil {
//		return internalError(err)
//	}
//
//	if len(*skills) == 0 {
//		return emptyResp()
//	}
//
//	t := getTeamFromQuery(r.URL.Query().Get("teamID"))
//	if t != nil {
//		fmt.Println(t)
//		return success(skills, partials.SkillList(*skills, t, "No skills found"))
//	}
//
//	m := getMemberFromQuery(r.URL.Query().Get("memID"))
//	return success(skills, partials.SkillList(*skills, m, "No skills found"))
//}
//
//func createSkill(w http.ResponseWriter, r *http.Request) *response {
//	skillReq, err := decodeRequest[skillRequest](r)
//	if err != nil {
//		return badRequest(err)
//	}
//
//	s := *entities.NewSkill(skillReq.Name, skillReq.Description, skillReq.Type, skillReq.Importance)
//
//	sk, err := entities.CreateSkill(db, s)
//	if err != nil {
//		return internalError(err)
//	}
//	w.Header().Add("HX-Trigger", "newSkill")
//	return success(sk, partials.SkillForm())
//}
//
//func getSkillById(w http.ResponseWriter, r *http.Request) *response {
//	strId := r.PathValue("id")
//	id, err := strconv.Atoi(strId)
//	if err != nil {
//		return idNotValid(strId)
//	}
//
//	mem, err := entities.GetSkillById(db, uint(id))
//	if err != nil {
//		return internalError(err)
//	}
//
//	if mem == nil {
//		return emptyResp()
//	}
//
//	return success(mem, nil)
//}
//
//func deleteSkill(w http.ResponseWriter, r *http.Request) *response {
//	strId := r.PathValue("id")
//	id, err := strconv.Atoi(strId)
//	if err != nil {
//		return idNotValid(strId)
//	}
//
//	err = entities.DeleteSkill(db, uint(id))
//	if err != nil {
//		return internalError(err)
//	}
//	return success("Done", nil)
//}
//
//func updateSkill(w http.ResponseWriter, r *http.Request) *response {
//	var skillReq skillRequest
//	err := json.NewDecoder(r.Body).Decode(&skillReq)
//	if err != nil {
//		return badRequest(err)
//	}
//
//	strId := r.PathValue("id")
//	id, err := strconv.Atoi(strId)
//	if err != nil {
//		return idNotValid(strId)
//	}
//
//	nT := &entities.Skill{
//		Name:        skillReq.Name,
//		Description: skillReq.Description,
//		Importance:  skillReq.Importance,
//		Type:        int(skillReq.Type),
//	}
//
//	nT.ID = uint(id)
//
//	err = entities.UpdateSkill(db, *nT)
//	if err != nil {
//		return internalError(err)
//	}
//
//	return success(nT, nil)
//}
//
//func getTeamFromQuery(id string) entities.SkillHolder {
//	if id != "" {
//		tID, err := strconv.Atoi(id)
//		if err == nil {
//			t, _ := entities.GetTeamById(db, uint(tID))
//			return t
//		}
//	}
//	return nil
//}
//
//func getMemberFromQuery(id string) entities.SkillHolder {
//	if id != "" {
//		mID, err := strconv.Atoi(id)
//		if err == nil {
//			m, _ := entities.GetMemberById(db, uint(mID))
//			return m
//		}
//	}
//	return nil
//}
