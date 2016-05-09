package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	// we can create variable explicitly
	var age int64 = 40

	// float or string defined
	var favNum float64 = 1.6180339

	// otherwise go can "guess what we want
	randNum := 1

	// if you do not use a variable; go raises and error
	fmt.Println(age, randNum, favNum)

	var numOne = 1.000
	var num99 = 0.999

	// Typical accuracy problem still exists
	fmt.Println(numOne - num99)
	// + - / %


	// immutable variable
	const pi float32 = 3.1415
	var (
		varA = 2
		varB = 3)
	fmt.Println(varA, varB, pi)

	// Strings
	myName := "Nicolas Feller" // same as var myName string = "Nicolas Feller"
	fmt.Println(len(myName))

	fmt.Println(myName + "is a cool \n guy")

	// bool
	var isOver40 bool = true // isOver40 := true

	if isOver40{
		// typical float print
		fmt.Printf("%.2f \n", pi)

		// type print
		fmt.Printf("%T \n", pi)

		// print integers
		fmt.Printf("%d \n", 123)
	}

	// logical operators
	fmt.Println(true && false, !true || false)

	// for loops
	var i int8 = 127
	j := 0
	for j <= 3 {
		fmt.Println(i)
		i ++
		j ++
	}
	for k := 0; k < 5; k ++ {
		fmt.Println(k)
	}

	// if statement -- if | else if | else

	// switch statement
	myAge := 26

	switch myAge{
		case 15 : fmt.Println("youre fifteen")
		case 26 : fmt.Println("youre more than 20")
		case 17 : fmt.Println("youre fifteen")
		default : fmt.Println("default case")
	}

	// Arrays
	var favNums1[5] float64
	favNums1[0] = 13
	favNums1[1] = 134
	favNums1[2] = 135
	favNums1[3] = 136
	//favNums1[4] = 138 // unasigned initialize to 0

	fmt.Println(favNums1)

	favNums2 := [5]float64 {54,65,76,87}//,98}
	fmt.Println(favNums2)

	for i, value := range favNums2 {
		fmt.Println(value + 100, i)
	}
	// if you dont want the index just put udnerscore
	for _, value := range favNums2 {
		fmt.Println(value + 100)
	}

	// number slices like python
	numSlice := []int {9,8,7,6,5,4}
	fmt.Println(numSlice[:3])
	fmt.Println(numSlice[2:5])

	// make range slice - start with 10 spots, max of 14
	numSlice2 := make([]int, 10, 14)
	copy(numSlice2, numSlice)
	fmt.Println(numSlice2)

	// append to missing space
	numSlice2 = append(numSlice2, 99, -1)
	numSlice2 = append(numSlice2, 98, -1)
	numSlice2 = append(numSlice2, 97, -1)
	fmt.Println(numSlice2)

	// maps or hash tables
	presidentAge := make(map[string] int)
	presidentAge["Teddy Roosevelt"] = 43
	presidentAge["Barry Owe"] = 12

	fmt.Println(presidentAge, "\nlength of Hmap", len(presidentAge))

	// Using functions
	listNums := []float64 {234,654,76,9,66}
	var result float64 = addThemUp(listNums)
	fmt.Println(result)

	// return next 2 vals
	var myval int8 = 10
	fmt.Println(nextTwoVal(myval))

	//variatic function - if we dont know how many things we are getting
	fmt.Println(subtractThem(1,2,3,4,5,6,7,8))

	// clojures
	numtodouble := 12
	doubleThis := func() int {
		// doubleThis can access the variable in the declared namespace
		numtodouble *= 2
		return numtodouble
	}
	fmt.Println(doubleThis())
	fmt.Println(doubleThis())

	// recursion
	fmt.Println(factorial(5))
	fmt.Println(fib(10))

	// defer - makes a function happen at the end of its namespace
	// can be used to clean up, while keeping code elegant
	//defer printTwo()
	printOne()

	// building safe functions
	fmt.Println(safeDivision(25,0))
	fmt.Println(safeDivision(25,5))

	// Panic demo
	demPanic()


}

func demPanic(){

	defer func(){
		fmt.Println(recover())
	}()
	panic("Panic!!")
}

func safeDivision (num1 int, num2 int) int {

	// this recovers the program if it raises and exception
	defer func(){
		// you could just not print anything
		fmt.Println(recover())
	}()
	solution := num1/num2
	return solution
}

func printOne() { fmt.Println(1)}
func printTwo() { fmt.Println(2)}

func fib(v int) int{
	if v < 2{
		return v
	}
	return fib(v-1) + fib(v-2)
}

func factorial(val int) int {
	if val == 0 {
		return 1
	}
	return val * factorial(val - 1)
}

// variatic function
func subtractThem(args ...int) int {
	result := 0 

	for _, v := range args {
		result -= v
	}

	return result
}

func nextTwoVal(v int8) (int8, int8){
	return v+1, v+2
}

func addThemUp(nums []float64) float64{

	var sum float64 = 0
	// or sum := 0.0

	for _, value := range nums{
		sum += value
	}

	return sum
}