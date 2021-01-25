package main

import (
	"github.com/gin-gonic/gin"
	"solid-go/controller"
	"solid-go/db"
	"solid-go/middleware"
	"solid-go/repository"
	"solid-go/service"
	"solid-go/storage"
)

var (
	//config
	mongo       = db.InitMongo()
	jwt         = service.NewJWTService()
	fileStorage = storage.NewMinIO()
	//setup repository
	userRepo = repository.NewUserRepo(mongo)
	//setup service
	userServ = service.NewUserService(userRepo)
	authServ = service.NewAuthService(userRepo, jwt, fileStorage)
	//setup controller
	userCont = controller.NewUserCont(userServ)
	authCont = controller.NewAuthCont(authServ)
)

func main() {
	r := gin.Default()

	r.POST("/login", authCont.Login)
	r.POST("/signup", authCont.SignUp)

	//private
	r.Use(middleware.AuthorizeJWT())
	r.GET("/user", userCont.GetUser)
	r.GET("/user/:guid", userCont.GetUserByID)

	r.Run(":9999")
}
