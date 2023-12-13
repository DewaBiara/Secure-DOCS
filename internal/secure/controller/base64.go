package controller

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

func EncodeHandler(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	src, err := file.Open()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer src.Close()

	fileBytes, err := ioutil.ReadAll(src)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	encodedString := base64.StdEncoding.EncodeToString(fileBytes)
	return c.String(http.StatusOK, encodedString)
}

func DecodeHandler(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	src, err := file.Open()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer src.Close()

	fileBytes, err := ioutil.ReadAll(src)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	decodedBytes, err := base64.StdEncoding.DecodeString(string(fileBytes))
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	tempFile, err := ioutil.TempFile("", "decoded_*.txt")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer tempFile.Close()

	_, err = tempFile.Write(decodedBytes)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, fmt.Sprintf("Decoded file saved as: %s", tempFile.Name()))
}
