package services

// 分页接口
type IPageNator interface {
	Page() int64
	Data() interface{}
	TotalRow() int64
	TotalPage() int64
	Limit() int64
}