package api

import (
	"net/http"
	"strconv"

	"github.com/pulsone21/powner/internal/entities"
	"github.com/pulsone21/powner/internal/ui/modals"
	"github.com/pulsone21/powner/internal/ui/partials"
)

type memberRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func getMembers(w http.ResponseWriter, r *http.Request) *response {
	mem, err := entities.GetMembers(db)
	if err != nil {
		return internalError(err)
	}

	if len(*mem) == 0 {
		return newResponse(nil, nil, 200, nil)
	}

	// return success(mem, partials.MemberList(*mem, nil, "No member found"))
	return success(mem, nil)
}

func createMember(w http.ResponseWriter, r *http.Request) *response {
	memReq, err := decodeRequest[memberRequest](r)
	if err != nil {
		return badRequest(err)
	}

	mem, err := entities.CreateMember(db, *entities.NewMember(memReq.Name, memReq.Age))
	if err != nil {
		return internalError(err)
	}

	w.Header().Add("HX-Trigger", "newMember")
	return newResponse(mem, partials.MemberForm(), 201, nil)
}

func getMemberById(w http.ResponseWriter, r *http.Request) *response {
	strId := r.PathValue("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return idNotValid(strId)
	}
	mem, err := entities.GetMemberById(db, uint(id))
	if err != nil {
		return internalError(err)
	}

	if mem == nil {
		return newResponse(nil, nil, 200, nil)
	}

	return success(mem, modals.MemberInfoModal(*mem))
}

func deleteMember(w http.ResponseWriter, r *http.Request) *response {
	strId := r.PathValue("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return idNotValid(strId)
	}
	err = entities.DeleteMember(db, uint(id))
	if err != nil {
		return internalError(err)
	}
	return success("Done", nil)
}

func updateMember(w http.ResponseWriter, r *http.Request) *response {
	memReq, err := decodeRequest[memberRequest](r)

	strId := r.PathValue("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return idNotValid(strId)
	}

	newMem := &entities.Member{
		Name: memReq.Name,
		Age:  memReq.Age,
	}

	newMem.ID = uint(id)
	err = entities.UpdateMember(db, *newMem)
	if err != nil {
		return internalError(err)
	}

	return success(newMem, nil)
}

func try_find_member(id uint) (*entities.Member, *response) {
	mem, err := entities.GetMemberById(db, id)
	if err != nil {
		return nil, internalError(err)
	}

	if mem == nil {
		return nil, idNotValid(string(id))
	}

	return mem, nil
}
