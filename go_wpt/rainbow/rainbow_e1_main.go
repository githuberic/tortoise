package main

import (
	"errors"
	"flag"
	"fmt"
	"sync/atomic"
	"time"
)

var (
	readNum       int64
	writeFail     bool
	exitLoop      bool
	extraSlowRead bool
)

type stream struct {
	stopRead  chan struct{}
	stopWrite chan struct{}
	responseC chan string
}

func init() {
	flag.BoolVar(&writeFail, "wf", false, "write fail")
	flag.BoolVar(&exitLoop, "el", false, "exit loop")
	flag.BoolVar(&extraSlowRead, "sr", false, "extra slow read")
	flag.Parse()
}

func (s *stream) handle() {
	defer func() {
		close(s.stopRead)
		fmt.Println("处理器完成stopRead\n")
	}()

	go s.loopRead()

	for {
		select {
		case <-s.stopWrite:
			fmt.Println("处理器收到stopWrite\n")
		case response, ok := <-s.responseC:
			if !ok {
				continue
			}
			fmt.Printf("处理器收到处理后响应 [%s]\n\n", response)
			if err := s.write(response); err != nil {
				fmt.Printf("rainbow写异常, err [%s]\n\n", err.Error())
				return
			}
		}
	}
}

func (s *stream) loopRead() {
	defer func() {
		close(s.stopWrite)
		fmt.Println("loopRead完成stopWrite\n")
	}()

	for {
		select {
		case <-s.stopRead:
			fmt.Println("loopRead收到stopRead信号\n")
			if exitLoop {
				return
			}
			break
		default:
			payload, err := s.reader()
			if err != nil {
				fmt.Printf("loopRead读取异常 err [%s]\n\n", err.Error())
				return
			}
			fmt.Printf("loopRead读取内容 [%d], err [%v]\n\n", payload, err)
			s.responseC <- s.internalProcess(payload)
		}
	}
}

func (s *stream) write(resp string) error {
	if writeFail {
		return errors.New("写包失败")
	}
	return nil
}

func (s *stream) reader() (int64, error) {
	v := atomic.AddInt64(&readNum, 1)
	if v == 1 {
		return time.Now().UnixNano(), nil
	}
	if extraSlowRead {
		time.Sleep(time.Millisecond * 10)
		fmt.Println("读包延迟\n")
		return time.Now().UnixNano(), nil
	}
	return 0, errors.New("读包失败")
}

func (s *stream) internalProcess(payload int64) string {
	return fmt.Sprintf("我是逻辑处理后的结果, %d", payload)
}

func main() {
	s := stream{stopRead: make(chan struct{}), stopWrite: make(chan struct{}), responseC: make(chan string, 10)}
	s.handle()

	time.Sleep(3 * time.Second)
	fmt.Println("正常完成请求,断开链接")
}
