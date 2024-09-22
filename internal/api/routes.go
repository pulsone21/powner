package api

import (
	"net/http"
)

func getRoutes() http.Handler {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./public/static/"))

	mux.HandleFunc("/", serveHtml)
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	api := http.NewServeMux()
	api.HandleFunc("GET /member", setupHandler(getMembers))
	api.HandleFunc("POST /member", setupHandler(createMember))
	api.HandleFunc("GET /member/{id}", setupHandler(getMemberById))
	api.HandleFunc("DELETE /member/{id}", setupHandler(deleteMember))
	api.HandleFunc("PUT /member/{id}", setupHandler(updateMember))
	api.HandleFunc("GET /member/{id}/skillrating", setupHandler(getSkillratingByMember))
	api.HandleFunc("POST /member/{id}/skillrating", setupHandler(addSkillrating))
	api.HandleFunc("PUT /member/{id}/skillrating/{rating_id}", setupHandler(updateSkillrating))
	api.HandleFunc("GET /member/{id}/skillrating/{rating_id}", setupHandler(getSkillrating))
	api.HandleFunc("DELETE /member/{id}/skillrating/{rating_id}", setupHandler(deleteSkillrating))

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
	api.HandleFunc("POST /team/{id}/member/{mem_id}", setupHandler(addMember))
	api.HandleFunc("DELETE /team/{id}/member/{mem_id}", setupHandler(removeMember))
	api.HandleFunc("POST /team/{id}/skill/{skill_id}", setupHandler(addSkill))
	api.HandleFunc("DELETE /team/{id}/skill/{skill_id}", setupHandler(removeSkill))

	mux.Handle("/api/", http.StripPrefix("/api", api))

	return mux
}

func setupHandler(fn responseFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fn(w, r).Respond(w, r)
	})
}
