package controllers

import (
	mongoInstance "fin-dashboard-api/app"
	"fin-dashboard-api/app/controllers/structs"
	"fin-dashboard-api/app/models"
	"fin-dashboard-api/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Login(c *gin.Context) {
	var user structs.UserBody
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data provided."})
		return
	}
	userCollection := mongoInstance.GetCollection("users")

	var existingUser models.UserModel

	err := userCollection.FindOne(c.Request.Context(), bson.M{"email": user.Email}).Decode(&existingUser)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed."})
		return
	}

	compareResult := utils.CheckPasswordHash(user.Password, existingUser.Password)

	if !compareResult {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authentication failed."})
		return
	}

	existingUser.Password = ""
	c.JSON(http.StatusOK, existingUser)

}
