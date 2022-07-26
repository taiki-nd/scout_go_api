package service

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/taiki-nd/scout_go_api/models"
)

/*
 * GetUserFromId
 * idからuser詳細情報を取得
 * @return models.User
 */
func GetUserFromId(c *fiber.Ctx) models.User {
	id, _ := strconv.Atoi(c.Params("id"))
	user := models.User{
		Id: uint(id),
	}
	return user
}
