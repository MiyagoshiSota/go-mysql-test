package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"test-mysql/article"
	"test-mysql/history"
	"test-mysql/player"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type PlayerName struct {
	Name string `json:"name"`
}

type OnlineUser struct {
	PlayerNames []PlayerName `json:"player_names"`
	Time        []string     `json:"time"`
}

var layout = "2006-01-02 15:04:05"

func open(path string, count uint) *sql.DB {
	db, err := sql.Open("mysql", path)
	if err != nil {
		log.Fatal("open error:", err)
	}

	if err = db.Ping(); err != nil {
		time.Sleep(time.Second * 2)
		count--
		fmt.Printf("retry... count:%v\n", count)
		return open(path, count)
	}

	fmt.Println("db connected!!")
	return db
}

func connectDB() *sql.DB {
	var path string = fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"))

	return open(path, 100)
}

// 文字列をtime型に変換
func stringToTime(str string) time.Time {
	t, _ := time.Parse(layout, str)
	return t
}

func main() {
	db := connectDB()
	defer db.Close()

	r := gin.Default()
	r.GET("/api/hello", func(c *gin.Context) {
		article.ReadAll(db)
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})
	r.GET("/api/history", func(c *gin.Context) {
		historyJSON := history.ReadAll(db)
		c.IndentedJSON(200, gin.H{
			"history": historyJSON,
		})
	})
	r.GET("/api/player", func(c *gin.Context) {
		playersJSON := player.ReadAll(db)
		c.IndentedJSON(200, gin.H{
			"players": playersJSON,
		})
	})
	r.POST("/api/player/add", func(c *gin.Context) {
		name := c.PostForm("name")
		player.AddPlayer(db, name)
		c.JSON(200, gin.H{
			"message": "success",
		})
	})
	r.POST("/api/history/set", func(c *gin.Context) {
		// リクエストをJSONに変換
		var requestJSON OnlineUser
		c.BindJSON(&requestJSON)
		fmt.Println(requestJSON)
		player_names := requestJSON.PlayerNames
		time := stringToTime(requestJSON.Time[0]) //time型に変換

		// オンラインユーザ一覧のプレイヤーIDを取得
		playerIDs := player.FindPlayerIDs(db, player_names) //player_namesが構造体の配列のためエラー

		// 過去のオンラインユーザを検索
		oldStayers := history.FindOldStayer(db)

		//オンラインユーザ一覧のプレイヤーIDを登録または更新
		if oldStayers != nil {
			var oldStayerID int
			none := oldStayers.Scan(&oldStayerID)
			if none != nil {
				// history.AddHistory(db, playerIDs, time)
			}
			// history.UpdateEndTime(db, oldStayerID, time)
		}
	})

	r.Run()
}
