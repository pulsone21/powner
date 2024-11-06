package api

// TODO: find a way to get dynamicly the correct type from the URL PathValue
/* TODO: and refactor all of these blocks

strId = r.PathValue("skill_id")
skillID, err := strconv.Atoi(strId)
if err != nil {
	return idNotValid(strId)
}

s, err := entities.GetSkillById(db, uint(skillID))
if err != nil {
	return badRequest(err)
}

*/
