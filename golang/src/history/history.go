package history

import (
	"database/sql"
	"time"
)

type History struct {
	Id         int
	Player_id  int
	Start_time time.Time
	End_time   time.Time
}

func ReadAll(db *sql.DB) interface{} {
	var history_array []History
	rows, err := db.Query("select * from History;")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		history := History{}
		err = rows.Scan(&history.Id, &history.Player_id, &history.Start_time, &history.End_time)
		if err != nil {
			panic(err)
		}
		history_array = append(history_array, history)
	}
	rows.Close()

	return history_array
}

// 履歴の追加
func AddHistory(db *sql.DB, player_id int, start_time time.Time) {
	_, err := db.Query("insert into History (player_id, start_time, end_time) values (1, '2020-01-01 00:00:00', '2020-01-01 00:00:00');")
	if err != nil {
		panic(err)
	}
}

// 前のオンラインユーザを検索
func FindOldStayer(db *sql.DB) *sql.Row {
	return db.QueryRow("select player_id from History where end_time = null;")
}

// 終了時間の更新
func UpdateEndTime(db *sql.DB, id int, end_time time.Time) {
	_, err := db.Query("update History set end_time = ? where id = ?;", end_time, id)
	if err != nil {
		panic(err)
	}
}
