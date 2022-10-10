package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	pb "github.com/wralith/mirket/pb/user"
)

type userHandler struct {
	UserService pb.UserServiceClient
}

func NewUserHandler(userSrv pb.UserServiceClient) *userHandler {
	return &userHandler{
		UserService: userSrv,
	}
}

func (h *userHandler) GetUser(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "invalid id")
	}

	user, err := h.UserService.GetUser(context.Background(), &pb.GetUserRequest{Id: uint32(id)})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, user.GetUser())

}
