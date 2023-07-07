package controllers

import (
	"net/http"
	"os"
	"strconv"
	"telkomsel/database"
	"telkomsel/models"
	"telkomsel/utils"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// Create
func CreateUser(c echo.Context) error {
	validate := validator.New()

	user := models.User{}
	user.Nama = c.FormValue("nama")
	user.Email = c.FormValue("email")
	user.TanggalLahir = c.FormValue("tanggal_lahir")
	user.NoHP = c.FormValue("no_hp")
	user.JenisKelamin = c.FormValue("jenis_kelamin")

	file, err := c.FormFile("url_foto")
	if err != nil {
		return err
	}

	user.UrlFoto = utils.SavePicture(file)

	err = validate.Struct(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed create input invalid",
			"error":   err.Error(),
		})
	}

	if err := database.DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed create",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Create",
		"status":  200,
	})
}

// Get user by id
func GetUserByID(c echo.Context) error {
	user := models.User{}
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := database.DB.First(&user, userID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get user by id",
		"data":    user,
	})
}

// Read all users
func ReadAllUsers(c echo.Context) error {
	users := []models.User{}
	if err := database.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get all users",
		"data":    users,
	})
}

// Update
func UpdateUserByID(c echo.Context) error {
	validate := validator.New()
	user := models.User{}
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := database.DB.First(&user, userID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	os.Remove("assets/images/" + user.UrlFoto)

	user.Nama = c.FormValue("nama")
	user.Email = c.FormValue("email")
	user.TanggalLahir = c.FormValue("tanggal_lahir")
	user.NoHP = c.FormValue("no_hp")
	user.JenisKelamin = c.FormValue("jenis_kelamin")

	file, err := c.FormFile("url_foto")
	if err != nil {
		return err
	}

	user.UrlFoto = utils.SavePicture(file)

	err = validate.Struct(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed create input invalid",
			"error":   err.Error(),
		})
	}

	if err := database.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success update",
	})
}

// Delete
func DeleteUserByID(c echo.Context) error {
	user := models.User{}
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	if err := database.DB.First(&user, userID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	os.Remove("assets/images/" + user.UrlFoto)

	if err := database.DB.Unscoped().Delete(&user, userID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Deleted successfully",
	})
}
