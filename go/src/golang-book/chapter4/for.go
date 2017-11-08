package main

import "fmt"

func main() {
	i := 1
	for i <= 10  {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println("Take 2")
	for j := 1 ; j <= 10; j++ {
		if j % 2 == 0 {
			fmt.Println(j , " is even")
		} else {
			fmt.Println(j , " is odd")
		}
	}

	fmt.Println("----------------------------------------")
	fmt.Println("----------------------------------------")
	fmt.Println("Exercises:")
	fmt.Println("----------------------------------------")
	fmt.Println("----------------------------------------")
	fmt.Println("1 : Small - 10 is not > 10 ")
	i = 10
	if i > 10 {
		fmt.Println("Big")
	} else {
		fmt.Println("Small")
	}


	fmt.Println("2 : Print out all numbers between 1 and 100 evenly divisible by 3 ")
	for j := 1 ; j <= 100; j++ {
		if j % 3 == 0 {
			fmt.Println(j , " is divisible by 3")
		}
	}


	fmt.Println("3 : Do FizzBuzz ")
	for j := 1 ; j <= 100; j++ {
		if j % 3 == 0 && j % 5 == 0 {
			fmt.Println(j , " FizzBuzz")
		} else if j % 3 == 0 {
			fmt.Println(j , " Fizz")
		} else if j % 5 == 0 {
			fmt.Println(j , " Buzz")
		} else {
			fmt.Println(j)
		}
	}
}
