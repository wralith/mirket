package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	pb "github.com/wralith/mirket/pb/user"
	"github.com/wralith/mirket/src/api-gateway/pkg/bcrypt"
	"github.com/wralith/mirket/src/api-gateway/pkg/protojson"
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

	pbRes, err := h.UserService.GetUser(context.Background(), &pb.GetUserRequest{Id: uint32(id)})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	user, err := protojson.Marshal(pbRes.GetUser())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSONBlob(http.StatusOK, user)

}

func (h *userHandler) AddUser(c echo.Context) error {
	var err error

	u := new(pb.User)
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	u.Password, err = bcrypt.HashPassword(u.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	in := &pb.AddUserRequest{User: u}

	pbRes, err := h.UserService.AddUser(context.Background(), in)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	res, err := protojson.Marshal(pbRes.GetUser())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSONBlob(http.StatusCreated, res)
}
