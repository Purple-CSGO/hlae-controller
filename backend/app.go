package backend

import (
	"context"
	mirvpgl "github.com/FlowingSPDG/HLAE-Server-GO"
	log "github.com/sirupsen/logrus"
	"hlae-controller/backend/tool"
	"os"
)

// App struct
type App struct {
	ctx          context.Context // 上下文
	AppName      string
	AppDeveloper string
	Version      string
	logf         *os.File
	hlaeserver   *mirvpgl.HLAEServer
}

type Config struct {
	CsgoPath string `json:"csgo_path"`
	LogFile  string `json:"log_file"`
	Port     string `json:"port"`
}

var config Config

// NewApp 创建一个新的 App 应用程序
func NewApp(appName string, appDeveloper string, version string) *App {
	var a = App{AppName: appName, AppDeveloper: appDeveloper, Version: version}

	// 初始化日志
	a.logf = a.initLog()

	cfg, err := tool.ReadJson[Config]("./config.json", nil)
	if err != nil {
		log.Error("read config:", err)
	} else {
		config = cfg
	}

	return &a
}

// Startup 在应用程序启动时调用
func (a *App) Startup(ctx context.Context) {
	// 执行初始化设置
	a.ctx = ctx
	log.Info("后端准备完成")
}

// DomReady 在前端Dom加载完毕后调用
func (a *App) DomReady(ctx context.Context) {
	log.Info("前端DOM准备完成")
}

// Shutdown 在前端Dom销毁 应用程序终止时被调用
func (a *App) Shutdown(ctx context.Context) {
	// 在此处做一些资源释放的操作
	err := tool.SaveJson[Config](config, "./config.json", true, nil)
	if err != nil {
		log.Error("save config:", err)
	}
	a.closeLog(a.logf)
}

// BeforeClose 关闭应用程序之前回调
func (a *App) BeforeClose(ctx context.Context) bool {
	// 同步前端的设置
	log.Info(">>> Before Close")

	return false
}
