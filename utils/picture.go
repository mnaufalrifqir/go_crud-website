package utils

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func SavePicture(pictureFile *multipart.FileHeader) string {
	encodedImage := EncodeBase64(pictureFile)
	url := DecodeAndWriteBase64(encodedImage)
	return url
}

func CheckPicture(c echo.Context) error {
	url := c.Param("url")
	url = "assets/images/" + url

	// Memeriksa keberadaan file gambar
	_, err := os.Stat(url)

	if os.IsNotExist(err) {
		// Jika file tidak ditemukan, kirimkan pesan error
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"status":  404,
			"message": "not found",
		})
	} else {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": 200,
		})
	}
}

func EncodeBase64(pictureFile *multipart.FileHeader) string {

	// Buka file gambar
	src, err_open := pictureFile.Open()
	if err_open != nil {
		log.Println(err_open)
		return "Failed to open image file"
	}
	defer src.Close()

	// Baca isi file gambar
	imageData, err_read := ioutil.ReadAll(src)
	if err_read != nil {
		log.Println(err_read)
		return "Failed to read image file"
	}

	encodedImage := base64.StdEncoding.EncodeToString(imageData)

	return encodedImage
}

func DecodeAndWriteBase64(stringBase64 string) string {
	dummyURL := "dummy.png"
	stringBase64 = strings.TrimPrefix(stringBase64, "data:image/png;base64,")

	decodedData, errDecode := base64.StdEncoding.DecodeString(stringBase64)
	if errDecode != nil {
		fmt.Println(errDecode)
		return dummyURL
	}

	imageURL := uuid.New().String() + ".png"
	file, errCreate := os.Create("assets/images/" + imageURL)
	if errCreate != nil {
		fmt.Println(errCreate)
		return dummyURL
	}
	defer file.Close()

	_, errWrite := file.Write(decodedData)
	if errWrite != nil {
		fmt.Println(errWrite)
		return dummyURL
	}

	return imageURL
}
