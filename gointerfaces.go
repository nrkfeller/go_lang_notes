package main

import "fmt"
import "math"

func main(){

	myrect := Rectangle{10,30}
	mycirc := Circle{15}
	mytri := Triangle{10,10}

	fmt.Println(getarea(mycirc))
	fmt.Println(getarea(myrect))
	fmt.Println(getarea(mytri))

	myCat := Cat{"MIAAAWW"}
	myDog := Dog{"WOOOOOF"}

	getnoise(myCat)
	getnoise(myDog)



}
// the animal interface has a get noise method that checks the passed object for a noise method of its own
type Animal interface{
	noise()
}

type Dog struct{
	barks string
}

type Cat struct{
	miaw string
}

func (c Cat) noise() {
	fmt.Println(c.miaw)
}

func (d Dog) noise() {
	fmt.Println(d.barks)
}

func getnoise(animal Animal){
	animal.noise()
}

//-------------------------------------------

type Shape interface{
	area() float64
}

type Rectangle struct{
	height float64
	width float64
}

type Circle struct{
	radius float64
}

type Triangle struct{
	base float64
	height float64
}

func (tri Triangle) area() float64{
	return tri.base*tri.height/2
}

func (rect Rectangle) area() float64{
	return rect.width*rect.height
}

func (circ Circle) area() float64{
	return math.Pi * math.Pow(circ.radius, 2)
}

func getarea(shape Shape) float64{
	return shape.area()
}