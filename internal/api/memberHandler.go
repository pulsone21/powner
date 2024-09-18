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
	Name   string                 `json:"name"`
	Age    int                    `json:"age"`
	Skills []entities.SkillRating `json:"skills"`
}

type ratingRequest struct {
	Rating  int  `json:"rating"`
	SkillId uint `json:"skill_id"`
}

func (s *ratingRequest) ValidateRating() bool {
	return s.Rating > 0 && s.Rating <= 5
}

func loadRating(r *http.Request) (*ratingRequest, error) {
	var ratReq ratingRequest
	err := json.NewDecoder(r.Body).Decode(&ratReq)
	if err != nil {
		return nil, err
	}

	if !ratReq.ValidateRating() {
		return nil, fmt.Errorf("not a valid rating: %b, needs to be between 1 - 5", ratReq.Rating)
	}
	return &ratReq, nil
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
		Name:   memReq.Name,
		Age:    memReq.Age,
		Skills: memReq.Skills,
	}

	newMem.ID = uint(id)
	err = entities.UpdateMember(db, *newMem)
	if err != nil {
		return nil, newRespErr(500, err)
	}

	return newMem, nil
}

func getSkillratingByMember(w http.ResponseWriter, r *http.Request) (any, *responseError) {
	strId := r.PathValue("id")
	memId, err := strconv.Atoi(strId)
	if err != nil {
		return nil, newRespErr(400, fmt.Errorf("id is not an uint: %v", strId))
	}

	mem, err2 := try_find_member(uint(memId))
	if err2 != nil {
		return nil, err2
	}

	return mem.Skills, nil
}

func addSkillrating(w http.ResponseWriter, r *http.Request) (any, *responseError) {
	strId := r.PathValue("id")
	memId, err := strconv.Atoi(strId)
	if err != nil {
		return nil, newRespErr(400, fmt.Errorf("id is not an uint: %v", strId))
	}

	mem, err2 := try_find_member(uint(memId))
	if err2 != nil {
		return nil, err2
	}

	ratReq, err := loadRating(r)
	if err != nil {
		return nil, newRespErr(400, err)
	}

	_, err = entities.GetSkillById(db, uint(ratReq.SkillId))
	if err != nil {
		return nil, newRespErr(400, err)
	}

	rating := entities.SkillRating{
		Rating:  ratReq.Rating,
		SkillID: ratReq.SkillId,
	}

	mem.Skills = append(mem.Skills, rating)

	err = entities.UpdateMember(db, *mem)
	if err != nil {
		return nil, newRespErr(500, err)
	}

	return mem, nil
}

func updateSkillrating(w http.ResponseWriter, r *http.Request) (any, *responseError) {
	strId := r.PathValue("rating_id")
	ratingId, err := strconv.Atoi(strId)
	if err != nil {
		return nil, newRespErr(400, fmt.Errorf("id is not an uint: %v", strId))
	}

	ratReq, err := loadRating(r)
	if err != nil {
		return nil, newRespErr(400, err)
	}

	err = entities.UpdateSkillRating(db, uint(ratingId), ratReq.Rating)
	if err != nil {
		return nil, newRespErr(400, err)
	}

	return "Done", nil
}

func getSkillrating(w http.ResponseWriter, r *http.Request) (any, *responseError) {
	strId := r.PathValue("rating_id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return nil, newRespErr(400, fmt.Errorf("id is not an uint: %v", strId))
	}

	rat, err := entities.GetSkillRating(db, uint(id))
	if err != nil {
		return nil, newRespErr(500, err)
	}

	return rat, nil
}

func deleteSkillrating(w http.ResponseWriter, r *http.Request) (any, *responseError) {
	strId := r.PathValue("id")
	_, err := strconv.Atoi(strId)
	if err != nil {
		return nil, newRespErr(400, fmt.Errorf("id is not an uint: %v", strId))
	}
	strId = r.PathValue("rating_id")
	_, err = strconv.Atoi(strId)
	if err != nil {
		return nil, newRespErr(400, fmt.Errorf("id is not an uint: %v", strId))
	}
	return nil, nil
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
