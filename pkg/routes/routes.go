package routes

import (
	secureControllerPkg "github.com/DewaBiara/Secure-DOCS/internal/secure/controller"
	userControllerPkg "github.com/DewaBiara/Secure-DOCS/internal/user/controller"
	"github.com/DewaBiara/Secure-DOCS/pkg/utils/validation"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Routes struct {
	userController       *userControllerPkg.UserController
	encryptionController *secureControllerPkg.EncryptionController
	decryptionController *secureControllerPkg.DecryptionController
	keyController        *secureControllerPkg.KeyController
}

func NewRoutes(userController *userControllerPkg.UserController, encryptionController *secureControllerPkg.EncryptionController, decryptionController *secureControllerPkg.DecryptionController,
	keyController *secureControllerPkg.KeyController) *Routes {
	return &Routes{
		userController:       userController,
		encryptionController: encryptionController,
		decryptionController: decryptionController,
		keyController:        keyController,
	}
}

func (r *Routes) Init(e *echo.Echo, conf map[string]string) {
	e.Pre(middleware.AddTrailingSlash())
	e.Use(middleware.Recover())

	e.Validator = &validation.CustomValidator{Validator: validator.New()}

	jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(conf["JWT_SECRET"]),
	})

	v1 := e.Group("/v1")

	// Users
	users := v1.Group("/users")
	users.POST("/signup/", r.userController.SignUpUser)
	users.POST("/login/", r.userController.LoginUser)

	usersWithAuth := users.Group("", jwtMiddleware)
	usersWithAuth.GET("/", r.userController.GetBriefUsers)
	usersWithAuth.PUT("/", r.userController.UpdateUser)

	// Encryptions
	encryptions := v1.Group("/encryptions")
	encryptions.POST("/", r.encryptionController.CreateEncryption, jwtMiddleware)
	encryptions.PUT("/", r.encryptionController.UpdateEncryption, jwtMiddleware)
	encryptions.GET("/:encryption_id/", r.encryptionController.GetSingleEncryption, jwtMiddleware)
	encryptions.GET("/", r.encryptionController.GetPageEncryption)
	encryptions.DELETE("/:encryption_id/", r.encryptionController.DeleteEncryption, jwtMiddleware)

	// Decryptions
	decryptions := v1.Group("/decryptions")
	decryptions.POST("/", r.decryptionController.CreateDecryption, jwtMiddleware)
	decryptions.PUT("/", r.decryptionController.UpdateDecryption, jwtMiddleware)
	decryptions.GET("/:decryption_id/", r.decryptionController.GetSingleDecryption, jwtMiddleware)
	decryptions.GET("/", r.decryptionController.GetPageDecryption)
	decryptions.DELETE("/:decryption_id/", r.decryptionController.DeleteDecryption, jwtMiddleware)

	// Keys
	keys := v1.Group("/keys")
	keys.POST("/", r.keyController.CreateKey)
	keys.PUT("/", r.keyController.UpdateKey, jwtMiddleware)
	keys.GET("/:key_id/", r.keyController.GetSingleKey, jwtMiddleware)
	keys.GET("/", r.keyController.GetPageKey)
	keys.DELETE("/:key_id/", r.keyController.DeleteKey, jwtMiddleware)

	// //Base 64
	// base64 := v1.Group("/base64")
	// base64.POST("/", controller.EncodeHandler)
	// //base64.POST("/decode", controller.DecodeHandler)
}
