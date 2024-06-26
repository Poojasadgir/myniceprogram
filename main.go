package main

import "Poojasadgir/myniceprogram"

func main() {
	fmt.Println("Hello ,World!")
	 
	var whatToSay string
	whatToSay = "Goodbye ,a cruel word"

	fmt.Println(whatToSay, " ")

	var i int

	i = 10
	fmt.Print(i, " ")
	fmt.Println(saySomethin())
	whatwasSaid := saySomethin()

	fmt.Println(whatwasSaid)

	fmt.Println(saySomething())

	whatwasSaid, theOther := saySomething()

	fmt.Println(whatwasSaid, theOther)

}
func saySomethin() string {
	return "something"
}

func saySomething() (string, string) {
	return "something", "else"
}
