package main

import (
	"gin-go-clean-architecture/controller"
	"gin-go-clean-architecture/db"
	"gin-go-clean-architecture/middleware"
	"gin-go-clean-architecture/repository"
	"gin-go-clean-architecture/service"
	"gin-go-clean-architecture/storage"
	"github.com/gin-gonic/gin"
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
