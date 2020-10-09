# Go Tanggal

[![Documentation](https://godoc.org/github.com/adibaulia/go.tanggal?status.svg)](http://godoc.org/github.com/adibaulia/go.tanggal)

Go Tanggal merupakan package sederhana untuk mengubah nama waktu ke bahasa Indonesia (October -> Oktober)

<br>

# Cara Menggunakan

Silahkan import package menggunakan Go module:
```
go get github.com/adibaulia/go.tanggal
```

lalu Anda bisa menggunakan package ini seperti berikut:
```
package main

import (
	"fmt"
	"time"

	tanggal "github.com/adibaulia/go.tanggal"
)

func main() {
	fmt.Println(tanggal.NewBulan(time.Now()).Long()) //print Oktober
}
```

<br>

# Rencana Lebih Lanjut
Akan ada tambahan fitur untuk mengubah "day" ke bahasa Indonesia

<br>

# Kontribusi
Monggo yang mau ikut berkontribusi melalui pull request. Silahkan berikan kritik dan saran. Terimakasih!