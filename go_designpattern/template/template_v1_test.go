package template

import (
	"fmt"
	"testing"
)

type Game struct {
	Initialize func()
	StartPlay  func()
	EndPlay    func()
}

//Play 模板基类的Play方法
func (p *Game) Play() {
	p.Initialize()
	p.StartPlay()
	p.EndPlay()
}

type FootBall struct {
	Game
}

func NewFootBall() *FootBall {
	ft := new(FootBall)
	ft.Game.Initialize = ft.Initialize
	ft.Game.StartPlay = ft.StartPlay
	ft.Game.EndPlay = ft.EndPlay
	return ft
}

func (p *FootBall) Initialize()  {
	fmt.Println("Football game initialize")
}

//StartPlay 子类的StartPlay方法
func (p *FootBall) StartPlay() {
	fmt.Println("Football game started.")
}

//EndPlay 子类的EndPlay方法
func (p *FootBall) EndPlay() {
	fmt.Println("Football game Finished!")
}


type Basketball struct {
	Game
}

func NewBasketball() *Basketball {
	ft := new(Basketball)
	ft.Game.Initialize = ft.Initialize
	ft.Game.StartPlay = ft.StartPlay
	ft.Game.EndPlay = ft.EndPlay
	return ft
}

func (p *Basketball) Initialize()  {
	fmt.Println("Basketball game initialize")
}

//StartPlay 子类的StartPlay方法
func (p *Basketball) StartPlay() {
	fmt.Println("Basketball game started.")
}

//EndPlay 子类的EndPlay方法
func (p *Basketball) EndPlay() {
	fmt.Println("Basketball game Finished!")
}

func TestVerify(t *testing.T)  {
	football := NewFootBall()
	football.Play()

	fmt.Println("-------------------")
	basketball := NewBasketball()
	basketball.Play()
}