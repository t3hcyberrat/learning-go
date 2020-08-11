package main

import "fmt"

var y = 911

func main() {

	fmt.Println(y)
	fmt.Printf("%T\n", y)  // print type
	fmt.Printf("%b\n", y)  // print binary
	fmt.Printf("%x\n", y)  // print hex without prededing 0x
	fmt.Printf("%#X\n", y) // print hex with preceding 0x

	y = 100

	fmt.Printf("%#X", y)
	fmt.Printf("%#X\t%b\t%x\n", y, y, y) // print using tab between values

	s := fmt.Sprintf("%#X\t%b\t%x\n", y, y, y) // transform to string
	fmt.Println(s)                             // print as str
	fmt.Printf("%T\n", s)                      // check if type is str

	fmt.Printf("%v\n", y) // back to normal value

}
