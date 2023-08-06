package service

import (
	"context"

	"github.com/iBaiYang/go-gin-blog-eddycjy/global"
	"github.com/iBaiYang/go-gin-blog-eddycjy/internal/dao"
	//otgorm "github.com/eddycjy/opentracing-gorm"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)

	// SQL追踪
	//svc.dao = dao.New(otgorm.WithContext(svc.ctx, global.DBEngine))

	return svc
}
