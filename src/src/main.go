package main

import (
	"go_gin_example/controller"
	"go_gin_example/envconfig"
	"log"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

func GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"users": envconfig.GetEnv("USERS"),
	})
}
func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Li, Tzu Hao is nice guy!",
	})
}

func main() {
	server := gin.Default()

	server.GET("/", GetIndex)      // 讀取首頁
	server.GET("/users", GetUsers) // 讀取Users

	//GET /products
	server.GET("/products", controller.GetProduct) // 讀取Products
	server.GET("/products/:ProductId", controller.GetProductById)
	server.GET("/products/:ProductId/category", controller.GetCategoryByProductId)
	//POST /products
	server.POST("/products", controller.CreateProduct) // 新增Products
	//PUT /products
	server.PUT("/products/:ProductId", controller.UpdateProductById) // 更新Products
	//DELETE /products
	server.DELETE("/products/:ProductId", controller.DeleteProductById) // 刪除Products

	//GET /customers
	server.GET("/customers", controller.GetCustomer) // 讀取Customers
	server.GET("/customers/:CustomerId", controller.GetCustomerById)
	//POST /customers
	server.POST("/customers", controller.CreateCustomer) // 新增Customers
	//PUT /customers
	server.PUT("/customers/:CustomerId", controller.UpdateCustomerById) // 更新Customers
	//DELETE /customers
	server.DELETE("/customers/:CustomerId", controller.DeleteCustomerById) // 刪除Customers

	//GET /orders
	server.GET("/orders", controller.GetOrder) // 讀取Orders
	server.GET("/orders/:OrderId", controller.GetOrderById)
	server.GET("/orders/:OrderId/customer", controller.GetCustomerByOrderId)
	//POST /orders
	server.POST("/orders", controller.CreateOrder) // 新增Orders
	//PUT /orders
	server.PUT("/orders/:OrderId", controller.UpdateOrderById) // 更新Orders
	//DELETE /orders
	server.DELETE("/orders/:OrderId", controller.DeleteOrderById) // 刪除Orders

	//GET /items
	server.GET("/items", controller.GetItem) // 讀取Items
	server.GET("/items/:ItemId", controller.GetItemById)
	server.GET("/items/:ItemId/order", controller.GetOrderByItemId)
	server.GET("/items/:ItemId/product", controller.GetProductByItemId)
	//POST /items
	server.POST("/items", controller.CreateItem) // 新增Items
	//PUT /items
	server.PUT("/items/:ItemId", controller.UpdateItemById) // 更新Items
	//DELETE /items
	server.DELETE("/items/:ItemId", controller.DeleteItemById) // 刪除Items

	//GET /category
	server.GET("/category", controller.GetCategory) // 讀取Category
	server.GET("/category/:CategoryId", controller.GetCategoryById)
	//POST /category
	server.POST("/category", controller.CreateCategory) // 新增Category
	//PUT /category
	server.PUT("/category/:CategoryId", controller.UpdateCategoryById) // 更新Category
	//DELETE /category
	server.DELETE("/category/:CategoryId", controller.DeleteCategoryById) // 刪除Category

	if err := server.Run(":" + envconfig.GetEnv("PORT")); err != nil {
		log.Fatalln(err.Error())
		return
	}
}
