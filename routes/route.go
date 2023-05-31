package routes

import (
	"log"
	"my-me/controllers"
	"my-me/middlewares"
	"my-me/repositories"
	"my-me/usecases"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func init() {
	middleware.ErrJWTMissing.Code = 401
	middleware.ErrJWTMissing.Message = "Unauthorized"
}

func Init(e *echo.Echo, db *gorm.DB) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	userRepository := repositories.NewUserRepository(db)
	userUsecase := usecases.NewUserUsecase(userRepository)
	userController := controllers.NewUserController(userUsecase)

	// Middleware CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // Izinkan semua domain
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	api := e.Group("/api/v1")

	// USER
	api.POST("/login", userController.UserLogin)
	api.POST("/register", userController.UserRegister)

	user := api.Group("/user")
	user.Use(middlewares.JWTMiddleware, middlewares.RoleMiddleware("user"))

	// user account
	user.Any("", userController.UserCredential)
	user.PATCH("/update-information", userController.UserUpdateInformation)
	user.PUT("/update-password", userController.UserUpdatePassword)
	user.PUT("/update-profile", userController.UserUpdateProfile)
	// user.PUT("/update-photo-profile", userController.UserUpdatePhotoProfile)
	user.DELETE("/delete-photo-profile", userController.UserDeletePhotoProfile)

	// ADMIN

	admin := api.Group("/admin")
	admin.Use(middlewares.JWTMiddleware, middlewares.RoleMiddleware("admin"))

	// crud arcticle

	// admin.GET("/article", articleController.GetAllArticles)
	// admin.GET("/article/:id", articleController.GetArticleByID)
	// admin.PUT("/article/:id", articleController.UpdateArticle)
	// admin.POST("/article", articleController.CreateArticle)
	// admin.DELETE("/article/:id", articleController.DeleteArticle)

}
