package global

import (
	"github.com/iBaiYang/go-gin-blog-eddycjy/pkg/logger"
	"github.com/iBaiYang/go-gin-blog-eddycjy/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
