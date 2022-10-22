package sql

import "database/sql"

func UpdatePassword(username string, password string) {
	_, err := Db.Exec("UPDATE `login` SET password='" + password + "' WHERE student_id='" + username + "'")
	if err != nil {
		panic(err)
	}
}

func QueryBuildingList(username string) (buildings []int) {
	queryString := "SELECT building.building_id FROM `building`\nINNER JOIN student ON building.gender=student.gender " +
		"WHERE student.student_id='" + username + "'"
	rows, err := Db.Query(queryString)
	if err != nil {
		panic(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	buildings = make([]int, 0)
	for rows.Next() {
		var building int
		err = rows.Scan(&building)
		if err != nil {
			panic(err)
		}
		buildings = append(buildings, building)
	}
	return buildings
}

func QueryAvaliableCount(building string) (cnt int) {
	queryString := "SELECT occupancy.capacity, occupancy.occupied FROM `room` \nINNER JOIN occupancy" +
		" On occupancy.id=room.id\nWHERE room.building=" + string(building)
	rows, err := Db.Query(queryString)
	if err != nil {
		panic(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	cnt = 0
	for rows.Next() {
		var capacity, occupied int
		err = rows.Scan(&capacity, &occupied)
		if err != nil {
			panic(err)
		}
		cnt += capacity - occupied
	}
	return cnt
}
