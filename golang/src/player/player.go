package player

import (
	"database/sql"
)

type Player struct {
	Id   int
	Name string
}

func ReadAll(db *sql.DB) interface{} {
	var players []Player
	rows, err := db.Query("select * from Player;")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		player := Player{}
		err = rows.Scan(&player.Id, &player.Name)
		if err != nil {
			panic(err)
		}
		players = append(players, player)
	}
	rows.Close()

	return players
}

func AddPlayer(db *sql.DB, name string) {
	_, err := db.Query("insert into Player (name) values (?);", name)
	if err != nil {
		panic(err)
	}
}

func DeletePlayer(db *sql.DB, id int) {
	_, err := db.Query("delete from Player where id = ?;", id)
	if err != nil {
		panic(err)
	}
}

func FindPlayerIDs(db *sql.DB, name []string) []int {
	var ids []int
	// プレイヤー名からプレイヤーIDを取得
	rows, err := db.Query("select id from Player where name = ?;", name)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// 取得したプレイヤーIDを配列(ids)に格納
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			panic(err)
		}
		ids = append(ids, id)
	}

	return ids
}
