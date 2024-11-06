package api

import (
	"net/http"
)

func getRoutes() http.Handler {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./public/static/"))

	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	api := http.NewServeMux()
	api.HandleFunc("GET /member", setupHandler(getMembers))
	api.HandleFunc("POST /member", setupHandler(createMember))
	api.HandleFunc("GET /member/{id}", setupHandler(getMemberById))
	api.HandleFunc("DELETE /member/{id}", setupHandler(deleteMember))
	api.HandleFunc("PUT /member/{id}", setupHandler(updateMember))
	api.HandleFunc("GET /member/{id}/skill", setupHandler(getSkillratingByMember))
	api.HandleFunc("POST /member/{id}/skill/{skill_id}", setupHandler(addSkillrating))
	api.HandleFunc("DELETE /member/{id}/skill/{skill_id}", setupHandler(deleteSkillrating))
	api.HandleFunc("GET /member/{id}/skillrating/{rating_id}/{rating}", setupHandler(updateSkillrating))
	api.HandleFunc("GET /member/{id}/skillrating/{rating_id}", setupHandler(getSkillrating))
	api.HandleFunc("POST /member/{mem_id}/team/{id}", setupHandler(addMember))
	api.HandleFunc("DELETE /member/{mem_id}/team/{id}", setupHandler(removeMember))

	api.HandleFunc("GET /skill", setupHandler(getSkills))
	api.HandleFunc("POST /skill", setupHandler(createSkill))
	api.HandleFunc("GET /skill/{id}", setupHandler(getSkillById))
	api.HandleFunc("DELETE /skill/{id}", setupHandler(deleteSkill))
	api.HandleFunc("PUT /skill/{id}", setupHandler(updateSkill))

	api.HandleFunc("GET /team", setupHandler(getTeams))
	api.HandleFunc("POST /team", setupHandler(createTeam))
	api.HandleFunc("GET /team/{id}", setupHandler(getTeamById))
	api.HandleFunc("DELETE /team/{id}", setupHandler(deleteTeam))
	api.HandleFunc("PUT /team/{id}", setupHandler(updateTeam))
	api.HandleFunc("GET /team/{id}/diagrams", setupHandler(getDiagrams))
	api.HandleFunc("GET /team/{id}/member", setupHandler(getMemberByTeam))
	api.HandleFunc("POST /team/{id}/member/{mem_id}", setupHandler(addMember))
	api.HandleFunc("DELETE /team/{id}/member/{mem_id}", setupHandler(removeMember))
	api.HandleFunc("GET /team/{id}/skill", setupHandler(getSkillsByTeam))
	api.HandleFunc("POST /team/{id}/skill/{skill_id}", setupHandler(addSkill))
	api.HandleFunc("DELETE /team/{id}/skill/{skill_id}", setupHandler(removeSkill))

	mux.Handle("/api/", http.StripPrefix("/api", api))

	mux.HandleFunc("/", indexPage)
	mux.HandleFunc("GET /modal/newTeam", serveNewTeamModal)
	mux.HandleFunc("GET /modal/member/{id}", serveMemberModal)
	mux.HandleFunc("GET /modal/skill/{id}", serverSkillModal)
	return mux
}

func setupHandler(fn responseFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fn(w, r).Respond(w, r)
	})
}
