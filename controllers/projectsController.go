package controllers

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/taiki-nd/scout_go_api/db"
	"github.com/taiki-nd/scout_go_api/models"
	"github.com/taiki-nd/scout_go_api/service"
)

/*
 * ProjectsIndex
 * project一覧を取得
 */
func ProjectsIndex(c *fiber.Ctx) error {
	log.Println("start to get projects")

	var projects []*models.Project

	// projectsレコードの取得
	err := db.DB.Find(&projects).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"project": false,
			"code":    "failed_db_get_projectes",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    projects,
		})
	}

	log.Println("success to get projects")

	return c.JSON(fiber.Map{
		"project": true,
		"code":    "success_project_index",
		"message": "",
		"data":    projects,
	})
}

/*
 * ProjectsCreate
 * DBにproject情報を登録
 */
func ProjectsCreate(c *fiber.Ctx) error {
	log.Println("start to create project")

	uuid := c.Query("uuid")
	if uuid == "" {
		log.Println("not_signin")
		return c.JSON(fiber.Map{
			"project": false,
			"code":    "not_signin",
			"message": "please signin",
			"data":    fiber.Map{},
		})
	}

	var project models.Project

	// リクエストボディのパース
	err := c.BodyParser(&project)
	if err != nil {
		log.Printf("parse error: %v", err)
		return c.JSON(fiber.Map{
			"project": false,
			"code":    "failed_parse_project_create",
			"message": fmt.Sprintf("parse error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// レコード作成
	err = db.DB.Create(&project).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"project": false,
			"code":    "failed_db_project_create",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	log.Println("success to create project")

	return c.JSON(fiber.Map{
		"project": true,
		"code":    "success_project_create",
		"message": "",
		"data":    project,
	})
}

/*
 * ProjectsShow
 * project詳細情報の取得
 */
func ProjectsShow(c *fiber.Ctx) error {
	log.Println("start to get project")

	id := c.Params("id")

	// project情報の取得
	project, err := service.GetProjectFromId(id)
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"project": false,
			"code":    "failed_db_project_show",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	log.Printf("success to get projects")

	return c.JSON(fiber.Map{
		"project": true,
		"code":    "success_project_show",
		"message": "",
		"data":    project,
	})
}

/*
 * ProjectsUpdate
 * project情報の更新
 */
func ProjectsUpdate(c *fiber.Ctx) error {
	log.Println("start to Update project")

	id := c.Params("id")
	uuid := c.Query("uuid")

	// project情報の取得
	project, err := service.GetProjectFromId(id)
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"project": false,
			"code":    "failed_db_project_show",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// userステータスの確認
	err = service.CheckUserStatus(uuid, project.WorkId)
	if err != nil {
		log.Printf("user stratus error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "user_status_error",
			"message": fmt.Sprintf("user stratus error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// リクエストボディのパース
	err = c.BodyParser(&project)
	if err != nil {
		log.Printf("parse error: %v", err)
		return c.JSON(fiber.Map{
			"project": false,
			"code":    "failed_parse_project_update",
			"message": fmt.Sprintf("parse error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// project情報の更新
	err = db.DB.Model(&project).Updates(project).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"project": false,
			"code":    "failed_db_project_update",
			"message": fmt.Sprintf("parse error: %v", err),
			"data":    fiber.Map{},
		})
	}

	log.Println("success to update project")

	return c.JSON(fiber.Map{
		"project": true,
		"code":    "success_project_update",
		"message": fmt.Sprintf("db error: %v", err),
		"data":    project,
	})
}

/*
 * ProjectsDelete
 * project情報の削除
 */
func ProjectsDelete(c *fiber.Ctx) error {
	log.Println("start to delete project")

	id := c.Params("id")
	uuid := c.Query("uuid")

	// project情報の取得
	project, err := service.GetProjectFromId(id)
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"project": false,
			"code":    "failed_db_project_show",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// userステータスの確認
	err = service.CheckUserStatus(uuid, project.WorkId)
	if err != nil {
		log.Printf("user stratus error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "user_status_error",
			"message": fmt.Sprintf("user stratus error: %v", err),
			"data":    fiber.Map{},
		})
	}

	// project情報の削除
	err = db.DB.Delete(project).Error
	if err != nil {
		log.Printf("db error: %v", err)
		return c.JSON(fiber.Map{
			"project": false,
			"code":    "failed_db_project_delete",
			"message": fmt.Sprintf("db error: %v", err),
			"data":    fiber.Map{},
		})
	}

	log.Println("success to delete project")

	return c.JSON(fiber.Map{
		"project": true,
		"code":    "success_project_delete",
		"message": "",
		"data":    fiber.Map{},
	})
}
