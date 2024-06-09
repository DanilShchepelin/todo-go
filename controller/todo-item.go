package controller

import (
	"awesomeProject/database"
	"awesomeProject/model"
	"github.com/gofiber/fiber/v2"
)

// GetAllTodoItems Получить список элементов
//	@Summary		Получить список элементов
//	@Description	Получить список элементов
//	@Tags			TodoItems
//	@Accept			json
//	@Produce		json
//	@Router			/api/todo [get]
func GetAllTodoItems(c *fiber.Ctx) error {
	db := database.DB
	var todoItems []model.TodoItem
	db.Find(&todoItems)
	return c.JSON(fiber.Map{"status": "success", "message": "All todo items", "data": todoItems})
}

// GetTodoItemById Получить элемент по id
//	@Summary		Получить элемент по id
//	@Description	Получить элемент по id
//	@Tags			TodoItems
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"TodoItem ID"
//	@Router			/api/todo/{id} [get]
func GetTodoItemById(c *fiber.Ctx) error {
	db := database.DB
	var todoItem model.TodoItem
	db.Find(&todoItem, c.Params("id"))
	return c.JSON(fiber.Map{"status": "success", "message": "Todo item by ID", "data": todoItem})
}

// CreateTodoItem Создать элемент
//	@Summary		Создать элемент
//	@Description	Создать элемент
//	@Tags			TodoItems
//	@Accept			json
//	@Produce		json
//	@Param			request	body model.TodoItem	true	"Данные для создания нового элемента"
//	@Router			/api/todo [post]
func CreateTodoItem(c *fiber.Ctx) error {
	db := database.DB
	var todoItem = new(model.TodoItem)
	if err := c.BodyParser(todoItem); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create todo item", "data": err})
	}

	db.Create(&todoItem)
	return c.JSON(fiber.Map{"status": "success", "message": "Created todo item", "data": todoItem})
}

// DeleteTodoItem Удалить элемент
//	@Summary		Удалить элемент
//	@Description	Удалить элемент
//	@Tags			TodoItems
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"TodoItem ID"
//	@Router			/api/todo/{id} [delete]
func DeleteTodoItem(c *fiber.Ctx) error {
	db := database.DB

	var todoItem model.TodoItem
	db.First(&todoItem, c.Params("id"))
	db.Delete(&todoItem)
	return c.JSON(fiber.Map{"status": "success", "message": "Todo item successfully deleted", "data": nil})
}
