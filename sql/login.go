package sql

import "database/sql"

type login struct {
	id       int
	username string
	password string
}

func QueryLogin(username string, password string) bool {
	row := Db.QueryRow("SELECT * FROM `login` WHERE student_id='" + username + "' AND password='" + password + "'")

	var l login

	if err := row.Scan(&l.id, &l.username, &l.password); err == sql.ErrNoRows {
		return false
	} else if err != nil {
		panic(err)
	}

	return true
}
