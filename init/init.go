package init

import (
	"github.com/labstack/echo"
)

func Initial() *echo.Echo {
	e := echo.New()
	e.Debug()

	return e
}
