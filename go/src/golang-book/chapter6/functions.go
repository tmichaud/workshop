package main

import "fmt"


func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}
func f() (int,int) {
	return 5,6
}

func average(scores []float64) float64 {
	//panic("Not Implemented")
	total := 0.0

	for _, value := range scores {
		total += value
	}
	return total / float64(len(scores))

}
func makeEvenGenerator() func() uint {
	i := uint(0)  // Note: i keeps it state below....
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x - 1)
}

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}


func testpanic(msg string) {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic(msg)
}

func catchPanic() {
	str := recover()
	fmt.Println(str)
}

func testpanic2a(msg string) {
	panic(msg)
}

func testpanic2(msg string) {
	defer catchPanic()
	testpanic2a(msg)
}

func testpanic3(msg string) {
	defer catchPanic()
}

func main() {
	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	fmt.Println(" functions ")
	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	scores := []float64 { 98, 93, 77, 82, 83 }

	//total := 0.0
	//for _, value := range scores {
	//	total += value
	//}
	//fmt.Println(total / float64(len(scores)))

	fmt.Println(average(scores))

	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	fmt.Println(" return multiple items             ")
	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	x,y := f()
	fmt.Println(" x = (", x, ") y = (", y, ")")

        fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	fmt.Println(" Variadic function                 ")
	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")

	fmt.Println(" Total of add (1,2,3) is : ", add(1,2,3))
	xs_array := []int{1,2,3}
	fmt.Println(" Adding a slice of ints : ", (add(xs_array...)))
	fmt.Println(" Adding a slice of ints : ", add([]int{1,2,3}...))


        fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	fmt.Println(" Closure                           ")
	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")

	x1 := 0
	increment := func() int {
		x1++
		return x1;
	}
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println("x1 = (", x1, ")")

	localadd := func(x, y int) int {
		return x + y
	}
	fmt.Println(localadd(1,1))
	fmt.Println("x = (", x , ")")


        fmt.Println("-----------------------------------")

	nextEven := makeEvenGenerator()  // Because of how we declare it here - it keeps state?
	fmt.Println("  nextEven is (should be 0) ", nextEven())
	fmt.Println("  nextEven is (should be 2) ", nextEven())
	fmt.Println("  nextEven is (should be 4) ", nextEven())


        fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	fmt.Println(" Recursion                         ")
	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")

	fmt.Println("Calling factorial(2) (2 * 1)     ", factorial(2))
	fmt.Println("Calling factorial(3) (3 * 2 * 1) ", factorial(3))


        fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	fmt.Println(" defer,panic and recover           ")
	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	defer second() // Moves second to the END of the main func....
	first()
	defer first()
	fmt.Println("Calling factorial(3) (3 * 2 * 1) ", factorial(3))


//	defer func() {
//		str := recover()
//		fmt.Println(str)
//	}()
//	panic ("PANIC!")

	testpanic("PANIC!")
	testpanic2("PANIC2!!!")
	testpanic3("PANIC3!!!")
	fmt.Println("Can't get here")


        fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	fmt.Println(" Pointers                          ")
	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	






	
}
