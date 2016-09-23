package middleware

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gogather/com"
	"github.com/gojudge/ojsite/models"
	"github.com/labstack/echo"
	"time"
)

func UserCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var lev string
		res := make(map[string]interface{})

		stn := time.Now()
		st := stn.UnixNano()
		res["start"] = st

		log.WithFields(log.Fields{
			"ua": c.Request().UserAgent(),
		}).Info("访问日志")

		userCookie, err := c.Cookie("username")
		if err != nil && userCookie != nil {
			user := userCookie.Value()
			log.WithFields(log.Fields{
				"cookie": userCookie,
			}).Info("用户Cookie")
			if user == "" {
				lev = "guest" // guest, not login
			} else {
				levelCookie, err := c.Cookie("level")
				if err != nil && levelCookie != nil {
					level := levelCookie.Value()
					if level == "" {
						lev = "user"
					} else {
						lev = level
					}

					username := user
					usr := models.User{}
					u, err := usr.GetUser(0, username, "", "")
					if err != nil {
						res["nickname"] = ""
						res["email_md5"] = ""
					} else {
						res["username"] = username
						res["nickname"] = u.Nickname
						res["email_md5"] = com.Md5(u.Email)
					}
				} else {
					lev = "user"
				}
			}
		} else {
			lev = "guest"
		}
		res["userIs"] = lev
		c.Set("res", res)

		return next(c)
	}
}
