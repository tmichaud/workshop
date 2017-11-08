package main

import "fmt"

var aString string = "Hello, World"

func showAString() {
	fmt.Println(aString)
}

func showBString() {
//	fmt.Println(bString) // No access to bString (yet)
}

func main() {
	var x string = "Hello, World"
	fmt.Println(x)

	var y string
	y = "Hello, World"
	fmt.Println(y)

	x = "first"
	fmt.Println(x)
	x = x + " second"
	fmt.Println(x)

	x = "hello"
	y = "world"
	fmt.Println(x == y)

	x = "hello"
	y = "hello"
	fmt.Println(x == y)	

	// Type can be inferred
	var z = "Hello, World"
	z1 := "Hello, World"
	fmt.Println(z == z1)

	// Scope
	fmt.Println(aString)
	showAString()	

	var bString string = "Hello, World"
	fmt.Println(bString)
	//showBString()	 // Errors out

	
	// Constants
	fmt.Println("\nConstants")
	const aConstant string = "Hello, World"
	fmt.Println(aConstant)

	//aConstant = "something else"  // Throws an compile time error		

	// Multiple Values
	fmt.Println("\nMultipleValues")
	var (
		mv1 = 5
		mv2 = 10
		mv3 = 15
	)
	const (
		cv1 = 5
		cv2 = 10
		cv3 = 15
	)
	fmt.Println("MV values: ", mv1, mv2, mv3)
	fmt.Println("CV values: ", cv1, cv2, cv3)

	// Example
	fmt.Println("\nExample of double")
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)


	// Exercises
	fmt.Println("\nExercises:")
	ex := 5
	ex += 1
	fmt.Println("what is x:= 5, x+1 afterward?  ", ex)

	fmt.Print("\nFarenheit into Celsius... enter a tempature: ")
	var inputF float64
	fmt.Scanf("%f", inputF)
	fmt.Println("Output in Celsius: ", (inputF - 32) * 5/9)
	
	fmt.Print("\nFeet into meter...enter feet:" )
	var inputFeet float64
	fmt.Scanf("%f", inputFeet)
	fmt.Println("Output in meters: ", inputFeet * 0.3048)	


	fmt.Println("-----------")
	fmt.Println("aString = " + aString);
	aString = "test2"
	fmt.Println("aString = " + aString);
	

}
