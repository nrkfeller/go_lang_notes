package main 

import ( "fmt" 
"time")

func main (){

	for j := 0; j < 10 ; j ++ {
		// calling the count routine
		go count(j)
	}

	time.Sleep(time.Millisecond*5000)

}

func count(id int) {

	// essentially j (or 10) count rountines are running simultaneously!
	for i := 0; i < 10 ; i ++ {
		fmt.Println(id, ":", i)

		time.Sleep(time.Millisecond * 500)
	}
}