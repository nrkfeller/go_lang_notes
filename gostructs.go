package main

import "fmt"

func main(){

	rect1 := Rectangle{leftX:0 , topy: 50, height:10, width:10}
	rect2 := Rectangle{30,60,70,70}

	fmt.Println(rect1.topy)

	fmt.Println("rect area is ", rect1.area())

}

type Rectangle struct {
	leftX float32
	topy float32
	height float32
	width float32

}

// class method type thing
func (rect *Rectangle) area() float32 {
	return rect.width * rect.height
}