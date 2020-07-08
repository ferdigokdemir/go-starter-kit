package controllers

import (
	"context"
	"github.com/gofiber/fiber"
	"go-starter-kit/config/database"
	"go-starter-kit/models"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

func GetCategories(ctx *fiber.Ctx) {
	collection := database.MongoInstance().Collection("categories")
	categories := []models.Category{}
	cursor, err := collection.Aggregate(context.TODO(), []bson.M{bson.M{"$lookup": bson.M{
		"from":         "upload_file", // other table name
		"localField":   "image",       // name of users table field
		"foreignField": "_id",         // name of userinfo table field
		"as":           "image",       // alias for userinfo table
	}}})

	if err != nil {
		log.Printf("Error while getting all todos, Reason: %v\n", err)
		ctx.SendStatus(http.StatusInternalServerError)
		ctx.Error()
	}

	// Iterate through the returned cursor.
	for cursor.Next(context.TODO()) {
		var category models.Category
		cursor.Decode(&category)
		categories = append(categories, category)
	}

	ctx.JSON(categories)
}
