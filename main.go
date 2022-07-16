package main

import (
	"log"

	"github.com/taiki-nd/scout_go_api/config"
	"github.com/taiki-nd/scout_go_api/utils"
)

func main() {
	utils.Logging(config.Config.LogFile)

	log.Println("test_log")
}
