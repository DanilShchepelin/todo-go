package controller

import (
	"awesomeProject/database"
	"awesomeProject/model"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
)

type ResponseHTTP struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func GetAllUsers(c *fiber.Ctx) error {
	db := database.DB
	var users []model.User
	if res := db.Find(&users); res.Error != nil {
		return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "All users",
		Data:    users,
	})
}

func GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	user := new(model.User)
	if err := db.First(&user, id).Error; err != nil {
		switch err.Error() {
		case "record not found":
			return c.Status(http.StatusNotFound).JSON(ResponseHTTP{
				Success: false,
				Message: fmt.Sprintf("User with ID %v not found.", id),
				Data:    nil,
			})
		default:
			return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			})

		}
	}

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success get user by ID.",
		Data:    *user,
	})
}

func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	var user = new(model.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	fullName := generateFullName(user)
	user.FullName = fullName

	hash, err := hashPassword(user.Password)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(ResponseHTTP{
			Success: false,
			Message: "Couldn't hash password",
			Data:    nil,
		})
	}
	user.Password = hash

	db.Create(&user)
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success create user.",
		Data:    *user,
	})
}

func DeleteUser(c *fiber.Ctx) error {
	db := database.DB
	userId := c.Params("id")

	var user model.User
	if err := db.First(&user, userId).Error; err != nil {
		switch err.Error() {
		case "record not found":
			return c.Status(http.StatusNotFound).JSON(ResponseHTTP{
				Success: false,
				Message: fmt.Sprintf("User with ID %v not found.", userId),
				Data:    nil,
			})
		default:
			return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			})

		}
	}

	db.Delete(&user)

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success delete user.",
		Data:    nil,
	})
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func generateFullName(user *model.User) string {
	fullName := strings.TrimSpace(user.LastName) + " " + strings.TrimSpace(user.Name)

	if user.Patronymic != nil {
		fullName += " " + strings.TrimSpace(*user.Patronymic)
	}

	return fullName
}
