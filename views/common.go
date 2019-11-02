package views

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nomal-gin/services"
)

type respData struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

func newRespData(code int, msg string, data interface{}) *respData {
	return &respData{Code: code, Msg: msg, Data: data}
}

func RespSuccess(c *gin.Context, data interface{}) {
	resp := newRespData(http.StatusOK, "success", convert2Resp(data))
	c.JSON(http.StatusOK, resp)
}

func RespErr(c *gin.Context, err *services.WarpErr) {
	if err == nil {
		c.JSON(http.StatusInternalServerError, newRespData(http.StatusInternalServerError, "internal error", nil))
		return
	}
	code := http.StatusInternalServerError
	if err.Code() >= http.StatusBadRequest && err.Code() < http.StatusInternalServerError {
		code = http.StatusBadRequest
	}
	c.JSON(code, newRespData(err.Code(), err.String(), err.Err()))
}
