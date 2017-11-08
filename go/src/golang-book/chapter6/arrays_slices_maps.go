package main

import "fmt"

func main() {
	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	fmt.Println(" Arrays ")
	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	var x [5]int
	x[4] = 100
	fmt.Println(x)

//	var scores [5]float64
//	scores[0] = 98
//	scores[1] = 93
//	scores[2] = 77
//	scores[3] = 82
//	scores[4] = 83
//	scores := [5]float64 { 98, 93, 77, 82, 83}
	
//	or
	scores := [5]float64 {
		98,
		93,
		77,
		82,
		83,
	}

	var total float64 = 0
//	for i:= 0; i < 5; i++ {
//		total += scores[i]
//	}
	for _, value := range scores {
		total += value
	}

	fmt.Println("Scores length : " , len(scores))
	fmt.Println(total / float64(len(scores)))


	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	fmt.Println(" Slices ")
	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	
	slice1 := []int {1,2,3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)

	slice3 := make([]int, 2)
	copy (slice3, slice1) 
	fmt.Println(slice1, slice3)
	
	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	fmt.Println(" Maps ")
	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	//var my_map map[string] int
	my_map := make(map[string]int)
	my_map["key"] = 10
	fmt.Println("my_map (string) value is: ", my_map["key"])

	fmt.Println("")
	my_map_int := make(map[int]int)
	fmt.Println("My_map_int length is ", len(my_map_int))
	my_map_int[1] = 10
	fmt.Println("My_map_int length is ", len(my_map_int))
	fmt.Println(my_map_int[1])
	fmt.Println("My_map_int length is ", len(my_map_int))
	delete(my_map_int, 1)
	fmt.Println("My_map_int length is ", len(my_map_int))

	fmt.Println("")
//	elements := make(map[string]string)
//	elements["H"] = "Hydrogen"
//	elements["He"] = "Helium"
//	elements["Li"] = "Lithium"
//	elements["Be"] = "Beryllium"
//	elements["B"] = "Boron"
//	elements["C"] = "Carbon"
//	elements["N"] = "Nitrogen"
//	elements["O"] = "Oxygen"
//	elements["F"] = "Fluorine"
//	elements["Ne"] = "Neon"

	// Shorter way to create a map
	elements := map[string]string {
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}

	fmt.Println(elements["Li"])
	fmt.Println(elements["Un"]) // Shows nothing...not good.

	name, ok := elements["Un"] // Another way to do it - but just shows false
	fmt.Println(name, ok)

	if name, ok := elements["Un"]; ok { // Don't do anything if not ok....
		fmt.Println(name, ok)
	}
	fmt.Println("")

	// Map of string to (map of string to string)
	elements2 := map[string]map[string]string{
		"H": map[string]string{
			"name":"Hydrogen",
			"state":"gas",
		},
		"He": map[string]string{
			"name":"Helium",
			"state":"gas",
		},
	        "Li": map[string]string{
			"name":"Lithium",
			"state":"solid",
		},
	        "Be": map[string]string{
			"name":"Beryllium",
			"state":"solid",
		},
	        "B":  map[string]string{
			"name":"Boron",
			"state":"solid",
		},
	        "C":  map[string]string{
			"name":"Carbon",
			"state":"solid",
		},
	        "N":  map[string]string{
			"name":"Nitrogen",
			"state":"gas",
		},
	        "O":  map[string]string{
			"name":"Oxygen",
			"state":"gas",
		},
	        "F":  map[string]string{
			"name":"Fluorine",
			"state":"gas",
		},
	        "Ne":  map[string]string{
			"name":"Neon",
			"state":"gas",
		},
	}

	if el, ok := elements2["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	} else {
		fmt.Println("Element not found.")
	}


	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	fmt.Println(" Exercises ")
	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	
	array := [6]float64 { 98, 93, 77, 82, 83, 24}
	slice := []int {1,2,3,4,5,6,7}

	fmt.Println("4 element of array: (", array[3], ")  and slice: (", slice[3], ")")
	
	slicex := make([]int, 3, 9)
	fmt.Println("Length should be 3 with a capacity of 9 : ", len(slicex))
		
	abc := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println("Given array a-f....array[2:5] should give : c, d, e...    Result: ", abc[2:5])

	numbers := []int {
		48,
		96,
		86,
		68,
		57,
		82,
		63,
		70,
		37,
		34,
		83,
		27,
		19,
		97,
		 9,
		17,
	}

	a_num := numbers[0]
	for _, value := range numbers {
		if (a_num > value) {
			a_num = value
		}
	}
	fmt.Println("a_num should be 9 now: (", a_num, ")")
}
