package models

import (
	"time"
)

type Warung struct {
	NamaWarung string
	Tanggal    time.Time
	Kasir      string
	Products   []Product
}
