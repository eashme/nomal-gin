package views

import (
	"nomal-gin/services"
	"nomal-gin/utils/errlog"
)

// 转换返回数据
func convert2Resp(data interface{}) interface{} {
	switch data.(type) {
	case []map[string]interface{}, map[string]interface{}, map[string]string, nil:
		return data
	case services.IPageNator:
		return respPageNator(data.(services.IPageNator))
	case services.IUser:
		return respUser(data.(services.IUser))
	case []services.IUser:
		return respUsers(data.([]services.IUser))
	default:
		errlog.Errorf("not have this type %+v", data)
		return data
	}
}

func respPageNator(val services.IPageNator) interface{} {
	type u struct {
		Page      int64       `json:"page"`
		Data      interface{} `json:"data"`
		TotalRow  int64       `json:"total_row"`
		TotalPage int64       `json:"total_page"`
		Limit     int64       `json:"limit"`
	}
	return &u{
		val.Page(),
		convert2Resp(val.Data()),
		val.TotalRow(),
		val.TotalPage(),
		val.Limit(),
	}
}

func respUser(val services.IUser) interface{} {
	type u struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	return &u{
		val.ID(),
		val.Name(),
		val.Age(),
	}
}

func respUsers(val []services.IUser) interface{} {
	var res []interface{}
	for _, u := range val {
		res = append(res, respUser(u))
	}
	return res
}
