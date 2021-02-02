package facade

import (
	"fmt"
	"testing"
)

// CPU
type CPU struct {
}
func (p *CPU) start() {
	fmt.Println("启动CPU...")
}

// 内存
type Memory struct {
}
func (p *Memory) start() {
	fmt.Println("启动内存管理...")
}

// 硬盘
type Disk struct {
}
func (p *Disk) start() {
	fmt.Println("启动硬盘...")
}

// 开机键
type StartBtn struct {
}
func (StartBtn) start() {
	cpu := &CPU{}
	cpu.start()

	memory := &Memory{}
	memory.start()

	disk := &Disk{}
	disk.start()
}

func TestVerify(t *testing.T)  {
	startBtn := &StartBtn{}
	startBtn.start()
}
// https://learnku.com/articles/33705