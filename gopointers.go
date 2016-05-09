package main
import "fmt"

func main () {

	x := 0

	// only passes the value of 0
	changeVal(x)

	fmt.Println("remains unchanged :",x)

	// pass in the memory address, instead of the value
	changeValNow(&x)
	fmt.Println("The Value in", &x, "is now :",x)


	y_pointer := new(int)
	changeYVal(y_pointer)
	fmt.Println("value of y pointer", *y_pointer)
}

func changeYVal(y *int) {
	*y = 10
}

func changeVal(v int) {
	v = 2
}

func changeValNow(v *int) { // pass in reference, access the memory address from the pointer

	// stores in the address of v
	*v = 5
}