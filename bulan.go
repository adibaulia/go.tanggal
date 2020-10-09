package bulan

import (
	"strings"
	"time"
)

type (
	bulan string
)

var datetime time.Time = time.Now()

//NewBulan creates initialization time for bulan
func NewBulan(time time.Time) bulan {
	datetime = time
	return bulan("")
}

//Long returns the Indonesia name of "bulan" in Long and Title Case
func (b bulan) Long() bulan {
	switch datetime.Month() {
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

//Short returns the Indonesia name of "bulan" in Short and Title Case
func (b bulan) Short() bulan {
	return b.Long()[:3]
}

//ToString returns the Indonesia name of "bulan" in to string type
func (b bulan) ToString() string {
	return string(b)
}

//LowerCase returns lowercase of "bulan"
func (b bulan) LowerCase() bulan {
	return bulan(strings.ToLower(b.ToString()))
}

//UpperCase returns UPPERCASE of "bulan"
func (b bulan) UpperCase() bulan {
	return bulan(strings.ToUpper(b.ToString()))
}
