package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/llr104/slgserver/constant"
	"github.com/llr104/slgserver/log"
	myhttp "github.com/llr104/slgserver/server/httpserver"
	"github.com/llr104/slgserver/server/httpserver/logic"
	"github.com/llr104/slgserver/server/httpserver/middleware"
)

type AccountController struct{}

func (self AccountController) RegisterRoutes(g *echo.Group) {
	g.Use(middleware.Cors())

	g.Any("/account/register", self.register)
	g.Any("/account/changepwd", self.changePwd)
	g.Any("/account/forgetpwd", self.forgetPwd)
	g.Any("/account/resetpwd", self.resetPwd)
}

func (self AccountController) register(ctx echo.Context) error {
	log.DefaultLog.Info("register")
	data := make(map[string]interface{})
	if err := logic.DefaultUser.CreateUser(ctx); err != nil {
		data["code"] = err.(*myhttp.MyError).Id()
		data["errmsg"] = err.(*myhttp.MyError).Error()
	} else {
		data["code"] = constant.OK
	}

	ctx.JSON(http.StatusOK, data)

	return nil
}

func (self AccountController) forgetPwd(ctx echo.Context) error {
	log.DefaultLog.Info("forgetPwd")
	ctx.String(http.StatusOK, "forgetPwd")
	return nil
}

func (self AccountController) changePwd(ctx echo.Context) error {
	log.DefaultLog.Info("changePwd")
	data := make(map[string]interface{})
	if err := logic.DefaultUser.ChangePassword(ctx); err != nil {
		data["code"] = err.(*myhttp.MyError).Id()
		data["errmsg"] = err.(*myhttp.MyError).Error()
	} else {
		data["code"] = constant.OK
	}

	ctx.JSON(http.StatusOK, data)
	return nil
}

func (self AccountController) resetPwd(ctx echo.Context) error {
	log.DefaultLog.Info("resetPwd")
	ctx.String(http.StatusOK, "resetPwd")
	return nil
}
