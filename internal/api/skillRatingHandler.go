package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/pulsone21/powner/internal/entities"
)

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

func getSkillratingByMember(w http.ResponseWriter, r *http.Request) *response {
	strId := r.PathValue("id")
	memId, err := strconv.Atoi(strId)
	if err != nil {
		return idNotValid(strId)
	}

	mem, err2 := try_find_member(uint(memId))
	if err2 != nil {
		return err2
	}

	return success(mem.Skills, nil)
}

func addSkillrating(w http.ResponseWriter, r *http.Request) *response {
	strId := r.PathValue("id")
	memId, err := strconv.Atoi(strId)
	if err != nil {
		return idNotValid(strId)
	}

	mem, err2 := try_find_member(uint(memId))
	if err2 != nil {
		return err2
	}

	ratReq, err := loadRating(r)
	if err != nil {
		return badRequest(err)
	}

	_, err = entities.GetSkillById(db, uint(ratReq.SkillId))
	if err != nil {
		return badRequest(err)
	}

	rating := entities.SkillRating{
		Rating:  ratReq.Rating,
		SkillID: ratReq.SkillId,
	}

	mem.Skills = append(mem.Skills, rating)

	err = entities.UpdateMember(db, *mem)
	if err != nil {
		return internalError(err)
	}

	return success(mem, nil)
}

func updateSkillrating(w http.ResponseWriter, r *http.Request) *response {
	strId := r.PathValue("rating_id")
	ratingId, err := strconv.Atoi(strId)
	if err != nil {
		return idNotValid(strId)
	}

	ratReq, err := loadRating(r)
	if err != nil {
		return badRequest(err)
	}

	err = entities.UpdateSkillRating(db, uint(ratingId), ratReq.Rating)
	if err != nil {
		return badRequest(err)
	}

	return success("Done", nil)
}

func getSkillrating(w http.ResponseWriter, r *http.Request) *response {
	strId := r.PathValue("rating_id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return idNotValid(strId)
	}

	rat, err := entities.GetSkillRating(db, uint(id))
	if err != nil {
		return internalError(err)
	}

	return success(rat, nil)
}

func deleteSkillrating(w http.ResponseWriter, r *http.Request) *response {
	strId := r.PathValue("rating_id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return idNotValid(strId)
	}

	err = entities.DeleteSkillRating(db, uint(id))
	if err != nil {
		return internalError(err)
	}

	return success("Done", nil)
}
