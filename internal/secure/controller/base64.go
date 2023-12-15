package controller

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"

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

	// Get the AES key from user input
	key := []byte(c.FormValue("key"))

	// Create a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Encrypt the file bytes using AES
	encryptedBytes := make([]byte, aes.BlockSize+len(fileBytes))
	iv := encryptedBytes[:aes.BlockSize]
	if _, err := rand.Read(iv); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(encryptedBytes[aes.BlockSize:], fileBytes)

	// Encode the encrypted bytes to base64
	encodedString := base64.StdEncoding.EncodeToString(encryptedBytes)

	// Save the encoded string to a file
	exportPath := "C:\\Users\\Dewa Biara\\OneDrive\\Documents\\Semester VII\\Secure-Docs\\export"
	err = os.MkdirAll(exportPath, os.ModePerm)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	tempFile, err := ioutil.TempFile(exportPath, "encoded_*.txt")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer tempFile.Close()

	_, err = tempFile.WriteString(encodedString)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "success creating encryption",
		"data":    tempFile.Name(),
	})
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

	// Get the AES key from user input
	key := []byte(c.FormValue("key"))

	// Decode the base64-encoded string
	encodedString := string(fileBytes)
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Create a new AES cipher block
	block, err := aes.NewCipher(key)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	// Extract the IV from the decoded bytes
	iv := decodedBytes[:aes.BlockSize]
	decodedBytes = decodedBytes[aes.BlockSize:]

	// Decrypt the decoded bytes using AES
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(decodedBytes, decodedBytes)

	// Save the decoded bytes to a file
	exportPath := "C:\\Users\\Dewa Biara\\OneDrive\\Documents\\Semester VII\\Secure-Docs\\export"
	err = os.MkdirAll(exportPath, os.ModePerm)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	tempFile, err := ioutil.TempFile(exportPath, "decoded_*.jpg")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer tempFile.Close()

	_, err = tempFile.Write(decodedBytes)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "success creating decryption",
		"data":    tempFile.Name(),
	})
}
