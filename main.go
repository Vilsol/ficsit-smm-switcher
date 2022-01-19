package main

import (
	"github.com/labstack/echo/v4"
	"strings"
)

const smmWindows = "https://github.com/satisfactorymodding/SatisfactoryModManager/releases/latest/download/Satisfactory-Mod-Manager-Setup.exe"
const smmLinux = "https://github.com/satisfactorymodding/SatisfactoryModManager/releases/latest/download/Satisfactory-Mod-Manager.AppImage"

func main() {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.GET("/", func(context echo.Context) error {
		userAgent := context.Request().UserAgent()
		if userAgent != "" {
			if strings.Contains(strings.ToLower(userAgent), "linux") {
				return context.Redirect(301, smmLinux)
			}
		}

		return context.Redirect(301, smmWindows)
	})

	_ = e.Start(":8080")
}
