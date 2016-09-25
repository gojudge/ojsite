package middleware

import (
	log "github.com/Sirupsen/logrus"
	"github.com/gogather/com"
	"github.com/gojudge/ojsite/models"
	"github.com/labstack/echo"
	"strconv"
	"time"
)

func UserCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var lev string
		userID := 0

		res := c.Get("res").(map[string]interface{})
		stn := time.Now()
		st := stn.UnixNano()
		res["start"] = st

		userCookie, _ := c.Cookie("userid")
		if userCookie != nil {
			user := userCookie.Value()
			if user == "" {
				lev = "guest" // guest, not login
			} else {
				userID, _ = strconv.Atoi(user)
				usr := models.User{}
				u, err := usr.GetUser(userID, "", "", "")
				if err != nil {
					res["nickname"] = ""
					res["email_md5"] = ""
					lev = "user not found"
				} else {
					c.Set("userID", u.Id)
					res["username"] = u.Username
					res["nickname"] = u.Nickname
					res["email_md5"] = com.Md5(u.Email)

					levelCookie, _ := c.Cookie("level")
					if levelCookie != nil {
						level := levelCookie.Value()
						if level == "" {
							lev = "user"
						} else {
							lev = level
						}
					} else {
						lev = "user"
					}
				}
			}
		} else {
			lev = "guest"
		}
		log.WithFields(log.Fields{
			"ua":     c.Request().UserAgent(),
			"userID": userID,
			"level":  lev,
		}).Info("用户访问")

		res["userIs"] = lev
		c.Set("userLevel", lev)
		c.Set("res", res)

		return next(c)
	}
}
