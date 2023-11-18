package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type Item struct {
	Id         uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Order_id   uuid.UUID `gorm:"foreignKey:Order_id" json:"order_id"`
	Product_id uuid.UUID `gorm:"foreignKey:Product_id" json:"product_id"`
	Is_ship    bool      `json:"is_ship"`
}

type ItemwithAll struct {
	Item
	Order   *Order   `json:"order,omitempty"`
	Product *Product `json:"product,omitempty"`
}

func (ItemwithAll) TableName() string {
	return "items"
}

type ItemwithOrder struct {
	Item
	Order *Order `json:"order,omitempty"`
}

func (ItemwithOrder) TableName() string {
	return "items"
}

type ItemwithProduct struct {
	Item
	Product *Product `json:"product,omitempty"`
}

func (ItemwithProduct) TableName() string {
	return "items"
}

func GetItem(c *gin.Context) {

	item_id := c.Query("item_id")
	order_id := c.Query("order_id")
	product_id := c.Query("product_id")
	is_ship := c.Query("is_ship")

	db := connectDB()
	var items []*Item
	if item_id != "" {
		log.Println("item_id: " + item_id)
		db.Preload(clause.Associations).Where("id = $1", item_id).Find(&items)
	} else if order_id != "" {
		log.Println("order_id: " + order_id)
		db.Preload(clause.Associations).Where("order_id = $1", order_id).Find(&items)
	} else if product_id != "" {
		log.Println("product_id: " + product_id)
		db.Preload(clause.Associations).Where("product_id = $1", product_id).Find(&items)
	} else if is_ship != "" {
		log.Println("is_ship: " + is_ship)
		db.Preload(clause.Associations).Where("is_ship = $1", is_ship).Find(&items)
	} else {
		db.Preload(clause.Associations).Find(&items)
	}
	closeDB(db)
	c.JSON(200, items)
}

func GetItemById(c *gin.Context) {
	id := c.Param("ItemId")

	db := connectDB()
	var item Item
	queryResult := db.Preload(clause.Associations).Where("id = $1", id).Find(&item)
	if queryResult.Error != nil {
		c.JSON(500, gin.H{
			"message": "query error" + queryResult.Error.Error(),
		})
		closeDB(db)
		return
	}
	closeDB(db)
	c.JSON(200, item)
}

func CreateItem(c *gin.Context) {
	db := connectDB()
	var item Item
	c.BindJSON(&item)
	item.Id = uuid.New()

	result := db.Create(&item)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "create error" + result.Error.Error(),
		})
		closeDB(db)
		return
	}
	closeDB(db)
	c.JSON(200, gin.H{
		"message": "create success",
	})
}

func UpdateItemById(c *gin.Context) {
	db := connectDB()
	var item Item
	queryResult := db.Where("id = $1", c.Param("ItemId")).Take(&item)
	if queryResult.Error != nil {
		c.JSON(500, gin.H{
			"message": "query error" + queryResult.Error.Error(),
		})
		closeDB(db)
		return
	}

	var itembody Item
	c.BindJSON(&itembody)
	itembody.Id = item.Id
	result := db.Model(&item).Where("id = ?", item.Id).Updates(itembody)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "update error" + result.Error.Error(),
		})
		closeDB(db)
		return
	}
	closeDB(db)
	c.JSON(200, gin.H{
		"message": "update success",
	})
}

func DeleteItemById(c *gin.Context) {
	db := connectDB()
	var item Item
	queryResult := db.Where("id = $1", c.Param("ItemId")).Take(&item)
	if queryResult.Error != nil {
		c.JSON(500, gin.H{
			"message": "query error" + queryResult.Error.Error(),
		})
		closeDB(db)
		return
	}

	result := db.Delete(&item)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": "delete error" + result.Error.Error(),
		})
		closeDB(db)
		return
	}
	closeDB(db)
	c.JSON(200, gin.H{
		"message": "delete success",
	})
}

func GetOrderByItemId(c *gin.Context) {
	db := connectDB()
	var item ItemwithOrder
	queryResult := db.Preload(clause.Associations).Where("id = $1", c.Param("ItemId")).Take(&item)
	if queryResult.Error != nil {
		c.JSON(500, gin.H{
			"message": "query error" + queryResult.Error.Error(),
		})
		closeDB(db)
		return
	}
	closeDB(db)
	c.JSON(200, item.Order)
}

func GetProductByItemId(c *gin.Context) {
	db := connectDB()
	var item ItemwithProduct
	queryResult := db.Preload(clause.Associations).Where("id = $1", c.Param("ItemId")).Take(&item)
	if queryResult.Error != nil {
		c.JSON(500, gin.H{
			"message": "query error" + queryResult.Error.Error(),
		})
		closeDB(db)
		return
	}
	closeDB(db)
	c.JSON(200, item.Product)
}
