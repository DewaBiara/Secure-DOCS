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

type DecryptionController struct {
	decryptionService service.DecryptionService
	jwtService        jwt_service.JWTService
}

func NewDecryptionController(decryptionService service.DecryptionService, jwtService jwt_service.JWTService) *DecryptionController {
	return &DecryptionController{
		decryptionService: decryptionService,
		jwtService:        jwtService,
	}
}

func (u *DecryptionController) CreateDecryption(c echo.Context) error {
	claims := u.jwtService.GetClaims(&c)
	role := claims["role"].(string)
	if role != "admin" {
		return echo.NewHTTPError(http.StatusForbidden, utils.ErrDidntHavePermission.Error())
	}
	decryption := new(dto.CreateDecryptionRequest)
	if err := c.Bind(decryption); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrBadRequestBody.Error())
	}

	if err := c.Validate(decryption); err != nil {
		return err
	}

	err := u.decryptionService.CreateDecryption(c.Request().Context(), decryption)

	if err != nil {
		switch err {
		case utils.ErrDecryptionAlreadyExist:
			fallthrough
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "success creating decryption",
		"data":    decryption,
	})
}

func (u *DecryptionController) UpdateDecryption(c echo.Context) error {
	claims := u.jwtService.GetClaims(&c)
	role := claims["role"].(string)
	if role != "admin" {
		return echo.NewHTTPError(http.StatusForbidden, utils.ErrDidntHavePermission.Error())
	}
	decryption := new(dto.UpdateDecryptionRequest)
	if err := c.Bind(decryption); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, utils.ErrBadRequestBody.Error())
	}

	if err := c.Validate(decryption); err != nil {
		return err
	}

	err := u.decryptionService.UpdateDecryption(c.Request().Context(), decryption.ID, decryption)
	if err != nil {
		switch err {
		case utils.ErrDecryptionNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		case utils.ErrDecryptionAlreadyExist:
			fallthrough
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success update decryption",
		"data":    decryption,
	})
}

func (u *DecryptionController) GetSingleDecryption(c echo.Context) error {
	decryptionID := c.Param("decryption_id")
	decryption, err := u.decryptionService.GetSingleDecryption(c.Request().Context(), decryptionID)
	if err != nil {
		if err == utils.ErrDecryptionNotFound {
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
			"message": "success getting decryption",
			"data":    decryption,
		})
	default:
		return echo.NewHTTPError(http.StatusForbidden, utils.ErrDidntHavePermission.Error())
	}
}

func (u *DecryptionController) GetPageDecryption(c echo.Context) error {

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

	decryption, err := u.decryptionService.GetPageDecryption(c.Request().Context(), int(pageInt), int(limitInt))
	if err != nil {
		if err == utils.ErrDecryptionNotFound {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success getting document",
		"data":    decryption,
		"meta": echo.Map{
			"page":  pageInt,
			"limit": limitInt,
		},
	})
}

func (d *DecryptionController) DeleteDecryption(c echo.Context) error {
	claims := d.jwtService.GetClaims(&c)
	role := claims["role"].(string)
	if role != "admin" {
		return echo.NewHTTPError(http.StatusForbidden, utils.ErrDidntHavePermission.Error())
	}
	decryptionID := c.Param("decryption_id")
	err := d.decryptionService.DeleteDecryption(c.Request().Context(), decryptionID)
	if err != nil {
		switch err {
		case utils.ErrDecryptionNotFound:
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		default:
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success deleting decryption",
	})
}
