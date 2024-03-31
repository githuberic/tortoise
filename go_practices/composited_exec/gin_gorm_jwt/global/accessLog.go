package global

import (
	"go.uber.org/zap"
	"go_practices/composited_exec/gin_gorm_jwt/pkg/zaplog"
)

var (
	AccessLogger *zap.SugaredLogger
)

//创建logger
func SetupAccessLogger() error {
	var err error
	filepath := AccessLogSetting.LogFilePath
	filename := AccessLogSetting.LogFileName
	//warnfilename:= LogSetting.LogWarnFileName
	fileext := AccessLogSetting.LogFileExt

	//infofilename,warnfilename,fileext string
	AccessLogger, err = zaplog.GetInitAccessLogger(filepath, filename, fileext)

	if err != nil {
		return err
	}
	defer AccessLogger.Sync()
	return nil
}
