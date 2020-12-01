package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

/**
Celsius 参数c出现在函数名称前面，名称叫string的方法关联到Celsius类型，返回c变量的数值型
 */
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
