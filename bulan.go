package main

import (
	"strings"
	"time"
)

type (
	Bulan string
)

func ParseFullString(time time.Time) Bulan {
	switch time.Month() {
	case 1:
		return "Januari"
	case 2:
		return "Februari"
	case 3:
		return "Maret"
	case 4:
		return "April"
	case 5:
		return "Mei"
	case 6:
		return "Juni"
	case 7:
		return "Juli"
	case 8:
		return "Agustus"
	case 9:
		return "September"
	case 10:
		return "Oktober"
	case 11:
		return "November"
	case 12:
		return "Desember"
	}
	return ""
}

func (b Bulan) Short() Bulan {
	return b[:3]
}

func (b Bulan) ToString() string {
	return string(b)
}

func (b Bulan) AllLowerCase() Bulan {
	b = Bulan(strings.ToLower(b.ToString()))
	return b
}
