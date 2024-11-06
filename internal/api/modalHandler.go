package api

import (
	"net/http"
	"strconv"

	"github.com/pulsone21/powner/internal/entities"
	"github.com/pulsone21/powner/internal/ui/modals"
)

func serveMemberModal(w http.ResponseWriter, r *http.Request) {
	strId := r.PathValue("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		writeError(err, w)
		return
	}

	t, err := entities.GetTeamById(db, uint(id))
	if err != nil {
		writeError(err, w)
		return
	}

	mems, err := entities.GetMembers(db)
	if err != nil {
		writeError(err, w)
		return
	}

	modals.MemberModal(*mems, t).Render(r.Context(), w)
}

func serverSkillModal(w http.ResponseWriter, r *http.Request) {
	strId := r.PathValue("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		writeError(err, w)
		return
	}

	t, err := entities.GetTeamById(db, uint(id))
	if err != nil {
		writeError(err, w)
		return
	}
	skills, err := entities.GetSkills(db)
	if err != nil {
		writeError(err, w)
		return
	}
	modals.SkillModal(*skills, t).Render(r.Context(), w)
}

func serveNewTeamModal(w http.ResponseWriter, r *http.Request) {
	modals.NewTeamModal().Render(r.Context(), w)
}
