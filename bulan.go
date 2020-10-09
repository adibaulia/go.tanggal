/*
Package bulan implements a simple library for parsing to Indonesia name of Month.
*/
package bulan

import (
	"time"
)

type (
	//Bulan is type aliases for having some functions
	Bulan string
)

var datetime time.Time = time.Now()

//NewBulan creates initialization time for bulan
func NewBulan(time time.Time) Bulan {
	datetime = time
	return Bulan("")
}

//Long returns the Indonesia name of "bulan" in Long and Title Case
func (b Bulan) Long() Bulan {
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
func (b Bulan) Short() Bulan {
	return b.Long()[:3]
}

//ToString returns the Indonesia name of "bulan" in to string type
func (b Bulan) ToString() string {
	return string(b)
}
