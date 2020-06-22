package user

import (
	"github.com/gofiber/fiber"
	"github.com/harrizsb/simple-crud/database"
	"github.com/harrizsb/simple-crud/helpers"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func CreateUser(c *fiber.Ctx) {
	u := new(User)

	if err := c.BodyParser(u); err != nil {
		c.JSON(Response{Success: false, Msg: err.Error()})
	}

	db := database.DBConn
	user := User{UserID: u.UserID, Email: u.Email, Address: u.Address, Password: u.Password}
	if err := db.Create(&user).Error; err != nil {
		c.JSON(Response{Success: false, Msg: err.Error()})
	} else {
		c.JSON(DatabaseResponse{Success: true, Data: user})
	}
}

func UpdateUser(c *fiber.Ctx) {
	u := new(User)

	if err := c.BodyParser(u); err != nil {
		c.JSON(Response{Success: false, Msg: err.Error()})
	}

	db := database.DBConn
	user := User{UserID: u.UserID, Email: u.Email, Address: u.Address, Password: u.Password}
	// exclude userID
	if err := db.Model(User{}).Omit("UserID").Updates(user).Error; err != nil {
		c.JSON(Response{Success: false, Msg: err.Error()})
	} else {
		c.JSON(DatabaseResponse{Success: true, Data: user})
	}
}

func Login(c *fiber.Ctx) {
	var err error
	u := new(User)

	if err = c.BodyParser(u); err != nil {
		c.JSON(Response{Success: false, Msg: err.Error()})
	}

	var user User
	db := database.DBConn

	if err = db.Where("user_id = ? AND password = ?", u.UserID, u.Password).
		First(&user).Error; err != nil {
		c.JSON(Response{Success: false, Msg: err.Error()})
	} else {
		jwt, err := helpers.CreateJWT(user.ID)

		if err != nil {
			c.JSON(Response{Success: false, Msg: err.Error()})
		} else {
			c.JSON(LoginResponse{Success: true, JWT: jwt})
		}
	}
}

func GetUser(c *fiber.Ctx) {
	var user User
	username := c.Params("username")
	db := database.DBConn

	if err := db.First(&user, &User{UserID: username}).Error; err != nil {
		c.JSON(Response{Success: false, Msg: err.Error()})
	} else {
		c.JSON(DatabaseResponse{Success: true, Data: user})
	}
}

func DeleteUser(c *fiber.Ctx) {
	username := c.Params("username")
	user := User{UserID: username}
	db := database.DBConn

	// permanent delete
	if err := db.Unscoped().Delete(&user).Error; err != nil {
		c.JSON(Response{Success: false, Msg: err.Error()})
	} else {
		c.JSON(DeleteUserResponse{Success: true, Msg: "User deleted successfully"})
	}
}
