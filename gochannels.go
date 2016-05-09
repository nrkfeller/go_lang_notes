package main 

import "fmt"
import "time"
import "strconv"

var pizzaNum = 0
var pizzaName = ""

func makeDough(stringChan chan string){

	pizzaNum++

	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)

	fmt.Println(pizzaName, "adding dough")

	stringChan <- pizzaName

	//time.Sleep(time.Millisecond * 10)
}
func makeSauce(stringChan chan string){

	pizza := <- stringChan

	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)

	fmt.Println(pizza, "adding sauce")

	stringChan <- pizzaName

	//time.Sleep(time.Millisecond * 10)
}
func makeToppings(stringChan chan string){

	pizza := <- stringChan

	fmt.Println(pizza, "adding toppings")

	//time.Sleep(time.Millisecond * 10)
}

func main (){

	// kind of like a safer global variable. Needs explicit passing
	stringChan := make(chan string)

	for i := 0; i < 3 ; i ++ {
		go makeDough(stringChan)
		go makeSauce(stringChan)
		go makeToppings(stringChan)

		time.Sleep(time.Millisecond * 1)
	}

	fmt.Println(pizzaNum, "pizzas total")
}