package bulan_test

import (
	"testing"
	"time"

	tanggal "github.com/adibaulia/go.tanggal"
	"github.com/stretchr/testify/assert"
)

const ISO_LAYOUT string = "2006-01-02"
const DATE_IN_STRING string = "2000-12-15"

var date time.Time

func init() {
	dateInDate, _ := time.Parse(ISO_LAYOUT, DATE_IN_STRING)
	date = dateInDate
}

func Test_InitNewBulan(t *testing.T) {
	bulan := tanggal.NewBulan(date)
	assert.Equal(t, "", string(bulan))
}

func Test_NewBulanToString(t *testing.T) {
	var bulan interface{} = tanggal.NewBulan(date).ToString()
	_, ok := bulan.(string)
	assert.Equal(t, true, ok)
}

func Test_NewBulanLong(t *testing.T) {
	bulan := tanggal.NewBulan(date).Long().ToString()
	assert.Equal(t, "Desember", bulan)
}

func Test_NewBulanShort(t *testing.T) {
	bulan := tanggal.NewBulan(date).Short().ToString()
	assert.Equal(t, "Des", bulan)
}
