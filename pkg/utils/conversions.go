package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

type Float interface {
	float32 | float64
}

func StrToFloat(s string) (float64, error) {
	if s == "" {
		return 0, nil
	}
	str := strings.ReplaceAll(s, ",", ".")
	str = strings.ReplaceAll(str, " ", "")
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, err
	}
	return float64(f), nil
}

func FloatToStr[T Float](val T) string {
	str := strconv.FormatFloat(float64(val), 'f', -1, 64)
	return str
}

func FloatToStr2[T Float](val T) string {
	v := math.Round(float64(val)*100) / 100
	str := strconv.FormatFloat(v, 'f', 2, 64)
	return str
}

func FloatToStr2Comma[T Float](val T) string {
	str := FloatToStr2(val)
	str = strings.ReplaceAll(str, ".", ",")
	return str
}

func FloatToStr2Hyphen[T Float](val T) string {
	str := FloatToStr2(val)
	str = strings.ReplaceAll(str, ".", "-")
	return str
}

func MoneyToInteger(roubles float64) int {
	return int(math.Round(float64(roubles) * 100.00))
}

func MoneyToDecimal(kopeyki int) float64 {
	v := float64(kopeyki) / 100.00
	return float64(v)
}

func RoundMoneyUp(value float64) float64 {
	r := math.Ceil(float64(value)*100) / 100
	return float64(r)
}

func RoundMoneyDown(value float64) float64 {
	r := math.Floor(float64(value)*100) / 100
	return float64(r)
}

func RoundMoney(value float64) float64 {
	r := math.Round(float64(value)*100) / 100
	return r
}

func TimeToStr(t time.Time) string {
	str := fmt.Sprintf("%02d.%02d.%04d %02d:%02d", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute())
	return str
}

func StrToUint32(s string) (uint32, error) {
	if s == "" {
		return 0, nil
	}
	val, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(val), nil
}

func StrToInt(s string) (int, error) {
	if s == "" {
		return 0, nil
	}
	val, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return int(val), nil
}

const float64EqualityThreshold = 1e-9

func FloatAlmostEqual[T Float](a, b T) bool {
	return math.Abs(float64(a)-float64(b)) <= float64EqualityThreshold
}
