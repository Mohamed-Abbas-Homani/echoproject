package api

import (
	"echoproject/internal/payload"
	"echoproject/pkg/client"
	pb "echoproject/proto"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserApI struct {
	UserClient pb.UserServiceClient
}

func NewUserAPI() *UserApI {
	userClient := client.ProvideUserClient()
	return &UserApI{UserClient: userClient}
}

func (u *UserApI) GetUsers(c echo.Context) error {
	getUsersResponse, err := u.UserClient.GetUsers(c.Request().Context(), &pb.GetUsersRequest{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, getUsersResponse.Users)
}

func (u *UserApI) GetUserByID(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		fmt.Println("Error:", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	getUserResponse, err := u.UserClient.GetUser(ctx, &pb.GetUserRequest{Id: id})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, getUserResponse)
}

func (u *UserApI) GetFullUser(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		fmt.Println("Error:", err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	getFullUserResponse, err := u.UserClient.GetFullUser(ctx, &pb.GetFullUserRequest{Id: id})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, getFullUserResponse)
}

func (u *UserApI) CreateUser(c echo.Context) error {
	var createUserPayload payload.CreateUserPayload
	if err := c.Bind(&createUserPayload); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	createUserResponse, err := u.UserClient.CreateUser(ctx, &pb.CreateUserRequest{
		Username:       createUserPayload.Username,
		Password:       createUserPayload.Password,
		Email:          createUserPayload.Email,
		OrganizationId: createUserPayload.OrganizationId,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, createUserResponse)
}

func (u *UserApI) UpdateUser(c echo.Context) error {
	var updateUserPayload payload.UserUpdatePayload
	if err := c.Bind(&updateUserPayload); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ctx := c.Request().Context()
	updateUserResponse, err := u.UserClient.UpdateUser(ctx, &pb.UpdateUserRequest{
		Id:             updateUserPayload.Id,
		Username:       updateUserPayload.Username,
		Email:          updateUserPayload.Email,
		OrganizationId: updateUserPayload.OrganizationId,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, updateUserResponse)
}

func (u *UserApI) DeleteUser(c echo.Context) error {
	ctx := c.Request().Context()
	var deleteUserPayload payload.UserDeletePayload
	if err := c.Bind(&deleteUserPayload); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	deleteUserResponse, err := u.UserClient.DeleteUser(ctx, &pb.DeleteUserRequest{
		Id: deleteUserPayload.Id,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, deleteUserResponse)
}
