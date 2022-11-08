package sql

import (
	"database/sql"
	"encoding/json"
)

func QueryStudentIncluded(username string) (teamId int) {
	row := Db.QueryRow("SELECT affiliation.team_id FROM `affiliation` \nINNER JOIN `student` ON "+
		"affiliation.student_id=student.id\nWHERE student.student_id=?;", username)
	if err := row.Scan(&teamId); err == sql.ErrNoRows {
		return -1
	} else if err != nil {
		panic(err)
	}
	return teamId
}

func CreateNewTeam(username string) {
	var Ids [1]int
	var gender bool
	if err := Db.QueryRow("SELECT id, gender FROM `student` WHERE student_id=?", username).Scan(&Ids[0], &gender); err != nil {
		panic(err)
	}
	studentIds, err := json.Marshal(Ids)
	if err != nil {
		panic(err)
	}
	var teamId int64
	if result, err := Db.Exec("INSERT INTO `team` (`id`, `student_ids`, `size`, `gender`) VALUES (NULL, ?, 1, ?);", studentIds, gender); err != nil {
		panic(err)
	} else if teamId, err = result.LastInsertId(); err != nil {
		panic(err)
	}
	if _, err := Db.Exec("INSERT INTO `affiliation` (`student_id`, `team_id`) VALUES (?, ?);", Ids[0], teamId); err != nil {
		panic(err)
	}
}
