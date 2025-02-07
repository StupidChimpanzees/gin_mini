package controller

import (
	"gin_work/common"
	"gin_work/extend/jwt"
	"gin_work/message"
	"gin_work/model"
	"gin_work/wrap/cache"
	"gin_work/wrap/cookie"
	"gin_work/wrap/response"
	"github.com/gin-gonic/gin"
)

type Login struct{}

type UserLoginForm struct {
	Username string `form:"username" json:"username" uri:"username" xml:"username" binding:"required,alphanum,min=5,max=20"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required,alphanum,min=8,max=20"`
	Remember bool   `form:"remember" json:"remember" uri:"remember" xml:"remember" binding:"omitempty,oneof=true false"`
}

func (l *Login) UserLogin(c *gin.Context) {
	var login UserLoginForm
	if err := c.ShouldBind(&login); err != nil {
		c.JSON(response.RequestFail(message.RequestError, err.Error()))
		return
	}

	var user model.User
	row, err := user.FindByUsername(login.Username)
	if err != nil {
		common.Warning(err, c, login)
		c.JSON(response.Fail(err))
		return
	} else if row == 0 {
		c.JSON(response.RequestFail(message.UsernameOrPasswordError))
		return
	}
	if code := common.CheckPwd(login.Password, user.Password, user.Salt); code != message.Success {
		c.JSON(response.RequestFail(code))
		return
	}

	// 生成token
	token, err := l.createToken(c, &user)
	if err != nil {
		common.Warning(err, c, login)
		c.JSON(response.Fail(message.ServerError, err.Error()))
		return
	}

	if login.Remember {
		encryptPwd := common.GetPwd(user.Password, c.ClientIP())
		cookie.Set(c, "username", login.Username)
		cookie.Set(c, "password", encryptPwd)
		c.JSON(response.Success(map[string]string{jwt.Token: token}))
	} else {
		c.JSON(response.Success(map[string]string{jwt.Token: token}))
	}

	return
}

func (*Login) createToken(c *gin.Context, user *model.User) (string, error) {
	accessToken, err := common.RefreshToken(user.Uuid, c.Request.Host, c.ClientIP())
	if err != nil {
		return "", err
	}
	_ = cache.Set(jwt.AuthPrefix+user.Uuid, accessToken)
	return accessToken, nil
}
