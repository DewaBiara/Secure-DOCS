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

type KeyController struct {
	keyService service.KeyService
	jwtService jwt_service.JWTService
}

func NewKeyController(keyService service.KeyService, jwtService jwt_service.JWTService) *KeyController {
	return &KeyController{
		keyService: keyService,
		jwtService: jwtService,
	}
}

func (u *KeyController) CreateKey(c echo.Context) error {
	claims := u.jwtService.GetClaims(&c)
	role := claims["role"].(string)
	if role != "admin" {
		return echo.NewHTTPError(http.StatusForbidden, utils.ErrDidntHavePermission.Error())
	}
	key := new(dto.CreateKeyRequest)
	if err := c.Bind(key); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrBadRequestBody.Error())
	}

	if err := c.Validate(key); err != nil {
		return err
	}

	err := u.keyService.CreateKey(c.Request().Context(), key)

	if err != nil {
		switch err {
		case utils.ErrKeyAlreadyExist:
			fallthrough
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "success creating key",
		"data":    key,
	})
}

func (u *KeyController) UpdateKey(c echo.Context) error {
	claims := u.jwtService.GetClaims(&c)
	role := claims["role"].(string)
	if role != "admin" {
		return echo.NewHTTPError(http.StatusForbidden, utils.ErrDidntHavePermission.Error())
	}
	key := new(dto.UpdateKeyRequest)
	if err := c.Bind(key); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrBadRequestBody.Error())
	}

	if err := c.Validate(key); err != nil {
		return err
	}

	err := u.keyService.UpdateKey(c.Request().Context(), key.ID, key)
	if err != nil {
		switch err {
		case utils.ErrKeyNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		case utils.ErrKeyAlreadyExist:
			fallthrough
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update key",
		"data":    key,
	})
}

func (u *KeyController) GetSingleKey(c echo.Context) error {
	keyID := c.Param("key_id")
	key, err := u.keyService.GetSingleKey(c.Request().Context(), keyID)
	if err != nil {
		if err == utils.ErrKeyNotFound {
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
			"message": "success getting key",
			"data":    key,
		})
	default:
		return echo.NewHTTPError(http.StatusForbidden, utils.ErrDidntHavePermission.Error())
	}
}

func (u *KeyController) GetPageKey(c echo.Context) error {

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

	key, err := u.keyService.GetPageKey(c.Request().Context(), int(pageInt), int(limitInt))
	if err != nil {
		if err == utils.ErrKeyNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success getting key",
		"data":    key,
		"meta": echo.Map{
			"page":  pageInt,
			"limit": limitInt,
		},
	})
}

func (u *KeyController) GetPageKeyByPenerima(c echo.Context) error {

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

	penerimaID := c.QueryParam("penerima_id")

	key, err := u.keyService.GetPageKeyByPenerima(c.Request().Context(), string(penerimaID), int(pageInt), int(limitInt))
	if err != nil {
		if err == utils.ErrKeyNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success getting key",
		"data":    key,
		"meta": echo.Map{
			"page":  pageInt,
			"limit": limitInt,
		},
	})
}

func (d *KeyController) DeleteKey(c echo.Context) error {
	claims := d.jwtService.GetClaims(&c)
	role := claims["role"].(string)
	if role != "admin" {
		return echo.NewHTTPError(http.StatusForbidden, utils.ErrDidntHavePermission.Error())
	}
	keyID := c.Param("key_id")
	err := d.keyService.DeleteKey(c.Request().Context(), keyID)
	if err != nil {
		switch err {
		case utils.ErrKeyNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success deleting key",
	})
}
