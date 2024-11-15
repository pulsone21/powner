package handler

import (
	"net/http"

	"github.com/pulsone21/powner/internal/server/response"
	"github.com/pulsone21/powner/internal/service"
)

func setupApiHandler(fn response.ResponseFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fn(w, r).Respond(w, r)
	})
}

func NewMemberHandler(ser service.MemberService) MemberHandler {
	return MemberHandler{
		service: ser,
	}
}

func NewSkillHandler(ser service.SkillService) SkillHandler {
	return SkillHandler{
		service: ser,
	}
}

func NewTeamHandler(ser service.TeamService) TeamHandler {
	return TeamHandler{
		service: ser,
	}
}

// TODO: Implement that correctly
func NewUIHandler(in interface{}) UIHandler {
	return UIHandler{
		// service: ser,
	}
}
