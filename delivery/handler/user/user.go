package user

import (
	"event-planner/delivery/middleware"
	"event-planner/delivery/response"
	"event-planner/entity"
	"event-planner/service/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	userUseCase user.UserUseCaseInterface
}

func NewUserHandler(userUseCase user.UserUseCaseInterface) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (uh *UserHandler) GetAllHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		users, err := uh.userUseCase.GetAll()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("success get all users", users))
	}
}

func (uh *UserHandler) GetUserById() echo.HandlerFunc {
	return func(c echo.Context) error {

		var id = middleware.ExtractToken(c)

		users, err := uh.userUseCase.GetUserById(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("success get user by id", users))
	}
}

func (uh *UserHandler) GetUserByIdUser() echo.HandlerFunc {
	return func(c echo.Context) error {

		var id = middleware.ExtractToken(c)

		users, err := uh.userUseCase.GetUserById(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccess("success get user by id", users))
	}
}

func (uh *UserHandler) CreateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var user entity.User
		c.Bind(&user)
		err := uh.userUseCase.CreateUser(user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed to create user"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccessWithoutData("success Create users"))
	}
}

func (uh *UserHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id = middleware.ExtractToken(c)

		err := uh.userUseCase.DeleteUser(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccessWithoutData("success delete user by id"))
	}
}

func (uh *UserHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var id = middleware.ExtractToken(c)
		var users entity.User
		c.Bind(&users)

		err := uh.userUseCase.UpdateUser(id, users)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, response.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, response.ResponseSuccessWithoutData("success update user by id"))
	}
}
