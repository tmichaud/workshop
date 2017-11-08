package main

import "fmt"

func main() {
	fmt.Println("Number operations" )
	fmt.Println("\t1 + 1 =", 1+1)
	fmt.Println("\t1.0 + 1.0 =", 1.0+1.0)
	fmt.Println("")

	fmt.Println("String operations" )
	fmt.Println(len("Hello, World"))    
	fmt.Println("Hello, World"[1])
	fmt.Println("                     e is 101...so...")    
	fmt.Println("Hello, " + "World")
	fmt.Println("")

	fmt.Println("Boolean operations" )
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
	fmt.Println("")


	fmt.Println("Exercise" )
	fmt.Println("\t32132 x 42452 = " , 32132*42452)
	fmt.Println((true && false) || (false && true) || !(false && false)) 


}
