package main

//import ("fmt"; "math"; "strings")
import ("fmt"; "strings")


func main() {
	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	fmt.Println(" Packages  ")
	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	
	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	fmt.Println(" Strings ")
	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	fmt.Println("contains .... ", strings.Contains("test", "es"))
	fmt.Println("contains .... ", strings.Contains("test", "yx"))
	fmt.Println("count ....... ", strings.Count("test", "t"))
	fmt.Println("HasPrefix ... ", strings.HasPrefix("test", "te"))
	fmt.Println("HasSuffix ... ", strings.HasSuffix("test", "st"))
	fmt.Println("Index ....... ", strings.Index("test", "e"))
	fmt.Println("Index ....... ", strings.Index("test", "s"))
	fmt.Println("Index ....... ", strings.Index("test", "y"))
	fmt.Println("Join ........ ", strings.Join([]string{"a string", "b string"}, " (seperator) "))
}

