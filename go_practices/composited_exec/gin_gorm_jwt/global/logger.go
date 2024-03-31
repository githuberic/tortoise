package global

import (
	"go.uber.org/zap"
	"go_practices/composited_exec/gin_gorm_jwt/pkg/zaplog"
)

var (
	Logger *zap.SugaredLogger
)

//创建logger
func SetupLogger() error {
	var err error
	filepath := LogSetting.LogFilePath
	infofilename := LogSetting.LogInfoFileName
	warnfilename := LogSetting.LogWarnFileName
	fileext := LogSetting.LogFileExt

	//infofilename,warnfilename,fileext string
	Logger, err = zaplog.GetInitLogger(filepath, infofilename, warnfilename, fileext)

	if err != nil {
		return err
	}
	defer Logger.Sync()
	return nil
}
