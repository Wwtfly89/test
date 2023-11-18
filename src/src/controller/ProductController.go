package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type Product struct {
	Id          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name        string    `json:"name"`
	Price       string    `json:"price"`
	Category_id uuid.UUID ` gorm:"foreignKey:Category_id" json:"category_id"`
}

type ProductwithAll struct {
	Product
	Category *Category `json:"category,omitempty"`
}

func (ProductwithAll) TableName() string {
	return "products"
}

func GetProduct(c *gin.Context) {
	product_id := c.Query("id")
	product_name := c.Query("name")
	product_price := c.Query("price")
	product_category_id := c.Query("category_id")

	db := connectDB()
	//
	var products []*ProductwithAll
	// 如果 product_name 不是空字串，就用 product_name 去找
	if product_id != "" {
		log.Print("product_id: " + product_id)
		db.Preload(clause.Associations).Where("id = $1", product_id).Find(&products)
	} else if product_name != "" {
		log.Print("product_name: " + product_name)
		db.Preload(clause.Associations).Where("name = $1", product_name).Find(&products)
	} else if product_price != "" {
		log.Print("product_price: " + product_price)
		db.Preload(clause.Associations).Where("price = $1", product_price).Find(&products)
	} else if product_category_id != "" {
		log.Print("product_category_id: " + product_category_id)
		db.Preload(clause.Associations).Where("category_id = $1", product_category_id).Find(&products)
	} else {
		db.Preload(clause.Associations).Find(&products)
		// db.Find(&products)
	}

	closeDB(db)

	c.JSON(200, products)
}

func GetProductById(c *gin.Context) {
	db := connectDB()
	var product *Product

	queryResult := db.Where("id = $1", c.Param("ProductId")).Take(&product)
	if queryResult.Error != nil {
		c.JSON(500, gin.H{
			"message": "Get product failed with error: " + queryResult.Error.Error(),
		})
		closeDB(db)
		return
	}

	closeDB(db)
	c.JSON(200, product)
}

func CreateProduct(c *gin.Context) {
	db := connectDB()
	var product Product
	c.BindJSON(&product)

	product.Id = uuid.New()

	result := db.Create(&product)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Create product failed with error: " + result.Error.Error(),
		})
		closeDB(db)
		return
	}

	closeDB(db)
	c.JSON(200, gin.H{
		"message": "Create product successfully!",
	})
}

func UpdateProductById(c *gin.Context) {
	db := connectDB()
	var product Product
	queryResult := db.Where("id = $1", c.Param("ProductId")).Take(&product)
	if queryResult.Error != nil {
		c.JSON(500, gin.H{
			"message": "Update product failed with error: " + queryResult.Error.Error(),
		})
		closeDB(db)
		return
	}

	var productbody Product

	c.BindJSON(&productbody)
	productbody.Id = product.Id

	result := db.Model(&product).Where("id = ?", product.Id).Updates(productbody)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Update product failed with error: " + result.Error.Error(),
		})
		closeDB(db)
		return
	}

	closeDB(db)
	c.JSON(200, gin.H{
		"message": "Update product successfully!",
	})
}

func DeleteProductById(c *gin.Context) {
	db := connectDB()
	var product Product

	queryResult := db.Where("id = $1", c.Param("ProductId")).Take(&product)
	if queryResult.Error != nil {
		c.JSON(500, gin.H{
			"message": "Delete product failed with error: " + queryResult.Error.Error(),
		})
		closeDB(db)
		return
	}

	result := db.Delete(&product)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "Delete product failed with error: " + result.Error.Error(),
		})
		closeDB(db)
		return
	}

	closeDB(db)
	c.JSON(200, gin.H{
		"message": "Delete product successfully!",
	})
}

func GetCategoryByProductId(c *gin.Context) {
	db := connectDB()
	var product *ProductwithAll

	queryResult := db.Preload(clause.Associations).Where("id = $1", c.Param("ProductId")).Take(&product)
	if queryResult.Error != nil {
		c.JSON(500, gin.H{
			"message": "Get category failed with error: " + queryResult.Error.Error(),
		})
		closeDB(db)
		return
	}

	closeDB(db)
	c.JSON(200, product.Category)
}
