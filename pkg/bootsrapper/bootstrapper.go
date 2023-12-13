package bootsrapper

import (
	"time"

	secureControllerPkg "github.com/DewaBiara/Secure-DOCS/internal/secure/controller"
	secureRepositoryPkg "github.com/DewaBiara/Secure-DOCS/internal/secure/repository/impl"
	secureServicePkg "github.com/DewaBiara/Secure-DOCS/internal/secure/service/impl"
	userControllerPkg "github.com/DewaBiara/Secure-DOCS/internal/user/controller"
	userRepositoryPkg "github.com/DewaBiara/Secure-DOCS/internal/user/repository/impl"
	userServicePkg "github.com/DewaBiara/Secure-DOCS/internal/user/service/impl"
	"github.com/DewaBiara/Secure-DOCS/pkg/routes"
	"github.com/DewaBiara/Secure-DOCS/pkg/utils/aes"
	jwtPkg "github.com/DewaBiara/Secure-DOCS/pkg/utils/jwt_service/impl"
	passwordPkg "github.com/DewaBiara/Secure-DOCS/pkg/utils/password/impl"
	"github.com/labstack/echo/v4"

	"gorm.io/gorm"
)

func InitController(e *echo.Echo, db *gorm.DB, conf map[string]string) {
	passwordFunc := passwordPkg.NewPasswordFuncImpl()
	jwtService := jwtPkg.NewJWTService(conf["JWT_SECRET"], 1*time.Hour)

	// User
	userRepository := userRepositoryPkg.NewUserRepositoryImpl(db)
	userService := userServicePkg.NewUserServiceImpl(userRepository, passwordFunc, jwtService)
	userController := userControllerPkg.NewUserController(userService, jwtService)

	//Encryption
	encryptionRepository := secureRepositoryPkg.NewEncryptionRepositoryImpl(db)
	encryptionService := secureServicePkg.NewEncryptionServiceImpl(encryptionRepository, aes.AESFileCrypter{})
	encryptionController := secureControllerPkg.NewEncryptionController(encryptionService, jwtService)

	//Decryption
	decryptionRepository := secureRepositoryPkg.NewDecryptionRepositoryImpl(db)
	decryptionService := secureServicePkg.NewDecryptionServiceImpl(decryptionRepository, aes.AESFileCrypter{})
	decryptionController := secureControllerPkg.NewDecryptionController(decryptionService, jwtService)

	//Key
	keyRepository := secureRepositoryPkg.NewKeyRepositoryImpl(db)
	keyService := secureServicePkg.NewKeyServiceImpl(keyRepository)
	keyController := secureControllerPkg.NewKeyController(keyService, jwtService)

	route := routes.NewRoutes(userController, encryptionController, decryptionController, keyController)
	route.Init(e, conf)
}
