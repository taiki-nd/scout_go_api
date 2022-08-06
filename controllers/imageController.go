package controllers

import (
	"context"
	"fmt"
	"image"
	"io"
	"log"
	"os"
	"time"

	"github.com/taiki-nd/scout_go_api/config"
	"github.com/taiki-nd/scout_go_api/service"

	"cloud.google.com/go/storage"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/api/option"
)

/*
 * ImageUpload
 * 画像をgcsにアップロード
 */
func ImageUpload(c *fiber.Ctx) error {
	log.Println("start to upload image")
	file, err := c.FormFile("image")
	if err != nil {
		log.Printf("uploads images error: %s", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "error_upload_image",
			"message": fmt.Sprintf("failed to upload image: %v", err),
			"data":    fiber.Map{},
		})
	}

	filename := file.Filename

	fileData, err := file.Open()
	if err != nil {
		log.Printf("failed to open image: %s", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "file_open_error",
			"message": fmt.Sprintf("failed to open image: %v", err),
			"data":    fiber.Map{},
		})
	}

	// 画像をimage.Image型にdecode
	img, data, err := image.Decode(fileData)
	if err != nil {
		log.Printf("failed to decode image: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "image_decode_error",
			"message": fmt.Sprintf("failed to decode image: %v", err),
			"data":    fiber.Map{},
		})
	}
	fileData.Close()

	// 画像のリサイズ
	if img.Bounds().Dx() > 800 {
		osPath, err := service.ResizeImg(img, fileData, data, filename)
		if err != nil {
			return c.JSON(fiber.Map{
				"status":  false,
				"code":    "resize_error",
				"message": fmt.Sprintf("failed to resize image: %v", err),
				"data":    fiber.Map{},
			})
		}
		fileData, err = os.Open(osPath)
		if err != nil {
			log.Printf("failed to open image: %s", err)
			return c.JSON(fiber.Map{
				"status":  false,
				"code":    "file_open_error",
				"message": fmt.Sprintf("failed to open image: %v", err),
				"data":    fiber.Map{},
			})
		}
	}

	log.Println("start to upload image to GCS")

	jsonPath := config.Config.GcsKeyPath
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(jsonPath))
	if err != nil {
		log.Printf("failed to create client: %s", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "gcp_client_error",
			"message": fmt.Sprintf("failed to create client: %s", err),
			"data":    fiber.Map{},
		})
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	bucketName := config.Config.GcsBucketName
	objectPath := config.Config.GcsObjectPath

	o := client.Bucket(bucketName).Object(filename)
	o = o.If(storage.Conditions{DoesNotExist: true})
	wc := o.NewWriter(ctx)
	_, err = io.Copy(wc, fileData)
	if err != nil {
		log.Printf("io.Copy: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "gcp_copy_error",
			"message": fmt.Sprintf("io.Copy: %v", err),
			"data":    fiber.Map{},
		})
	}
	err = wc.Close()
	if err != nil {
		log.Printf("Writer.Close: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "gcp_close_error",
			"message": fmt.Sprintf("Writer.Close: %v", err),
			"data":    fiber.Map{},
		})
	}

	err = os.Remove("images/uploads/" + filename)
	if err != nil {
		log.Printf("failed to remove uploads/%v: %v", filename, err)
		return c.JSON(fiber.Map{
			"status":  false,
			"code":    "gcp_remove_error",
			"message": fmt.Sprintf("failed to remove images/uploads/%v: %v", filename, err),
			"data":    fiber.Map{},
		})
	}

	log.Println("Success to upload image")
	return c.JSON(fiber.Map{
		"status":   true,
		"code":     "success_upload_image",
		"url":      objectPath + filename,
		"filename": filename,
	})
}

/*
 * ImageDelete
 * imageをgcsから削除
 */
func ImageDelete(filename string) error {
	log.Printf("start to delete image form GCS: %v", filename)

	jsonPath := config.Config.GcsKeyPath
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(jsonPath))
	if err != nil {
		log.Printf("failed to create client: %s", err)
		return fmt.Errorf("failed to create client: %s", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	bucketName := config.Config.GcsBucketName

	o := client.Bucket(bucketName).Object(filename)

	attrs, err := o.Attrs(ctx)
	if err != nil {
		log.Printf("object.Attrs: %v", err)
		return fmt.Errorf("object.Attrs: %v", err)
	}
	o = o.If(storage.Conditions{GenerationMatch: attrs.Generation})

	err = o.Delete(ctx)
	if err != nil {
		log.Printf("Object.Delete: %v", err)
		return fmt.Errorf("Object.Delete: %v", err)
	}

	return nil
}
