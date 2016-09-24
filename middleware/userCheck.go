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

		res := c.Get("res").(map[string]interface{})
		stn := time.Now()
		st := stn.UnixNano()
		res["start"] = st

		userCookie, _ := c.Cookie("username")
		if userCookie != nil {
			user := userCookie.Value()
			if user == "" {
				lev = "guest" // guest, not login
			} else {
				log.WithFields(log.Fields{
					"ua":       c.Request().UserAgent(),
					"username": user,
				}).Info("用户访问")
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
			log.WithFields(log.Fields{
				"ua": c.Request().UserAgent(),
			}).Info("用户访问")
			lev = "guest"
		}
		res["userIs"] = lev
		c.Set("res", res)

		return next(c)
	}
}
