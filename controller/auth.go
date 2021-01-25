package controller

import (
	"gin-go-clean-architecture/entity"
	"gin-go-clean-architecture/model"
	"gin-go-clean-architecture/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type AuthCont struct {
	service service.AuthService
}

func NewAuthCont(service service.AuthService) *AuthCont {
	return &AuthCont{service: service}
}

func (a AuthCont) Login(c *gin.Context) {
	var l model.LoginRequest
	err := c.BindJSON(&l)
	if err != nil {
		res := model.ApiResponse{
			Success: false,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    gin.H{},
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	token, err := a.service.Login(l.Email, l.Password)
	if err != nil {
		res := model.ApiResponse{
			Success: false,
			Message: http.StatusText(http.StatusUnauthorized),
			Data:    gin.H{},
		}
		c.JSON(http.StatusUnauthorized, res)
		return
	}

	res := model.ApiResponse{
		Success: true,
		Message: "OK",
		Data:    gin.H{"token": token},
	}
	c.JSON(http.StatusOK, res)
}

func (a AuthCont) SignUp(c *gin.Context) {
	var u entity.User
	err := c.Bind(&u)
	if err != nil {
		res := model.ApiResponse{
			Success: false,
			Message: http.StatusText(http.StatusInternalServerError),
			Data:    gin.H{},
		}
		c.JSON(http.StatusInternalServerError, res)
		return
	}

	file, _, err := c.Request.FormFile("photo")

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err.Error()})
		return
	}
	defer file.Close()

	token, err := a.service.SignUp(u, file)
	if err != nil {
		res := model.ApiResponse{
			Success: false,
			Message: err.Error(),
			Data:    gin.H{},
		}
		c.JSON(http.StatusBadRequest, res)
		return
	}
	res := model.ApiResponse{
		Success: true,
		Message: "OK",
		Data:    gin.H{"email": u.Email, "token": token},
	}
	c.JSON(http.StatusOK, res)
}
