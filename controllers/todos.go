package controllers

import (
	"go-starter-kit/config/database"
	"go-starter-kit/models"
	"strconv"

	"github.com/gofiber/fiber"
)

// GetTodos is getting todo list
func GetTodos(ctx *fiber.Ctx) {
	limit := 10
	start := 0

	if ctx.Query("limit") != "" {
		limit, _ = strconv.Atoi(ctx.Query("limit"))
		if limit > 100 {
			limit = 10
		}
	}

	if ctx.Query("start") != "" {
		start, _ = strconv.Atoi(ctx.Query("start"))
	}

	var todos []*models.Todo
	database.DB().Offset(start).Limit(limit).Find(&todos)
	ctx.JSON(todos)
}
