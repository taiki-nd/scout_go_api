package main

import (
	"github.com/taiki-nd/scout_go_api/config"
	"github.com/taiki-nd/scout_go_api/db"
	"github.com/taiki-nd/scout_go_api/utils"
)

func main() {
	// log出力の有効化
	utils.Logging(config.Config.LogFile)

	// db接続
	db.ConnectToDb()
}
