package main

import (
	"fmt"
	"strings"
	"os"
	"io/ioutil"
	"path/filepath"
)

func create() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("But HOW to I handle an error?")
		return
	}
	defer file.Close()
	file.WriteString("Meow.")
}

func read() {
        file,err := os.Open("test.txt")
	if err != nil {
		fmt.Println("But HOW to I handle an error?")
	}
	defer file.Close()

	// Get filesize
	stat,err := file.Stat()
	if err != nil {
		fmt.Println("But HOW to I handle an error?")
		return
	}

	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		fmt.Println("But HOW to I handle an error?")
		return
	}

	str1 := string(bs)
	fmt.Println("File test.txt is size : (", stat.Size(), ") and contains (", str1, ")")
}

func fastread() {
	bs2, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("But HOW to I handle an error?")
		return
	}
	str2 := string(bs2)
	fmt.Println("File test.txt contains (", str2, ")")
}

func readdir() {
	dir, err := os.Open(".")
	if err != nil {
		fmt.Println("But HOW to I handle an error?")
		return
	}
	defer dir.Close()

       fileInfos, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("But HOW to I handle an error?")
		return
	}
	for _, fi := range fileInfos {
		fmt.Println("-- directory contains...", fi.Name());
	}
}

func walkdir() {
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(".... directory path is", path)
		return nil;
	})
}



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
	fmt.Println("contains .......... ", strings.Contains("test", "es"))
	fmt.Println("contains .......... ", strings.Contains("test", "yx"))
	fmt.Println("count ............. ", strings.Count("test", "t"))
	fmt.Println("HasPrefix ......... ", strings.HasPrefix("test", "te"))
	fmt.Println("HasSuffix ......... ", strings.HasSuffix("test", "st"))
	fmt.Println("Index ............. ", strings.Index("test", "e"))
	fmt.Println("Index ............. ", strings.Index("test", "s"))
	fmt.Println("Index ............. ", strings.Index("test", "y"))
	fmt.Println("Join .............. ", strings.Join([]string{"a string", "b string"}, " (seperator) "))
        fmt.Println("Join a,b - ........ ", strings.Join([]string{"a","b"},"-"))
        fmt.Println("Repeat a, 5 ....... ", strings.Repeat("a", 5))
        fmt.Println("Replace aaaa a b 2. ", strings.Replace("aaaa", "a", "b", 2))
	fmt.Println("Split a-b-c, -..... ", strings.Split("a-b-c-d-e", "-"))
	fmt.Println("ToLower............ ", strings.ToLower("THIS IS UPPERCASE"))
	fmt.Println("ToUpper............ ", strings.ToUpper("this is lowercase"))
        //arr := []byte("test")
	//str := string([]byte{'T', 'E', 'S', 'T'})
	fmt.Println("To binary.......... no example here.")

	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	fmt.Println(" Input/Output")
	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")

	fmt.Println("-- Creating a file            ------")
	create();
	fmt.Println("-- Reading a file            ------")
	read();
	fmt.Println("-- Faster method to read a file ---")
	fastread();
	fmt.Println("-- Reading the directory ---")
	readdir();
	fmt.Println("-- Walking the directory (and subdirectories) ---")
	walkdir();
}

