/*
小数计算工具包 , 使用decimal进行精确小数计算

*/
package utils

import (
	"github.com/shopspring/decimal"
	"math"
)
// 判断有无小数，有的话向上取整
func UpToInteger(n float64) float64 {
	return math.Ceil(n)
}

// 四舍五入
func RoundOff(n float64) float64 {
	return math.Floor(n+0.5)
}

func Add(a, b float64) float64 {
	val, _ := decimal.NewFromFloat(a).Add(decimal.NewFromFloat(b)).Float64()
	return val
}

func Sub(a, b float64) float64 {
	val, _ := decimal.NewFromFloat(a).Sub(decimal.NewFromFloat(b)).Float64()
	return val
}

func Mul(a, b float64) float64 {
	val, _ := decimal.NewFromFloat(a).Mul(decimal.NewFromFloat(b)).Float64()
	return val
}

func Div(a, b float64) float64 {
	val, _ := decimal.NewFromFloat(a).Div(decimal.NewFromFloat(b)).Float64()
	return val
}
