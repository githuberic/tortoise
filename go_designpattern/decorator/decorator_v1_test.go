package decorator

import (
	"fmt"
	"testing"
)

type Shape interface {
	Draw()
}

type Rectangle struct {
}

func (r *Rectangle) Draw() {
	fmt.Println("Shape: Rectangle")
}

type Circle struct {
}

func (c *Circle) Draw() {
	fmt.Println("Shape: Circle")
}

// 定义抽象装饰类
type ShapeDecorator struct {
	decoratedShape Shape
}

func (s *ShapeDecorator) ShapeDecorator(decoratedShape Shape) {
	s.decoratedShape = decoratedShape
}
func (s *ShapeDecorator) Draw() {
	s.decoratedShape.Draw()
}

// 定义具体的装饰类
type RedShapeDecorator struct {
	shapeDecorator ShapeDecorator
}

func (s *RedShapeDecorator) RedShapeDecorator(decoratedShape Shape) {
	s.shapeDecorator.ShapeDecorator(decoratedShape)
}
func (s *RedShapeDecorator) Draw() {
	s.shapeDecorator.Draw()
	s.setRedBorder(s.shapeDecorator.decoratedShape)
}
func (s *RedShapeDecorator) setRedBorder(decoratedShape Shape) {
	fmt.Println("Border Color: Red")
}

func TestVerify(t *testing.T) {
	var rectangle = Rectangle{}
	rectangle.Draw()

	var circle = Circle{}
	circle.Draw()

	var shape Shape = new(Circle)
	// 增加Draw其他颜色属性
	var shapeDecorator ShapeDecorator
	shapeDecorator.decoratedShape = shape
	shapeDecorator.Draw()

	var redShapeDecorator RedShapeDecorator
	redShapeDecorator.RedShapeDecorator(shape)
	redShapeDecorator.Draw()
}
