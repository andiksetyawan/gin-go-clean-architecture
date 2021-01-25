package controller

import (
	"gin-go-clean-architecture/model"
	"gin-go-clean-architecture/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserCont struct {
	userService service.UserService
	//postService service.PostService
}

func NewUserCont(service service.UserService) *UserCont {
	return &UserCont{userService: service}
}

func (u *UserCont) GetUser(c *gin.Context) {
	users, err := u.userService.FindUser()
	if err != nil {
		res := model.ApiResponse{Message: http.StatusText(http.StatusInternalServerError), Data: gin.H{}}
		c.JSON(http.StatusInternalServerError, &res)
	}
	res := model.ApiResponse{Message: "OK", Success: true, Data: users}
	c.JSON(200, &res)
}

func (u *UserCont) GetUserByID(c *gin.Context) {
	guid, _ := c.Params.Get("guid")
	user, err := u.userService.FindUserByID(guid)
	if err != nil {
		res := model.ApiResponse{Message: http.StatusText(http.StatusInternalServerError), Data: gin.H{}}
		c.JSON(http.StatusInternalServerError, &res)
		return
	}
	res := model.ApiResponse{Message: "OK", Success: true, Data: user}
	c.JSON(200, &res)
}
