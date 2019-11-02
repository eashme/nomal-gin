package services

import "nomal-gin/utils/errlog"

type respCode uint16

const (
	Success respCode = 200
)

//  400+ 异常
const (
	ErrParamsNotValid respCode = 401 + iota
	ErrUserForbidden
	ErrDataNotFound
	ErrMethodNotAllowed
)

// 500+ 异常
const (
	ErrInternalServer = 500 + iota
	ErrServerNotImplement
	ErrBadGetWay
	ErrServiceUnavailable
)

var httpRespInfo = map[respCode]string{
	ErrParamsNotValid: "Parameters Error !",
	ErrUserForbidden: "Forbidden !",
	ErrDataNotFound:"Not Found !",
	ErrMethodNotAllowed:"Method Not Allowed !",
	ErrInternalServer:"Internal Server Error !",
	ErrServerNotImplement:"Server Not Implement",
	ErrBadGetWay: "Bad Gateway !",
	ErrServiceUnavailable:"Server Unavailable !",
}


// service层异常封装
type WarpErr struct {
	code respCode
	err  error
}

func NewWarpErr(code respCode, err error) *WarpErr {
	return &WarpErr{
		code: code,
		err:  err,
	}
}


func (w *WarpErr) String() string {
	if info, ok := httpRespInfo[w.code]; ok {
		return info
	}
	errlog.Debug("not has the respCode:", w.code)
	return ""
}

func (w *WarpErr) Code() int {
	return int(w.code)
}

func (w *WarpErr) Err() error {
	return w.err
}