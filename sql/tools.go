package sql

import (
	"database/sql"
	"dormitory_interface/redis"
	"log"
)

func UpdatePassword(username string, password string) {
	_, err := Db.Exec("UPDATE `login` SET password=? WHERE student_id=?", password, username)
	if err != nil {
		panic(err)
	}
}

func QueryBuildingList(username string) (buildings []int) {
	queryString := "SELECT building.building_id FROM `building`\nINNER JOIN student ON building.gender=student.gender WHERE student.student_id=?"
	rows, err := Db.Query(queryString, username)
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

func QueryAvailableCount(building string) (cnt int) {
	cnt, err := redis.QueryAvaliableCountviaRedis(building)
	if err != nil {
		log.Println("Redis query failed, fallback to SQL")
		cnt = QueryAvailableCountViaSql(building)
		redis.UpdateAvailableCount(building, cnt)
	} else {
		log.Println("Redis query success")
	}
	return cnt
}
func QueryAvailableCountViaSql(building string) (cnt int) {
	queryString := "SELECT occupancy.capacity, occupancy.occupied FROM `room` \nINNER JOIN occupancy" +
		" On occupancy.id=room.id\nWHERE room.building=" + string(building)
	print(queryString)
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

func InsertApplicationRaw(username string, buildingNo string) {
	if _, err := Db.Exec("INSERT INTO `raw_order` (`id`, `student_id`, `building_id`) VALUES (NULL, ?, ?);", username, buildingNo); err != nil {
		panic(err)
	}
}
