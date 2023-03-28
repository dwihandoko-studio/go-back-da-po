package controllers

import (
	"backbone-dapodik/models"
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/labstack/echo/v4"
)

func GETPembelajaran(c echo.Context) error {
	keyGet := c.Request().Header.Get("X-API-TOKEN")

	if keyGet != "0b4e06f30dc26c36f322580591e0a07b" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 401, Message: "Token tidak valid"},
		)
	}

	var id_sekolah string

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)

	if err != nil {
		id_sekolah = c.FormValue("id_sekolah")
	} else {
		id_sekolah = fmt.Sprintf("%s", json_map["id_sekolah"])
	}

	if id_sekolah == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 400, Message: "ID Sekolah tidak boleh kosong dan harus valid."},
		)
	}

	result, err := models.GetPembelajaranBackbone(id_sekolah)
	if err != nil {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: err.Error()})
	}

	return c.JSON(http.StatusOK,
		result,
		// models.Response{
		// 	Status:  200,
		// 	Message: "success",
		// 	Data:    result,
		// }
	)
}

func GETPtk(c echo.Context) error {
	keyGet := c.Request().Header.Get("X-API-TOKEN")

	if keyGet != "0b4e06f30dc26c36f322580591e0a07b" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 401, Message: "Token tidak valid"},
		)
	}

	var id_sekolah string

	json_map := make(map[string]interface{})
	err := json.NewDecoder(c.Request().Body).Decode(&json_map)

	if err != nil {
		id_sekolah = c.FormValue("id_sekolah")
	} else {
		id_sekolah = fmt.Sprintf("%s", json_map["id_sekolah"])
	}

	if id_sekolah == "" {
		return c.JSON(http.StatusOK,
			models.Response{Status: 400, Message: "ID Sekolah tidak boleh kosong dan harus valid."},
		)
	}

	result, err := models.GetPtkBackbone(id_sekolah)
	if err != nil {
		return c.JSON(http.StatusOK,
			models.Response{Status: 404, Message: err.Error()})
	}

	return c.JSON(http.StatusOK,
		result,
		// models.Response{
		// 	Status:  200,
		// 	Message: "success",
		// 	Data:    result,
		// }
	)
}
