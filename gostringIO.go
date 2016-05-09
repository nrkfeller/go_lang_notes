package main

import "fmt" 
import ("strings" 
"sort" 
"os" 
"log" 
"io/ioutil"
"strconv")

func main(){

	hw := "hello world"

	fmt.Println(strings.Contains(hw, "hell"))
	fmt.Println(strings.Index(hw, " w"))
	fmt.Println(strings.Count(hw, "o"))
	fmt.Println(strings.Replace(hw, "l", "p",2))

	stringWithCommas := "1,2,3,4,5,6"
	fmt.Println(strings.Split(stringWithCommas, ","))

	listOfLetters := []string{"c","D","f"}
	sort.Strings(listOfLetters)
	fmt.Println(listOfLetters)

	listOfNums := []string {"1","3","6","4"}
	strings.Join(listOfNums, ",")
	fmt.Println(listOfNums)

//-----------FILE IO-----------------------------
	file, err := os.Create("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("new text yay!!")

	file.Close()

	stream,err := ioutil.ReadFile("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	readString := string(stream)

	fmt.Println(readString)

//-----------EXPECTING INPUT--------------------
	fmt.Println("Enter name :")
	var name string
	fmt.Scan(&name)

	fmt.Println("Your name is :", name)

//-----------CASTING----------------------------

	randInt := 1
	randFloat := 12.43
	randString := "100"
	randString2 := "12.43"

	fmt.Println(float64(randFloat))
	fmt.Println(int(randFloat)) // truncates
	fmt.Println(int(randInt))

	// parse to find numbers in strings
	newInt, _ := strconv.ParseInt(randString, 0, 64)
    fmt.Println(newInt)
    
    newFloat, _ := strconv.ParseFloat(randString2, 64)
    fmt.Println(newFloat)

}

