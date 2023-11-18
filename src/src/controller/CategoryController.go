package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Category struct {
	Id   uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name string    `json:"name"`
}

func GetCategory(c *gin.Context) { //所有的類別
	category := c.Query("name")

	db := connectDB()

	if category != "" {
		var category *Category
		db.Where("name = $1", category).Find(&category)
		closeDB(db)
		c.JSON(200, category)
		return
	}
	var categories []*Category

	db.Find(&categories)

	closeDB(db)
	c.JSON(200, categories)
}

func GetCategoryById(c *gin.Context) {
	db := connectDB()
	var category *Category

	log.Println("CategoryId: " + c.Param("CategoryId"))
	queryResult := db.Where("id = $1", c.Param("CategoryId")).Take(&category)
	if queryResult.Error != nil {
		c.JSON(500, gin.H{
			"error": "query error" + queryResult.Error.Error(),
		})
		closeDB(db)
		return
	}
	closeDB(db)
	c.JSON(200, category)
}

func CreateCategory(c *gin.Context) {
	db := connectDB()
	var category Category
	c.BindJSON(&category)
	category.Id = uuid.New()

	result := db.Create(&category)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": "create error" + result.Error.Error(),
		})
		closeDB(db)
		return
	}
	closeDB(db)
	c.JSON(200, gin.H{
		"message": "create success",
	})
}

func UpdateCategoryById(c *gin.Context) {
	db := connectDB()
	var category Category

	queryResult := db.Where("id = $1", c.Param("CategoryId")).Take(&category)
	if queryResult.Error != nil {
		c.JSON(500, gin.H{
			"error": "query error" + queryResult.Error.Error(),
		})
		closeDB(db)
		return
	}

	var categorybody Category
	c.BindJSON(&categorybody)
	categorybody.Id = category.Id
	result := db.Model(&category).Where("id = ?", category.Id).Updates(categorybody)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": "update error" + result.Error.Error(),
		})
		closeDB(db)
		return
	}
	closeDB(db)
	c.JSON(200, gin.H{
		"message": "update success",
	})
}

func DeleteCategoryById(c *gin.Context) {
	db := connectDB()
	var category Category

	queryResult := db.Where("id = $1", c.Param("CategoryId")).Take(&category)
	if queryResult.Error != nil {
		c.JSON(500, gin.H{
			"error": "query error" + queryResult.Error.Error(),
		})
		closeDB(db)
		return
	}
	result := db.Delete(&category)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": "delete error" + result.Error.Error(),
		})
		closeDB(db)
		return
	}
	closeDB(db)
	c.JSON(200, gin.H{
		"message": "delete success",
	})
}
