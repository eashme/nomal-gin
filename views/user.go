package views

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"nomal-gin/cache"
	"nomal-gin/services"
	"nomal-gin/views/verify"
)

type userView struct {
	us services.IUserService
}

func NewUserView(db *sql.DB, c cache.ICache) *userView {
	return &userView{
		services.NewUserService(db, c),
	}
}

func (uv *userView) BindRoute(route gin.IRouter) {
	route.GET("/users", uv.Users)
	route.GET("/user/:id", uv.User)
	route.POST("/user", uv.Add)
	route.PUT("/user", uv.Update)
	route.DELETE("/user",uv.Delete)
}

// @Summary 用户列表接口
// @Description 根据条件筛选用户
// @Tags user
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param name query  int  false "姓名"
// @Param age query  int  false "年龄"
// @Success 200 {object} docs.ResponseUsers "success"
// @Failure 400 {object} docs.ResponseInvalidParam "bad parameters"
// @Failure 500 {object} docs.ResponseServerErr "internal server error"
// @Router /users [get]
func (uv *userView) Users(c *gin.Context) {
	paramErr := verify.NewParamErr()
	pageSize := paramErr.CheckLimit(c.Query("page_size"))
	page := paramErr.CheckPage(c.Query("page"))
	if !paramErr.ParamValid() {
		RespErr(c, paramErr.Error())
		return
	}
	data, err := uv.us.GetAll(int(page), int(pageSize))
	if err != nil {
		RespErr(c, err)
		return
	}
	RespSuccess(c, data)
}

// @Summary 根据用户ID
// @Description 根据用户ID,获取用户
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {array} docs.ResponseUser "success"
// @Failure 400 {object} docs.ResponseInvalidParam "bad parameters"
// @Failure 500 {object} docs.ResponseServerErr "internal server error"
// @Router /user/{id} [get]
func (uv *userView) User(c *gin.Context) {
	paramErr := verify.NewParamErr()
	id := paramErr.CheckId(c.Param("id"))
	if !paramErr.ParamValid() {
		RespErr(c, paramErr.Error())
		return
	}
	data, err := uv.us.Get(int(id))
	if err != nil {
		RespErr(c, err)
		return
	}
	RespSuccess(c, data)
}

// @Summary 增加用户
// @Description 增加用户
// @Tags user
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param agent_id formData int true "代理人ID"
// @Param name formData string true "姓名"
// @Param age formData int true "年龄"
// @Success 200 {array} docs.ResponseOk "success"
// @Failure 400 {object} docs.ResponseInvalidParam "bad parameters"
// @Failure 500 {object} docs.ResponseServerErr "internal server error"
// @Router /user [post]
func (uv *userView) Add(c *gin.Context) {

}

// @Summary 更新用户信息
// @Description 更新用户信息
// @Tags user
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param id path int true "用户ID"
// @Param name formData string true "姓名"
// @Param email formData int true "年龄"
// @Success 200 {array} docs.ResponseOk "success"
// @Failure 400 {object} docs.ResponseInvalidParam "bad parameters"
// @Failure 500 {object} docs.ResponseServerErr "internal server error"
// @Router /user/{id} [put]
func (uv *userView) Update(c *gin.Context) {

}

// @Summary 删除用户
// @Description 根据用户ID,删除用户
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {array} docs.ResponseOk "success"
// @Failure 400 {object} docs.ResponseInvalidParam "bad parameters"
// @Failure 500 {object} docs.ResponseServerErr "internal server error"
// @Router /user/{id} [delete]
func(uv *userView)Delete(c *gin.Context){

}