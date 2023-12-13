package controller

import (
	"net/http"
	"strconv"

	"github.com/DewaBiara/Secure-DOCS/internal/secure/dto"
	"github.com/DewaBiara/Secure-DOCS/internal/secure/service"
	"github.com/DewaBiara/Secure-DOCS/pkg/utils"
	"github.com/DewaBiara/Secure-DOCS/pkg/utils/jwt_service"
	"github.com/labstack/echo/v4"
)

type EncryptionController struct {
	encryptionService service.EncryptionService
	jwtService        jwt_service.JWTService
}

func NewEncryptionController(encryptionService service.EncryptionService, jwtService jwt_service.JWTService) *EncryptionController {
	return &EncryptionController{
		encryptionService: encryptionService,
		jwtService:        jwtService,
	}
}

func (u *EncryptionController) CreateEncryption(c echo.Context) error {
	claims := u.jwtService.GetClaims(&c)
	role := claims["role"].(string)
	if role != "admin" {
		return echo.NewHTTPError(http.StatusForbidden, utils.ErrDidntHavePermission.Error())
	}
	encryption := new(dto.CreateEncryptionRequest)
	if err := c.Bind(encryption); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrBadRequestBody.Error())
	}

	if err := c.Validate(encryption); err != nil {
		return err
	}

	err := u.encryptionService.CreateEncryption(c.Request().Context(), encryption)

	if err != nil {
		switch err {
		case utils.ErrEncryptionAlreadyExist:
			fallthrough
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "success creating encryption",
		"data":    encryption,
	})
}

func (u *EncryptionController) UpdateEncryption(c echo.Context) error {
	claims := u.jwtService.GetClaims(&c)
	role := claims["role"].(string)
	if role != "admin" {
		return echo.NewHTTPError(http.StatusForbidden, utils.ErrDidntHavePermission.Error())
	}
	encryption := new(dto.UpdateEncryptionRequest)
	if err := c.Bind(encryption); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrBadRequestBody.Error())
	}

	if err := c.Validate(encryption); err != nil {
		return err
	}

	err := u.encryptionService.UpdateEncryption(c.Request().Context(), encryption.ID, encryption)
	if err != nil {
		switch err {
		case utils.ErrEncryptionNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		case utils.ErrEncryptionAlreadyExist:
			fallthrough
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update encryption",
		"data":    encryption,
	})
}

func (u *EncryptionController) GetSingleEncryption(c echo.Context) error {
	encryptionID := c.Param("encryption_id")
	encryption, err := u.encryptionService.GetSingleEncryption(c.Request().Context(), encryptionID)
	if err != nil {
		if err == utils.ErrEncryptionNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	claims := u.jwtService.GetClaims(&c)
	role := claims["role"].(string)

	switch {
	case role == "pegawai":
		fallthrough
	case role == "admin":
		return c.JSON(http.StatusOK, echo.Map{
			"message": "success getting encryption",
			"data":    encryption,
		})
	default:
		return echo.NewHTTPError(http.StatusForbidden, utils.ErrDidntHavePermission.Error())
	}
}

func (u *EncryptionController) GetPageEncryption(c echo.Context) error {

	page := c.QueryParam("page")
	if page == "" {
		page = "1"
	}
	pageInt, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrInvalidNumber.Error())
	}

	limit := c.QueryParam("limit")
	if limit == "" {
		limit = "20"
	}
	limitInt, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrInvalidNumber.Error())
	}

	encryption, err := u.encryptionService.GetPageEncryption(c.Request().Context(), int(pageInt), int(limitInt))
	if err != nil {
		if err == utils.ErrEncryptionNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success getting document",
		"data":    encryption,
		"meta": echo.Map{
			"page":  pageInt,
			"limit": limitInt,
		},
	})
}

func (d *EncryptionController) DeleteEncryption(c echo.Context) error {
	claims := d.jwtService.GetClaims(&c)
	role := claims["role"].(string)
	if role != "admin" {
		return echo.NewHTTPError(http.StatusForbidden, utils.ErrDidntHavePermission.Error())
	}
	encryptionID := c.Param("encryption_id")
	err := d.encryptionService.DeleteEncryption(c.Request().Context(), encryptionID)
	if err != nil {
		switch err {
		case utils.ErrEncryptionNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success deleting encryption",
	})
}
