package models

import (
	"github.com/uptrace/bun"
	"time"
)

type Example struct {
	bun.BaseModel `bun:"table:example,alias:e"`
	ID            int          `json:"id" bun:"id,pk,autoincrement"`
	Value         int          `json:"value" bun:"value,notnull" `
	CreatedAt     time.Time    `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt     bun.NullTime `json:"updated_at"`
}
