package juniper

import (
	"github.com/labstack/echo/v4"
)

var (
	ResponseSuccess      = echo.Map{"outcome": true, "message": nil}
	ResponseNotFound     = echo.Map{"outcome": false, "message": "Not Found"}
	ResponseFailed       = echo.Map{"outcome": false, "message": "Failed"}
	ResponseBadInput     = echo.Map{"outcome": false, "message": "Bad Input"}
	ResponseUnauthorized = echo.Map{"outcome": false, "message": "Unauthorized"}
)
