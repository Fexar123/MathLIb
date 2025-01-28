package main

import (
	"MathLib/Math"
	"fmt"
)

func main() {

	var com1 Math.ComplexNumber
	com1.SetRe(5)
	com1.SetIm(6)
	var com2 Math.ComplexNumber
	com2.SetRe(20)
	com2.SetIm(10)

	result := Math.SumComplexNumber(com1, com2)

	fmt.Println(result.GetRe())
	fmt.Println(result.GetIm())

	result = Math.MultComplexNumber(com1, com2)

	fmt.Println(result.GetRe())
	fmt.Println(result.GetIm())

	result = Math.DivComplexNumber(com1, com2)

	fmt.Println(result.GetRe())
	fmt.Println(result.GetIm())

	com1.ConstMult(5)
	fmt.Println(com1.GetRe())
	fmt.Println(com1.GetIm())
}
