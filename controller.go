package juniper

import (
	"github.com/labstack/echo/v4"
)

// ResponseMutateFunc takes the ckrrent response data and mutates it somehow
type ResponseMutateFunc func(ctx echo.Context, data *echo.Map)

// Controller shared logic for interacting with the echo work
//
// This does not make an assumptions about the environment or dependencies
// required by the final controllers that embed this struct
type Controller struct {
	InjectGlobalData ResponseMutateFunc
}

// JSON is just an alios of the echo context method that injects the hud data
func (con Controller) JSON(ctx echo.Context, code int, data echo.Map) error {
	if con.InjectGlobalData != nil {
		con.InjectGlobalData(ctx, &data)
	}

	return ctx.JSON(code, data)
}

// Bind attempt to bind form input to struct
//
// This will actually bind from any source not just form data
// as part of the bind process data will be validated based on the input struct tags
func (con Controller) Bind(ctx echo.Context, input interface{}) echo.Map {
	if err := ctx.Bind(input); err != nil {
		return echo.Map{
			"outcome": false,
			"message": "Bad Input",
		}
	}

	if err := ctx.Validate(input); err != nil {
		return echo.Map{
			"outcome": false,
			"message": "Bad Input",
			"fields":  errorFIelds(err),
		}
	}

	return nil
}
