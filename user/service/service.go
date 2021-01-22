package service

import (
	"context"
	"user/dao"
)

type Service_User struct {
	ctx *context.Context
	dao *dao.Dao
}
