package ioc

import (
	"fmt"
	"testing"
)

//-------------------
//抽象层
//-------------------
type Car interface {
	Run()
}
type Driver interface {
	Drive(car Car)
}

type Benz struct{}
func (b *Benz) Run() {
	fmt.Println("Mercedes-Benz")
}

type BMW struct{}
func (b *BMW) Run() {
	fmt.Println("BMW")
}

type Toyota struct{}
func (f *Toyota) Run() {
	fmt.Println("Toyota")
}

//driver
type ZhangSan struct{}
func (zs *ZhangSan) Drive(car Car) {
	car.Run()
}

type LiSi struct {}
func (ls *LiSi)Drive(car Car){
	car.Run()
}

type WanWu struct {}
func (ww *WanWu)Drive(car Car){
	car.Run()
}

func TestVerifyV1(t *testing.T) {
	// 业务1：ZhangSan开奔驰
	var zhangSan ZhangSan
	var car Car
	car = &Toyota{} //多态
	zhangSan.Drive(car)
}
