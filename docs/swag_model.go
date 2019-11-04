/*
此文档只用于构建swagger API
*/
package docs

type ResponseOk struct {
	Code int    `json:"code" example:"200"`
	Msg  string `json:"message" example:"ok"`
	Data string `json:"data" example:"success"`
}

type ResponseInvalidParam struct {
	Code int    `json:"code" example:"400"`
	Msg  string `json:"message" example:"param is not valid"`
	Data string `json:"data" example:""`
}

type ResponseServerErr struct {
	Code int    `json:"code" example:"500"`
	Msg  string `json:"message" example:"internal server error"`
	Data string `json:"data" example:""`
}

type User struct {
	Name string `json:"name"`
	Age uint8 `json:"age"`
}

type ResponseUser struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data User `json:"data"`
}

type ResponseUsers struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data []User `json:"data"`
}