package controller

import (
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/xdorro/golang-fiber-base-project/app/dto"
	"github.com/xdorro/golang-fiber-base-project/app/repository"
	"github.com/xdorro/golang-fiber-base-project/pkg/config"
	"github.com/xdorro/golang-fiber-base-project/pkg/util"
	"time"
)

// AuthToken : Find user by Username and Password and Status = 1
func AuthToken(c *fiber.Ctx) error {
	var userRequest dto.UserRequest

	if err := c.BodyParser(&userRequest); err != nil {
		return util.ResponseBadRequest(c, "Đăng nhập không thành công", err)
	}

	user, err := repository.FindUserByUsernameAndStatus(userRequest.Username, 1)
	if user == nil || user.Username == "" || err != nil {
		return util.ResponseUnauthenticated(c, "Tài khoản không tồn tại", err)
	}

	if !util.CheckPasswordHash(userRequest.Password, user.Password) {
		return util.ResponseUnauthenticated(c, "Mật khẩu không chính xác", nil)
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.UserId
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	result, err := token.SignedString([]byte(config.GetJwt().Secret))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return util.ResponseSuccess(c, "Thành công", result)
}