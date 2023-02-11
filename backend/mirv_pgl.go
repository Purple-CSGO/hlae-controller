package backend

import (
	"fmt"
	mirvpgl "github.com/FlowingSPDG/HLAE-Server-GO"
	"github.com/hpcloud/tail"
	log "github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
	"time"
)

func (a *App) res(res string) {
	log.Println(res)
	runtime.EventsEmit(a.ctx, "res", res)
}

// ExampleHandler for HLAE Server
func ExampleHandler(cmd string) {
	log.Println("Received" + cmd)
	fmt.Printf("Received %s\n", cmd)
}

// ExampleCamHandler for cam datas
func ExampleCamHandler(cam *mirvpgl.CamData) {
	fmt.Printf("Received cam data %v\n", cam)
}

// ExampleEventHandler for cam datas
func ExampleEventHandler(ev *mirvpgl.GameEventData) {
	fmt.Printf("Received event data %v\n", ev)
}

func (a *App) StartServer(port string) (err error) {
	// New Instance
	config.Port = port
	a.hlaeserver, err = mirvpgl.New(":"+port, "/mirv")
	if err != nil {
		if err != nil {
			log.Error(err)
			a.error("服务启动失败", err.Error())
		}
		return err
	}

	a.hlaeserver.RegisterHandler(a.res)
	// a.hlaeserver.RegisterCamHandler(ExampleCamHandler)
	// a.hlaeserver.RegisterEventHandler(ExampleEventHandler)

	// Start Server
	go func() {
		err = a.hlaeserver.Start()
		if err != nil {
			log.Error(err)
			a.error("服务启动失败", err.Error())
		}
	}()

	return err
}

func (a *App) Send(cmd string) (err error) {
	log.Info("exec:", cmd)

	err = a.hlaeserver.BroadcastRCON(cmd)
	if err != nil {
		log.Error(err)
		a.error("指令执行失败", err.Error())
	}

	return
}

func (a *App) StartLog(csgoPath string, logfile string) (err error) {
	config.CsgoPath = csgoPath
	config.LogFile = logfile
	filePath := csgoPath + "\\" + logfile
	log.Info("startlog: ", filePath)

	config := tail.Config{
		ReOpen:    true,                                 // 重新打开
		Follow:    true,                                 // 是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件的哪个地方开始读
		MustExist: false,                                // 文件不存在不报错
		Poll:      true,                                 // 监听新行，使用tail -f，这个参数非常重要
	}

	tails, err := tail.TailFile(filePath, config)
	if err != nil {
		fmt.Println("tail file failed, err: ", err)
		log.Error(err)
		return err
	}

	var line *tail.Line
	var ok bool

	for {
		line, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}

		log.Info("getlog: ", line.Text)
		a.res(line.Text)
	}

	return
}

func (a *App) EndLog() (err error) {
	log.Info("endlog...")

	return
}

func (a *App) ClearLog(csgoPath string, logfile string) (err error) {
	config.CsgoPath = csgoPath
	config.LogFile = logfile
	filePath := csgoPath + "\\" + logfile
	log.Info("clearlog: ", filePath)

	return os.Remove(filePath)
}
