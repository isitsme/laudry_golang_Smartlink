package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"laundry-api/config"
	"laundry-api/models"
	"laundry-api/utils"
)

type ServiceInput struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
}

func CreateService(c *gin.Context) {
	var input ServiceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}
	svc := models.Service{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
	}
	if err := config.DB.Create(&svc).Error; err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to create service")
		return
	}
	utils.RespondSuccess(c, http.StatusCreated, svc)
}

func ListServices(c *gin.Context) {
	var services []models.Service
	config.DB.Find(&services)
	utils.RespondSuccess(c, http.StatusOK, services)
}

func GetService(c *gin.Context) {
	id := c.Param("id")
	var svc models.Service
	if err := config.DB.First(&svc, id).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "Service not found")
		return
	}
	utils.RespondSuccess(c, http.StatusOK, svc)
}

func UpdateService(c *gin.Context) {
	id := c.Param("id")
	var svc models.Service
	if err := config.DB.First(&svc, id).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "Service not found")
		return
	}
	var input ServiceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}
	svc.Name = input.Name
	svc.Description = input.Description
	svc.Price = input.Price
	config.DB.Save(&svc)
	utils.RespondSuccess(c, http.StatusOK, svc)
}

func DeleteService(c *gin.Context) {
	id := c.Param("id")
	var svc models.Service
	if err := config.DB.First(&svc, id).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "Service not found")
		return
	}
	config.DB.Delete(&svc)
	utils.RespondSuccess(c, http.StatusOK, gin.H{"message": "Service deleted"})
}
