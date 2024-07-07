package routes

import (
	"echoproject/internal/api"
	"github.com/labstack/echo/v4"
	"os"
)

var router *echo.Echo

func InitRoutes() {
	router = echo.New()
	userAPI := api.NewUserAPI()
	userRoutes := router.Group("/user")
	userRoutes.GET("", userAPI.GetUsers)
	userRoutes.GET("/:id", userAPI.GetUserByID)
	userRoutes.GET("/full/:id", userAPI.GetFullUser)
	userRoutes.POST("", userAPI.CreateUser)
	userRoutes.PUT("/:id", userAPI.UpdateUser)
	userRoutes.DELETE("/:id", userAPI.DeleteUser)
}

func Run() {
	router.Logger.Fatal(router.Start(":" + os.Getenv("PORT")))
}
