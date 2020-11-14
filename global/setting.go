package global

import (
	"github.com/KarasWinds/blog-service/pkg/logger"
	"github.com/KarasWinds/blog-service/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
