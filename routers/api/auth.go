package api

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/orzalter/go-blog/models"
	"github.com/orzalter/go-blog/pkg/e"
	"github.com/orzalter/go-blog/pkg/logging"
	"github.com/orzalter/go-blog/pkg/util"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// GetAuth use to check password and generate token.
func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS
	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			logging.Error(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
