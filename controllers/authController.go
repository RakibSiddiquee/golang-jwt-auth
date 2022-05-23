package controllers

import (
	"strconv"
	"time"

	"github.com/RakibSiddiquee/go-fiber-jwt-auth/database"
	"github.com/RakibSiddiquee/go-fiber-jwt-auth/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	database.DB.Create(&user)

	return c.JSON(user)
}

// func Register(c *fiber.Ctx) error {
// 	fmt.Println("Register")
// 	var data map[string]string
// 	user := new(models.User)
// 	fmt.Println("c", c.BodyParser(&data))
// 	if err := c.BodyParser(&user); err != nil {
// 		// return err
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}

// 	errors := validation.ValidateStruct(*user)
// 	if errors != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(errors)
// 	}

// 	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

// 	newuser := models.User{
// 		Name:     user.Name,
// 		Email:    user.Email,
// 		Password: password,
// 	}

// 	fmt.Println("dd", user, newuser)

// 	database.DB.Create(&newuser)

// 	// return c.SendString("Hello, World 👋!")
// 	return c.JSON(newuser)
// }

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// pp, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"]))

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"error": "Invalid password",
		})
	}

	// Create token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: jwt.NewNumericDate(time.Unix(time.Now().Add(time.Hour*24).Unix(), 0)), // 1 day
	})

	// time.Now().Add(time.Hour * 24).Unix()

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"error": "could not login",
		})
	}

	// Create cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	// return c.JSON(token)
	return c.JSON(fiber.Map{
		"success": "success message",
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.RegisteredClaims)

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
	// Remove cookie
	// -time.Hour = expires before one hour
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
