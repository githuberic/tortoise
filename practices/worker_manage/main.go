package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"tortoise/practices/worker_manage/work_manager"
)

type LoadWorker struct {
	RunningFlag bool
	ExitedFlag  bool
	ExitChan    chan struct{}
}

func NewLoadWorker() *LoadWorker {
	worker := &LoadWorker{RunningFlag: true, ExitedFlag: false}
	worker.ExitChan = make(chan struct{})
	return worker
}

func (worker *LoadWorker) Start() {
	log.Println("[start]LoadWorker")
	for worker.RunningFlag {
		select {
		case <-time.After(1 * time.Minute):
			//do some thing
			log.Println("LoadWorker do something")
			time.Sleep(time.Second * 3)

		case <-worker.ExitChan:
			log.Println("LoadWorker execute exit logic")
		}

	}
	worker.ExitedFlag = true
}

func (worker *LoadWorker) Stop() {
	log.Println("LoadWorker exit...")
	worker.RunningFlag = false
	close(worker.ExitChan)
	for !worker.ExitedFlag {
		time.Sleep(50 * time.Millisecond)
	}
	log.Println("[end]LoadWorker")
}

type WebServer struct {
	Server *http.Server
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (worker *WebServer) Start() {
	log.Println("[start]WebServer")

	ginHandler := gin.Default()
	ginHandler.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/plain", []byte("hello world!"))
	})

	worker.Server = &http.Server{
		Addr:           ":8080",
		Handler:        ginHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	worker.Server.ListenAndServe()
}

func (worker *WebServer) Stop() {
	log.Println("WebServer exit...")
	cxt, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// gracefull exit web server
	err := worker.Server.Shutdown(cxt)
	if err != nil {
		log.Printf("shutdown error, %v", err)
	}
	log.Println("[end]WebServer exit")
}

func prepareAllWorker() *work_manager.Master {
	m := work_manager.NewMaster()

	WorkerCount := 2
	for i := 0; i < WorkerCount; i++ {
		m.AddWorker(NewLoadWorker())
	}

	m.AddWorker(NewWebServer())
	return m
}

func GracefulExit(m *work_manager.Master) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch)

	for sig := range ch {
		switch sig {
		case syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT:
			log.Println("got a signal, execute stop", sig)
			close(ch)
			m.Stop()
		case syscall.SIGPIPE:
			log.Println("got a signal, ignore", sig)
		default:
			log.Println("got a signal, default", sig)
		}
	}
}

func main()  {
	// 1. init some worker
	wm := prepareAllWorker()

	// 2. start
	wm.Start()

	// 3. register grace exit
	GracefulExit(wm)

	// 4. block and wait
	wm.Wait()
}

// https://github.com/vearne/worker_manager/blob/master/README_zh.md