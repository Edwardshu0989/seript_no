package logger

import (
	"github.com/yeeyuntech/yeego"
	"github.com/yeeyuntech/yeego/yeeFile"
	"gitlab.yeeyuntech.com/yee/easyweb"
	"gitlab.yeeyuntech.com/yee/easyweb/log4go"
)

type LogConfig struct {
	WarnPath  string
	DebugPath string
	ErrorPath string
}

func InitLogger(cfg LogConfig) {
	runMode := yeego.Config.GetString("app.RunMode")
	Logger := log4go.NewLogger()
	if runMode == "debug" {
		Logger.AddFilter("stdout_warn", log4go.WARNING, log4go.NewConsoleLogWriter())
		Logger.AddFilter("stdout_debug", log4go.DEBUG, log4go.NewConsoleLogWriter())
		Logger.AddFilter("stdout_err", log4go.ERROR, log4go.NewConsoleLogWriter())
	} else {
		var err error
		if cfg.WarnPath != "" {
			err = yeeFile.MkdirForFile(cfg.WarnPath)
			if err != nil {
				panic("can not create file dir " + cfg.WarnPath)
			}
			flwWarn := log4go.NewFileLogWriter(cfg.WarnPath, true)
			flwWarn.SetFormat("[%D %T] [%L] (%S) %M")
			flwWarn.SetRotateDaily(true)
			flwWarn.SetRotateSize(1024 * 1024 * 500)
			Logger.AddFilter("warn_file", log4go.WARNING, flwWarn)
		}
		err = yeeFile.MkdirForFile(cfg.DebugPath)
		if err != nil {
			panic("can not create file dir " + cfg.DebugPath)
		}
		err = yeeFile.MkdirForFile(cfg.ErrorPath)
		if err != nil {
			panic("can not create file dir " + cfg.ErrorPath)
		}

		flwDebug := log4go.NewFileLogWriter(cfg.DebugPath, true)
		flwDebug.SetFormat("[%D %T] [%L] (%S) %M")
		flwDebug.SetRotateDaily(true)
		flwDebug.SetRotateMaxBackup(10)
		flwDebug.SetRotateSize(1024 * 1024 * 500)
		Logger.AddFilter("debug_file", log4go.DEBUG, flwDebug)

		flwError := log4go.NewFileLogWriter(cfg.ErrorPath, true)
		flwError.SetFormat("[%D %T] [%L] (%S) %M")
		flwError.SetRotateDaily(true)
		flwError.SetRotateMaxBackup(10)
		flwError.SetRotateSize(1024 * 1024 * 500)
		Logger.AddFilter("error_file", log4go.ERROR, flwError)
	}
	easyweb.Logger = Logger
}
