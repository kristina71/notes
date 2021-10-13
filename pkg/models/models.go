package models

import (
	"time"
)

type Note struct {
	Id          uint16    `json:"id" db:"id"`
	Text        string    `json:"text" db:"text"`
	CreatedTime time.Time `json:"created_at" db:"created_time"`
	UpdatedTime  time.Time `json:"updated_at" db:"updated_time"`
}
