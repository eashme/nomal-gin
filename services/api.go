package services

// 分页接口
type IPageNator interface {
	Page() int64
	Data() interface{}
	TotalRow() int64
	TotalPage() int64
	Limit() int64
}

type IBaseService interface {
}

type IUserService interface {
	GetAll(page, pageSize int) (IPageNator, *WarpErr)
}

// 用户信息接口
type IUser interface {
	ID() int64
	Name() string
	Age() int
}
