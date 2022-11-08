package sql

import "fmt"

type Student struct {
	studentId string
	Name      string
	Gender    bool
	Grade     string
	Major     string
	Contact   string
	Room      string
}

func QueryStudent(studentId string) (s Student) {
	queryString := "SELECT student.student_id, student.Name, student.gender, student.grade, major.Name, " +
		"student.contact, student.room FROM `student`\nINNER JOIN major ON student.major=major.id WHERE student.student_id=" +
		"'" + studentId + "'"
	fmt.Println(queryString)
	row := Db.QueryRow(queryString)

	if err := row.Scan(&s.studentId, &s.Name, &s.Gender, &s.Grade, &s.Major, &s.Contact, &s.Room); err != nil {
		panic(err)
	}
	return s
}
