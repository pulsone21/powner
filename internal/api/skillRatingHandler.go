package api

//
//import (
//	"fmt"
//	"net/http"
//	"strconv"
//
//	"github.com/pulsone21/powner/internal/entities"
//	"github.com/pulsone21/powner/internal/ui/partials"
//)
//
//type ratingRequest struct {
//	Rating  int  `json:"rating"`
//	SkillId uint `json:"skill_id"`
//}
//
//func (s *ratingRequest) ValidateRating() bool {
//	return s.Rating >= 0 && s.Rating <= 5
//}
//
//func getSkillratingByMember(w http.ResponseWriter, r *http.Request) *response {
//	strId := r.PathValue("id")
//	memId, err := strconv.Atoi(strId)
//	if err != nil {
//		return idNotValid(strId)
//	}
//
//	mem, err2 := try_find_member(uint(memId))
//	if err2 != nil {
//		return err2
//	}
//
//	return success(mem.Skills, partials.SkillAdjustList(*mem))
//}
//
//func addSkillrating(w http.ResponseWriter, r *http.Request) *response {
//	strId := r.PathValue("id")
//	memId, err := strconv.Atoi(strId)
//	if err != nil {
//		return idNotValid(strId)
//	}
//
//	mem, err2 := try_find_member(uint(memId))
//	if err2 != nil {
//		return err2
//	}
//
//	strId = r.PathValue("skill_id")
//	skillID, err := strconv.Atoi(strId)
//	if err != nil {
//		return idNotValid(strId)
//	}
//
//	s, err := entities.GetSkillById(db, uint(skillID))
//	if err != nil {
//		return badRequest(err)
//	}
//
//	rating := entities.SkillRating{
//		Rating: 0,
//		Skill:  *s,
//	}
//
//	mem.Skills = append(mem.Skills, rating)
//
//	err = entities.UpdateMember(db, *mem)
//	if err != nil {
//		return internalError(err)
//	}
//	w.Header().Add("HX-Trigger", "skillRatingChange")
//	return success(mem, partials.SkillListItem(*s, string(mem.GetID()), mem.GetType(), mem.HasSkill(uint(skillID))))
//}
//
//func updateSkillrating(w http.ResponseWriter, r *http.Request) *response {
//	strId := r.PathValue("rating_id")
//	ratingId, err := strconv.Atoi(strId)
//	if err != nil {
//		return idNotValid(strId)
//	}
//
//	rating, err := strconv.Atoi(r.PathValue("rating"))
//	if err != nil {
//		return newResponse(nil, nil, 400, fmt.Errorf("couldn't parse rating input to int"))
//	}
//
//	ratReq := &ratingRequest{
//		SkillId: uint(ratingId),
//		Rating:  rating,
//	}
//
//	if ratReq.ValidateRating() != true {
//		return badRequest(fmt.Errorf("not a valid rating: %b, needs to be between 1 - 5", ratReq.Rating))
//	}
//
//	err = entities.UpdateSkillRating(db, uint(ratingId), ratReq.Rating)
//	if err != nil {
//		return badRequest(err)
//	}
//
//	s, err := entities.GetSkillRating(db, uint(ratingId))
//	if err != nil {
//		return badRequest(err)
//	}
//
//	fmt.Printf("%+v\n", s)
//
//	w.Header().Add("HX-Trigger", "skillRatingChange")
//	return success("Done", partials.SkillAddjustItem(r.PathValue("id"), *s))
//}
//
//func getSkillrating(w http.ResponseWriter, r *http.Request) *response {
//	strId := r.PathValue("rating_id")
//	id, err := strconv.Atoi(strId)
//	if err != nil {
//		return idNotValid(strId)
//	}
//
//	rat, err := entities.GetSkillRating(db, uint(id))
//	if err != nil {
//		return internalError(err)
//	}
//
//	return success(rat, nil)
//}
//
//func deleteSkillrating(w http.ResponseWriter, r *http.Request) *response {
//	strId := r.PathValue("skill_id")
//	id, err := strconv.Atoi(strId)
//	if err != nil {
//		return idNotValid(strId)
//	}
//
//	s, err := entities.GetSkillById(db, uint(id))
//
//	strId = r.PathValue("id")
//	memId, err := strconv.Atoi(strId)
//	if err != nil {
//		return idNotValid(strId)
//	}
//
//	mem, err2 := try_find_member(uint(memId))
//	if err2 != nil {
//		return err2
//	}
//
//	sR := mem.GetSkillRatingBySkill(s.ID)
//	if sR == nil {
//		return badRequest(fmt.Errorf("Member don't has the Skill: %v", s.Name))
//	}
//
//	err = entities.DeleteSkillRating(db, sR.ID)
//	if err != nil {
//		return internalError(err)
//	}
//
//	w.Header().Add("HX-Trigger", "skillRatingChange")
//	return success("Done", partials.SkillListItem(*s, string(mem.GetID()), mem.GetType(), false))
//}
