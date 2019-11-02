package verify

import (
	"fmt"
	"nomal-gin/services"
	"strconv"
	"strings"
)

type paramErr struct {
	we *services.WarpErr
}

func (p *paramErr) Error() *services.WarpErr {
	return p.we
}

func NewParamErr() *paramErr {
	return &paramErr{}
}

func (p *paramErr) ParamValid() bool {
	if p.we != nil {
		return false
	}
	return true
}

func (p *paramErr) CheckName(name string) string {
	if !p.ParamValid() {
		return ""
	}
	name = strings.TrimSpace(name)
	l := len(name)
	if l < 3 || l > 10 {
		p.we = services.NewWarpErr(services.ErrParamsNotValid, nil)
	}
	return name
}

func (p *paramErr) String2Int64(str string) int64 {
	if !p.ParamValid() {
		return 0
	}
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		p.we = services.NewWarpErr(services.ErrParamsNotValid, err)
	}
	return i
}

func (p *paramErr) String2TimeInt64(str string) int64 {
	if !p.ParamValid() {
		return 0
	}
	if str == "" {
		return 0
	}
	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		p.we = services.NewWarpErr(services.ErrParamsNotValid, err)
	}
	return i
}

func (p *paramErr) String2Float64(str string) float64 {
	if !p.ParamValid() {
		return 0
	}
	i, err := strconv.ParseFloat(str, 64)
	if err != nil {
		p.we = services.NewWarpErr(services.ErrParamsNotValid, err)
	}
	return i
}

func (p *paramErr) CheckStartID(start string) int64 {
	if !p.ParamValid() {
		return 0
	}
	s := p.String2Int64(start)
	if !p.ParamValid() {
		return 0
	}
	if s < 0 {
		p.we = services.NewWarpErr(services.ErrParamsNotValid, nil)
	}
	return s
}

func (p *paramErr) CheckLimit(limit string) int64 {
	if !p.ParamValid() {
		return 0
	}
	l := p.String2Int64(limit)
	if !p.ParamValid() {
		return 0
	}
	if l < 0 || l > 50 {
		p.we = services.NewWarpErr(services.ErrParamsNotValid, nil)
	}
	return l
}

func (p *paramErr) CheckLimitInt(limit int64) int64 {
	if !p.ParamValid() {
		return 0
	}
	if limit < 0 || limit > 50 {
		p.we = services.NewWarpErr(services.ErrParamsNotValid, nil)
	}
	return limit
}

func (p *paramErr) CheckPage(page string) int64 {
	// page参数最小为1
	if !p.ParamValid() {
		return 0
	}
	pg := p.String2Int64(page)
	if pg < 1 {
		p.we = services.NewWarpErr(services.ErrParamsNotValid, nil)
	}
	return pg
}

func (p *paramErr) CheckOffset(offset string) int64 {
	if !p.ParamValid() {
		return 0
	}
	o := p.String2Int64(offset)
	if !p.ParamValid() {
		return 0
	}
	if o < 0 {
		p.we = services.NewWarpErr(services.ErrParamsNotValid, nil)
	}
	return o
}

func (p *paramErr) CheckId(id string) int64 {
	if !p.ParamValid() {
		return 0
	}
	o := p.String2Int64(id)
	if o <= 0 {
		p.we = services.NewWarpErr(services.ErrParamsNotValid, nil)
	}
	return o
}

func (p *paramErr) CheckJsonId(id int64) int64 {
	if !p.ParamValid() {
		return 0
	}
	if id <= 0 {
		p.we = services.NewWarpErr(services.ErrParamsNotValid, fmt.Errorf("id not valid "))
	}
	return id
}

func (p *paramErr) CheckJsonEnumerate(enumerate uint8) uint8 {
	if !p.ParamValid() {
		return 0
	}
	if enumerate <= 0 {
		p.we = services.NewWarpErr(services.ErrParamsNotValid, fmt.Errorf("enumerate not valid "))
	}
	return enumerate
}
