package routes

import (
	"backbone-dapodik/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Static("/", "uploads")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Bagaimana Kabar Anda!")
	})

	e.GET("/syncwilayah", controllers.SyncWilayah)
	e.GET("/syncsekolah", controllers.SyncSekolah)
	e.GET("/syncsekolahid", controllers.SyncSekolahById)
	e.GET("/syncptk", controllers.SyncPtk)
	e.GET("/ptk", controllers.GETPtk)
	e.GET("/pembelajaran", controllers.GETPembelajaran)
	e.GET("/listfile", controllers.ListFile)
	return e
}
