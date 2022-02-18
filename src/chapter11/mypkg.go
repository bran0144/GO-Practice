package chapter11

import "fmt"

type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParameter(float64)
	MethodWithReturnValue() string
}

type MyType int

func (m MyType) MethodWithoutParameters() {
	fmt.Println("MethodWithoutParameters called")
}
func (m MyType) MethodWithParameter(f float64) {
	fmt.Println("MethodWithParameter called")
}
func (m MyType) MethodWithReturnValue() string {
	return "Hi from MethodWithReturnValue"
}
func (my MyType) MethodNotInterface() {
	fmt.Println("MethodNotInterface called")
}
