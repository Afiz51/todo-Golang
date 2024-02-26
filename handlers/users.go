package handlers

import (
	"net/http"

	"github.com/Afiz51/TodoGoRest/db"
	"github.com/Afiz51/TodoGoRest/models"
	"github.com/Afiz51/TodoGoRest/validation"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateUserRequest represents the request structure for creating a user

func CreateUserHandler(c *gin.Context) {
	var createUserRequest models.CreateUserRequest

	// Bind JSON request body to createUserRequest struct
	if err := c.BindJSON(&createUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	errors := validation.ValidateUserRequest(createUserRequest)
	// If there are validation errors, return them
	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors})
		return
	}

	// Here you would typically process the request, such as creating a user in a database

	res, err := db.Users.InsertOne(c, createUserRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to add users"})
		return
	}

	user := models.CreateUserRequest{
		ID:        res.InsertedID.(primitive.ObjectID),
		FirstName: createUserRequest.FirstName,
		LastName:  createUserRequest.LastName,
		Username:  createUserRequest.Username,
		Email:     createUserRequest.Email,
	}
	// For this example, we'll just echo back the received data
	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"data":    user,
	})
}

func GetUsers(c *gin.Context) {
	cursor, err := db.Users.Find(c, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch users"})
		return
	}

	var users []models.CreateUserRequest
	if err = cursor.All(c, &users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not fetch users"})
		return
	}

	c.JSON(http.StatusOK, users)
}
