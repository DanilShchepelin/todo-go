package controller

import (
	"awesomeProject/database"
	"awesomeProject/model"
	"github.com/gofiber/fiber/v2"
)

// GetAllTodoItems is a function to get all to-do items data from database
// @Summary Get all to-do
// @Description Get all to-do items
// @Tags to-doItems
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=[]models.TodoItem}
// @Failure 503 {object} ResponseHTTP{}
// @Router /api/todo [get]
func GetAllTodoItems(c *fiber.Ctx) error {
	db := database.DB
	var todoItems []model.TodoItem
	db.Find(&todoItems)
	return c.JSON(fiber.Map{"status": "success", "message": "All todo items", "data": todoItems})
}

func GetTodoItemById(c *fiber.Ctx) error {
	db := database.DB
	var todoItem model.TodoItem
	db.Find(&todoItem, c.Params("id"))
	return c.JSON(fiber.Map{"status": "success", "message": "Todo item by ID", "data": todoItem})
}

func CreateTodoItem(c *fiber.Ctx) error {
	db := database.DB
	var todoItem = new(model.TodoItem)
	if err := c.BodyParser(todoItem); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create todo item", "data": err})
	}

	db.Create(&todoItem)
	return c.JSON(fiber.Map{"status": "success", "message": "Created todo item", "data": todoItem})
}

func DeleteTodoItem(c *fiber.Ctx) error {
	db := database.DB

	var todoItem model.TodoItem
	db.First(&todoItem, c.Params("id"))
	db.Delete(&todoItem)
	return c.JSON(fiber.Map{"status": "success", "message": "Todo item successfully deleted", "data": nil})
}
