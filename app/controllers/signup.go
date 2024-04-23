package controllers

import (
	mongoInstance "fin-dashboard-api/app"
	"fin-dashboard-api/app/controllers/structs"
	"fin-dashboard-api/app/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var newUser structs.UserBody
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid data provided."})
		return
	}

	userCollection := mongoInstance.GetCollection("users")

	var existingUser models.UserModel
	err := userCollection.FindOne(c.Request.Context(), bson.M{"email": newUser.Email}).Decode(&existingUser)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
			if hashErr != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not hash password."})
				return
			}

			_, insertErr := userCollection.InsertOne(c.Request.Context(), bson.M{
				"email":     newUser.Email,
				"password":  string(hashedPassword),
				"createdAt": time.Now(),
				"updatedAt": time.Now(),
			})
			if insertErr != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create user."})
				return
			}

			c.JSON(http.StatusCreated, gin.H{"message": "Your account has been created!"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error checking for existing user."})
			return
		}
	}
	c.JSON(http.StatusConflict, gin.H{"message": "User with such Email already exists."})
}
