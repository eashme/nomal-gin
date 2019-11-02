package views

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"nomal-gin/cache"
	"nomal-gin/services"
	"nomal-gin/views/verify"
)

type UserView struct {
	us services.IUserService
}

func NewUserView(db *sql.DB, c cache.ICache) *UserView {
	return &UserView{
		services.NewUserService(db, c),
	}
}

func (uv *UserView) BindRoute(route gin.IRouter) {
	route.GET("/users",uv.Users)
}

func (uv *UserView) Users(c *gin.Context){
	paramErr := verify.NewParamErr()
	pageSize := paramErr.CheckLimit(c.Query("page_size"))
	page := paramErr.CheckPage(c.Query("page"))
	if !paramErr.ParamValid(){
		RespErr(c,paramErr.Error())
		return
	}
	data ,err := uv.us.GetAll(int(page),int(pageSize))
	if err != nil{
		RespErr(c,err)
		return
	}
	RespSuccess(c,data)
}