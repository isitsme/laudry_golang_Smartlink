package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"laundry-api/config"
	"laundry-api/models"
	"laundry-api/utils"
)

type OrderInput struct {
	CustomerName string    `json:"customer_name" binding:"required"`
	CustomerID   uint      `json:"customer_id"`
	ServiceID    uint      `json:"service_id" binding:"required"`
	Quantity     int       `json:"quantity" binding:"required,gt=0"`
	DueDate      time.Time `json:"due_date"`
}

func CreateOrder(c *gin.Context) {
	var input OrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}
	var svc models.Service
	if err := config.DB.First(&svc, input.ServiceID).Error; err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Service not found")
		return
	}
	total := float64(input.Quantity) * svc.Price
	order := models.Order{
		CustomerName: input.CustomerName,
		CustomerID:   input.CustomerID,
		ServiceID:    input.ServiceID,
		Quantity:     input.Quantity,
		TotalPrice:   total,
		Status:       "pending",
		DueDate:      input.DueDate,
	}
	if err := config.DB.Create(&order).Error; err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to create order")
		return
	}
	config.DB.Preload("Service").First(&order, order.ID)
	utils.RespondSuccess(c, http.StatusCreated, order)
}

func ListOrders(c *gin.Context) {
	var orders []models.Order
	config.DB.Preload("Service").Find(&orders)
	utils.RespondSuccess(c, http.StatusOK, orders)
}

func GetOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := config.DB.Preload("Service").First(&order, id).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "Order not found")
		return
	}
	utils.RespondSuccess(c, http.StatusOK, order)
}

type UpdateOrderInput struct {
	Quantity *int       `json:"quantity"`
	Status   *string    `json:"status"`
	DueDate  *time.Time `json:"due_date"`
}

func UpdateOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := config.DB.First(&order, id).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "Order not found")
		return
	}
	var input UpdateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}
	if input.Quantity != nil {
		order.Quantity = *input.Quantity
		var svc models.Service
		config.DB.First(&svc, order.ServiceID)
		order.TotalPrice = float64(order.Quantity) * svc.Price
	}
	if input.Status != nil {
		order.Status = *input.Status
	}
	if input.DueDate != nil {
		order.DueDate = *input.DueDate
	}
	config.DB.Save(&order)
	config.DB.Preload("Service").First(&order, order.ID)
	utils.RespondSuccess(c, http.StatusOK, order)
}

func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := config.DB.First(&order, id).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "Order not found")
		return
	}
	config.DB.Delete(&order)
	utils.RespondSuccess(c, http.StatusOK, gin.H{"message": "Order deleted"})
}
