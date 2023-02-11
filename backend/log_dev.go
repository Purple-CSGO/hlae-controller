//go:build !build
// +build !build

package backend

import (
	"hlae-controller/backend/tool"
	"os"
	"path"
	"runtime"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

// 初始化日志
func (a *App) initLog() *os.File {
	var logDir string
	logName := "[debug]" + time.Now().Format("2006年01月02日_15点04分_05秒_") + strconv.FormatInt(time.Now().Unix(), 10) + ".log"
	logDir = "./log"
	logFile := path.Join(logDir, logName)

	if !tool.IsFileExisted(logDir) {
		if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
			log.Println(err)
			return nil
		}
	}

	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	log.SetOutput(f)
	log.SetLevel(log.DebugLevel)

	// 输出系统信息
	log.Infof("[%v] %v - %v @ %v/%v", a.AppDeveloper, a.AppName, a.Version, runtime.GOOS, runtime.GOARCH)
	log.Info("日志初始化完成")
	log.Debug("Debug输出测试")

	return f
}

// 关闭日志
func (a *App) closeLog(f *os.File) {
	defer f.Close()
	log.Info("日志已关闭")
}
