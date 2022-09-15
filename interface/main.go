package main

import "fmt"

// interface 被賦值後，它的型別值會變成實作它的 Type 的型別和值
type Shape interface {
	Area() float64
}

// Rect 實作了 Shape interface
// Rect 所建立的變數同時會符合 Rect（Struct Type）和 Shape（Interface Type）
type Rect struct {
	width  float64
	height float64
}

func (r Rect) Area() float64 {
	return r.height * r.width
}

func main() {
	var s Shape = Rect{3, 5}
	fmt.Printf("(%T, %v) \n", s, s) // (main.Rect, {3 5})
	fmt.Println(s.Area())           // 可以直接用 Shape interface 來呼叫方法
}
