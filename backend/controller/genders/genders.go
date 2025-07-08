package genders

import (
   "net/http"
   "example.com/LAB2/backend/config"
   "example.com/LAB2/backend/entity"
   "github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
   db := config.DB()
   var genders []entity.Genders
   db.Find(&genders)
   c.JSON(http.StatusOK, &genders)
}