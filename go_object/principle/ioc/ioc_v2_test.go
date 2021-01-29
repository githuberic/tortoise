package ioc

import (
	"fmt"
	"testing"
)

//------------抽象层------------
type Card interface {
	Display()
}
type Memory interface {
	Storage()
}
type Cpu interface {
	Calculate()
}

type Computer struct {
	ItemCard   Card
	ItemMemory Memory
	ItemCpu    Cpu
}

//NewComputer  初始化一个computer类对象
func NewComputer(cpu Cpu, card Card, memory Memory) (c *Computer) {
	c = &Computer{ //多态性
		ItemCard:   card,
		ItemCpu:    cpu,
		ItemMemory: memory,
	}
	return c
}

func (c *Computer) Work(){
	c.ItemCard.Display()
	c.ItemMemory.Storage()
	c.ItemCpu.Calculate()
}

//------------实现层------------
type IntelCpu struct {
}
func (i *IntelCpu) Calculate(){
	fmt.Println("IntelCpu")
}

type IntelMemory struct {
}
func (i *IntelMemory)Storage(){
	fmt.Println("IntelMemory")
}

type KingMemory struct {
}
func (k *KingMemory)Storage(){
	fmt.Println("KingMemory")
}


type IntelCard struct {
}
func (i *IntelCard) Display(){
	fmt.Println("IntelCard")
}
type NVIDCard struct {
}
func (n *NVIDCard)Display(){
	fmt.Println("NVIDCard")
}


func TestVerifyV2(t *testing.T) {
	//组装Intel系列的电脑 并运行
	intelComputer:=NewComputer(&IntelCpu{},&IntelCard{},&IntelMemory{})
	intelComputer.Work()

	fmt.Printf("\n\n")
	//组装Intel Cpu、Kingston Memory、NVIDIA Card混合的电脑并运行
	mixComputer:=NewComputer(&IntelCpu{},&NVIDCard{},&KingMemory{})
	mixComputer.Work()
}

// https://blog.csdn.net/wzb_wzt/article/details/107380699
