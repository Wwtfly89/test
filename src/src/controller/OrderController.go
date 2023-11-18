package controller

import (
	//"go_gin_example/model"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type Order struct {
	Id          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`
	Customer_id uuid.UUID `gorm:"foreignKey:Customer_id" json:"customer_id"`
	Is_paid     bool      `json:"is_paid"`
}

type OrderwithAll struct {
	Order
	Customer *Customer `json:"customer,omitempty"`
}

func (OrderwithAll) TableName() string {
	return "orders"
}

func GetOrder(c *gin.Context) {
	order_id := c.Query("order_id")
	customer_id := c.Query("customer_id")
	is_paid := c.Query("is_paid")
	customer_name := c.Query("customer_name")
	db := connectDB()
	var orders []*OrderwithAll
	if order_id != "" {
		log.Println("order_id: " + order_id)
		db.Preload(clause.Associations).Where("id = $1", order_id).Find(&orders)
	} else if customer_id != "" {
		log.Println("customer_id: " + customer_id)
		db.Preload(clause.Associations).Where("customer_id = $1", customer_id).Find(&orders)
	} else if is_paid != "" {
		log.Println("is_paid: " + is_paid)
		db.Preload(clause.Associations).Where("is_paid = $1", is_paid).Find(&orders)
	} else if customer_name != "" {
		log.Println("customer_name: " + customer_name)
		var customer_id uuid.UUID
		db.Table("customers").Select("id").Where("name = $1", customer_name).Find(customer_id)
		db.Preload(clause.Associations).Where("customer_id = $1", customer_id).Find(&orders)
	} else {
		db.Preload(clause.Associations).Find(&orders)
	}
	closeDB(db)
	c.JSON(200, orders)
}

func GetOrderById(c *gin.Context) {
	id := c.Param("OrderId")

	db := connectDB()
	var order Order
	queryresult := db.Preload(clause.Associations).Where("id = $1", id).Take(&order)
	if queryresult.Error != nil {
		c.JSON(500, gin.H{
			"message": "query error" + queryresult.Error.Error(),
		})
		closeDB(db)
		return
	}
	closeDB(db)
	c.JSON(200, order)
}

func CreateOrder(c *gin.Context) {
	db := connectDB()
	var order Order
	c.BindJSON(&order)
	order.Id = uuid.New()
	result := db.Create(&order)
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

func UpdateOrderById(c *gin.Context) {
	db := connectDB()
	var order Order

	queryResult := db.Where("id = $1", c.Param("OrderId")).Take(&order)
	if queryResult.Error != nil {
		c.JSON(500, gin.H{
			"message": "query error" + queryResult.Error.Error(),
		})
		closeDB(db)
		return
	}

	var orderbody Order
	c.BindJSON(&orderbody)
	orderbody.Id = order.Id

	result := db.Model(&order).Where("id = ?", order.Id).Updates(orderbody)
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

func DeleteOrderById(c *gin.Context) {
	db := connectDB()
	var order Order

	queryResult := db.Where("id = $1", c.Param("OrderId")).Take(&order)
	if queryResult.Error != nil {
		c.JSON(500, gin.H{
			"message": "query error" + queryResult.Error.Error(),
		})
		closeDB(db)
		return
	}

	result := db.Delete(&order)
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

func GetCustomerByOrderId(c *gin.Context) {
	db := connectDB()
	var order OrderwithAll
	queryresult := db.Preload(clause.Associations).Where("id = $1", c.Param("OrderId")).Take(&order)
	if queryresult.Error != nil {
		c.JSON(500, gin.H{
			"message": "query error" + queryresult.Error.Error(),
		})
		closeDB(db)
		return
	}
	closeDB(db)
	c.JSON(200, order.Customer)
}
