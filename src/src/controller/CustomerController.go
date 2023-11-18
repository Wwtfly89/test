package controller

import (
	//"go_gin_example/model"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Customer struct {
	Id   uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Name string    `json:"name"`
}

func GetCustomer(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")
	db := connectDB()
	var customers []*Customer
	if name != "" {
		log.Println("name: " + name)
		db.Where("name = $1", name).Find(&customers)
	} else if id != "" {
		log.Println("id: " + id)
		db.Where("id = $1", id).Find(&customers)
	} else {
		db.Find(&customers)
	}
	closeDB(db)
	c.JSON(200, customers)
}

func GetCustomerById(c *gin.Context) {
	db := connectDB()
	var customer Customer
	queryResult := db.Where("id = $1", c.Param("CustomerId")).Take(&customer)
	if queryResult.Error != nil {
		c.JSON(500, gin.H{
			"message": "query error" + queryResult.Error.Error(),
		})
		closeDB(db)
		return
	}
	closeDB(db)
	c.JSON(200, customer)
}

func CreateCustomer(c *gin.Context) {
	db := connectDB()
	var customer Customer
	c.BindJSON(&customer)
	customer.Id = uuid.New()

	result := db.Create(&customer)
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

func UpdateCustomerById(c *gin.Context) {
	db := connectDB()
	var customer Customer
	queryResult := db.Where("id = $1", c.Param("CustomerId")).Take(&customer)
	if queryResult.Error != nil {
		c.JSON(500, gin.H{
			"message": "query error" + queryResult.Error.Error(),
		})
		closeDB(db)
		return
	}

	var customerbody Customer
	c.BindJSON(&customerbody)
	customerbody.Id = customer.Id

	result := db.Model(&customer).Where("id = ?", customer.Id).Updates(customerbody)
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

func DeleteCustomerById(c *gin.Context) {
	db := connectDB()
	var customer Customer

	queryResult := db.Where("id = $1", c.Param("CustomerId")).Take(&customer)
	if queryResult.Error != nil {
		c.JSON(500, gin.H{
			"message": "query error" + queryResult.Error.Error(),
		})
		closeDB(db)
		return
	}

	result := db.Delete(&customer)
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
