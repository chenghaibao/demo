package carbon

import (
	"strings"
	"time"
)

// Season gets season name according to the meteorological division method, i18n is supported.
// 获取当前季节(以气象划分)，支持i18n
func (c Carbon) Season() string {
	if c.IsInvalid() {
		return ""
	}
	if len(c.lang.resources) == 0 {
		c.lang.SetLocale(defaultLocale)
	}
	index := -1
	switch {
	case c.Month() == 3 || c.Month() == 4 || c.Month() == 5:
		index = 0
	case c.Month() == 6 || c.Month() == 7 || c.Month() == 8:
		index = 1
	case c.Month() == 9 || c.Month() == 10 || c.Month() == 11:
		index = 2
	case c.Month() == 12 || c.Month() == 1 || c.Month() == 2:
		index = 3
	}
	if seasons, ok := c.lang.resources["seasons"]; ok {
		slice := strings.Split(seasons, "|")
		if len(slice) == 4 {
			return slice[index]
		}
	}
	return ""
}

// StartOfSeason returns a Carbon instance for start of the season.
// 本季节开始时间
func (c Carbon) StartOfSeason() Carbon {
	if c.IsInvalid() {
		return c
	}
	if c.Month() == 1 || c.Month() == 2 {
		c.Time = time.Date(c.Year()-1, time.Month(12), 1, 0, 0, 0, 0, c.loc)
		return c
	}
	c.Time = time.Date(c.Year(), time.Month(c.Month()/3*3), 1, 0, 0, 0, 0, c.loc)
	return c
}

// EndOfSeason returns a Carbon instance for end of the season.
// 本季节结束时间
func (c Carbon) EndOfSeason() Carbon {
	if c.IsInvalid() {
		return c
	}
	if c.Month() == 1 || c.Month() == 2 {
		c.Time = time.Date(c.Year(), time.Month(2), 1, 23, 59, 59, 999999999, c.loc).AddDate(0, 1, -1)
		return c
	}
	if c.Month() == 12 {
		c.Time = time.Date(c.Year()+1, time.Month(2), 1, 23, 59, 59, 999999999, c.loc).AddDate(0, 1, -1)
		return c
	}
	c.Time = time.Date(c.Year(), time.Month(c.Month()/3*3+2), 1, 23, 59, 59, 999999999, c.loc).AddDate(0, 1, -1)
	return c
}

// IsSpring whether is spring.
// 是否是春季
func (c Carbon) IsSpring() bool {
	if c.IsInvalid() {
		return false
	}
	if c.Month() == 3 || c.Month() == 4 || c.Month() == 5 {
		return true
	}
	return false
}

// IsSummer whether is summer.
// 是否是夏季
func (c Carbon) IsSummer() bool {
	if c.IsInvalid() {
		return false
	}
	if c.Month() == 6 || c.Month() == 7 || c.Month() == 8 {
		return true
	}
	return false
}

// IsAutumn whether is autumn.
// 是否是秋季
func (c Carbon) IsAutumn() bool {
	if c.IsInvalid() {
		return false
	}
	if c.Month() == 9 || c.Month() == 10 || c.Month() == 11 {
		return true
	}
	return false
}

// IsWinter whether is winter.
// 是否是冬季
func (c Carbon) IsWinter() bool {
	if c.IsInvalid() {
		return false
	}
	if c.Month() == 12 || c.Month() == 1 || c.Month() == 2 {
		return true
	}
	return false
}
