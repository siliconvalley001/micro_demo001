package service

import (
	"context"
	"github.com/siliconvalley001/micro_demo001/user/dao"
)

type Service_User struct {
	ctx *context.Context
	dao *dao.Dao
}
