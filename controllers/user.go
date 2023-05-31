package controllers

import (
	"my-me/dtos"
	"my-me/helpers"
	"my-me/middlewares"
	"my-me/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase usecases.UserUsecase
}

func NewUserController(userUsecase usecases.UserUsecase) UserController {
	return UserController{userUsecase}
}

func (c *UserController) UserLogin(ctx echo.Context) error {
	var userInput dtos.UserLoginInput

	err := ctx.Bind(&userInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to login",
				helpers.GetErrorData(err),
			),
		)
	}

	user, err := c.userUsecase.UserLogin(userInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to login",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully logged in",
			user,
		),
	)
}

func (c *UserController) UserRegister(ctx echo.Context) error {
	var userInput dtos.UserRegisterInput
	err := ctx.Bind(&userInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to register",
				helpers.GetErrorData(err),
			),
		)
	}

	user, err := c.userUsecase.UserRegister(userInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to register",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusCreated,
		helpers.NewResponse(
			http.StatusCreated,
			"Successfully registered",
			user,
		),
	)
}

func (c *UserController) UserUpdateInformation(ctx echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(ctx.Request())
	if tokenString == "" {
		return ctx.JSON(
			http.StatusUnauthorized,
			helpers.NewErrorResponse(
				http.StatusUnauthorized,
				"No token provided",
				helpers.GetErrorData(nil),
			),
		)
	}

	userId, err := middlewares.GetUserIdFromToken(tokenString)
	if err != nil {
		return ctx.JSON(
			http.StatusUnauthorized,
			helpers.NewErrorResponse(
				http.StatusUnauthorized,
				"No token provided",
				helpers.GetErrorData(err),
			),
		)
	}

	var userInput dtos.UserUpdateInformationInput
	err = ctx.Bind(&userInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to update information",
				helpers.GetErrorData(err),
			),
		)
	}

	user, err := c.userUsecase.UserUpdateInformation(userId, userInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to update information",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully updated information",
			user,
		),
	)
}

func (c *UserController) UserUpdatePassword(ctx echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(ctx.Request())
	if tokenString == "" {
		return ctx.JSON(
			http.StatusUnauthorized,
			helpers.NewErrorResponse(
				http.StatusUnauthorized,
				"No token provided",
				helpers.GetErrorData(nil),
			),
		)
	}

	userId, err := middlewares.GetUserIdFromToken(tokenString)
	if err != nil {
		return ctx.JSON(
			http.StatusUnauthorized,
			helpers.NewErrorResponse(
				http.StatusUnauthorized,
				"No token provided",
				helpers.GetErrorData(err),
			),
		)
	}

	var userInput dtos.UserUpdatePasswordInput
	err = ctx.Bind(&userInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to update password",
				helpers.GetErrorData(err),
			),
		)
	}

	user, err := c.userUsecase.UserUpdatePassword(userId, userInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to update password",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully updated password",
			user,
		),
	)
}

func (c *UserController) UserUpdateProfile(ctx echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(ctx.Request())
	if tokenString == "" {
		return ctx.JSON(
			http.StatusUnauthorized,
			helpers.NewErrorResponse(
				http.StatusUnauthorized,
				"No token provided",
				helpers.GetErrorData(nil),
			),
		)
	}

	userId, err := middlewares.GetUserIdFromToken(tokenString)
	if err != nil {
		return ctx.JSON(
			http.StatusUnauthorized,
			helpers.NewErrorResponse(
				http.StatusUnauthorized,
				"No token provided",
				helpers.GetErrorData(err),
			),
		)
	}

	var userInput dtos.UserUpdateProfileInput
	err = ctx.Bind(&userInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to update profile",
				helpers.GetErrorData(err),
			),
		)
	}

	user, err := c.userUsecase.UserUpdateProfile(userId, userInput)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to update profile",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully updated profile",
			user,
		),
	)
}

func (c *UserController) UserCredential(ctx echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(ctx.Request())
	if tokenString == "" {
		return ctx.JSON(
			http.StatusUnauthorized,
			helpers.NewErrorResponse(
				http.StatusUnauthorized,
				"No token provided",
				helpers.GetErrorData(nil),
			),
		)
	}

	userId, err := middlewares.GetUserIdFromToken(tokenString)
	if err != nil {
		return ctx.JSON(
			http.StatusUnauthorized,
			helpers.NewErrorResponse(
				http.StatusUnauthorized,
				"No token provided",
				helpers.GetErrorData(err),
			),
		)
	}

	user, err := c.userUsecase.UserCredential(userId)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to get user credentials",
				helpers.GetErrorData(err),
			),
		)
	}
	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully get user credentials",
			user,
		),
	)
}

// func (c *UserController) UserUpdatePhotoProfile(ctx echo.Context) error {
// 	tokenString := middlewares.GetTokenFromHeader(ctx.Request())
// 	if tokenString == "" {
// 		return ctx.JSON(
// 			http.StatusUnauthorized,
// 			helpers.NewErrorResponse(
// 				http.StatusUnauthorized,
// 				"No token provided",
// 				helpers.GetErrorData(nil),
// 			),
// 		)
// 	}

// 	userId, err := middlewares.GetUserIdFromToken(tokenString)
// 	if err != nil {
// 		return ctx.JSON(
// 			http.StatusUnauthorized,
// 			helpers.NewErrorResponse(
// 				http.StatusUnauthorized,
// 				"No token provided",
// 				helpers.GetErrorData(err),
// 			),
// 		)
// 	}

// 	var userInput dtos.UserUpdatePhotoProfileInput
// 	err = ctx.Bind(&userInput)
// 	if err != nil {
// 		return ctx.JSON(
// 			http.StatusBadRequest,
// 			helpers.NewErrorResponse(
// 				http.StatusBadRequest,
// 				"Failed to update information",
// 				helpers.GetErrorData(err),
// 			),
// 		)
// 	}

// 	if userInput.ProfilePicture == "" {
// 		formHeader, err := ctx.FormFile("file")
// 		if err != nil {
// 			if err != nil {
// 				return ctx.JSON(
// 					http.StatusInternalServerError,
// 					helpers.NewErrorResponse(
// 						http.StatusInternalServerError,
// 						"Error uploading photo",
// 						helpers.GetErrorData(err),
// 					),
// 				)
// 			}
// 		}

// 		//get file from header
// 		formFile, err := formHeader.Open()
// 		if err != nil {
// 			return ctx.JSON(
// 				http.StatusInternalServerError,
// 				helpers.NewErrorResponse(
// 					http.StatusInternalServerError,
// 					"Error uploading photo",
// 					helpers.GetErrorData(err),
// 				),
// 			)
// 		}

// 		var re = regexp.MustCompile(`.png|.jpeg|.jpg`)

// 		if !re.MatchString(formHeader.Filename) {
// 			return ctx.JSON(
// 				http.StatusBadRequest,
// 				helpers.NewErrorResponse(
// 					http.StatusBadRequest,
// 					"The provided file format is not allowed. Please upload a JPEG or PNG image",
// 					helpers.GetErrorData(err),
// 				),
// 			)
// 		}

// 		uploadUrl, err := usecases.NewMediaUpload().FileUpload(models.File{File: formFile})

// 		if err != nil {
// 			return ctx.JSON(
// 				http.StatusInternalServerError,
// 				helpers.NewErrorResponse(
// 					http.StatusInternalServerError,
// 					"Error uploading photo",
// 					helpers.GetErrorData(err),
// 				),
// 			)
// 		}
// 		userInput.ProfilePicture = uploadUrl
// 	} else {
// 		var url models.Url
// 		url.Url = userInput.ProfilePicture

// 		var re = regexp.MustCompile(`.png|.jpeg|.jpg`)
// 		if !re.MatchString(userInput.ProfilePicture) {
// 			return ctx.JSON(
// 				http.StatusBadRequest,
// 				helpers.NewErrorResponse(
// 					http.StatusBadRequest,
// 					"The provided file format is not allowed. Please upload a JPEG or PNG image",
// 					helpers.GetErrorData(err),
// 				),
// 			)
// 		}

// 		uploadUrl, err := usecases.NewMediaUpload().RemoteUpload(url)
// 		if uploadUrl == "" || err != nil {
// 			return ctx.JSON(
// 				http.StatusInternalServerError,
// 				helpers.NewErrorResponse(
// 					http.StatusInternalServerError,
// 					"Error uploading photo",
// 					helpers.GetErrorData(err),
// 				),
// 			)
// 		}

// 		userInput.ProfilePicture = uploadUrl
// 	}

// 	user, err := c.userUsecase.UserUpdatePhotoProfile(userId, userInput)
// 	if err != nil {
// 		return ctx.JSON(
// 			http.StatusBadRequest,
// 			helpers.NewErrorResponse(
// 				http.StatusBadRequest,
// 				"Failed to update information",
// 				helpers.GetErrorData(err),
// 			),
// 		)
// 	}

// 	return ctx.JSON(
// 		http.StatusOK,
// 		helpers.NewResponse(
// 			http.StatusOK,
// 			"Successfully updated information",
// 			user,
// 		),
// 	)
// }

func (c *UserController) UserDeletePhotoProfile(ctx echo.Context) error {
	tokenString := middlewares.GetTokenFromHeader(ctx.Request())
	if tokenString == "" {
		return ctx.JSON(
			http.StatusUnauthorized,
			helpers.NewErrorResponse(
				http.StatusUnauthorized,
				"No token provided",
				helpers.GetErrorData(nil),
			),
		)
	}

	userId, err := middlewares.GetUserIdFromToken(tokenString)
	if err != nil {
		return ctx.JSON(
			http.StatusUnauthorized,
			helpers.NewErrorResponse(
				http.StatusUnauthorized,
				"No token provided",
				helpers.GetErrorData(err),
			),
		)
	}

	user, err := c.userUsecase.UserDeletePhotoProfile(userId)
	if err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			helpers.NewErrorResponse(
				http.StatusBadRequest,
				"Failed to update profile",
				helpers.GetErrorData(err),
			),
		)
	}

	return ctx.JSON(
		http.StatusOK,
		helpers.NewResponse(
			http.StatusOK,
			"Successfully updated profile",
			user,
		),
	)
}
