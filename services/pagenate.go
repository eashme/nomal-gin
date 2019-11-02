package services

// 分页结构体
type PageNator struct {
	page     int64
	data     interface{}
	totalRow int64
	limit    int64
}

func NewPageNator(page int64, totalRow int64, limit int64,data interface{}) *PageNator {
	return &PageNator{
		page,
		data,
		totalRow,
		limit,
	}
}

func (p *PageNator) Page() int64 {
	return p.page
}

func (p *PageNator) Data() interface{} {
	return p.data
}

func (p *PageNator) TotalRow() int64 {
	return p.totalRow
}

func (p *PageNator) TotalPage() int64 {
	if p.limit == 0 {
		return p.totalRow
	}
	var totalPage int64
	totalPage = p.totalRow / p.limit
	if p.totalRow%p.limit > 0 {
		totalPage += 1
	}
	return totalPage
}

func (p *PageNator) Limit() int64 {
	return p.limit
}
