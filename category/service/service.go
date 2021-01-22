package service

import (
	"category/dao"
	"context"
)

type Service_Category struct {
	ctx *context.Context
	dao *dao.Dao
}
