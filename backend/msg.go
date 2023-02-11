package backend

import (
	log "github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// 通知消息
func (a *App) info(title string, content string) {
	log.Info(title, content)
	runtime.EventsEmit(a.ctx, "info", title, content)
}

func (a *App) success(title string, content string) {
	log.Info("success: ", title, content)
	runtime.EventsEmit(a.ctx, "success", title, content)
}

func (a *App) warning(title string, content string) {
	log.Warning(title, content)
	runtime.EventsEmit(a.ctx, "warning", title, content)
}

func (a *App) error(title string, content string) {
	log.Error(title, content)
	runtime.EventsEmit(a.ctx, "error", title, content)
}
