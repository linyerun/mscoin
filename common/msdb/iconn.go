package msdb

import (
	"context"
	"gorm.io/gorm"
)

type IConnection interface {
	Begin()
	Rollback()
	Commit()
	Session(ctx context.Context) *gorm.DB
	Tx(ctx context.Context) *gorm.DB
}
