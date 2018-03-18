package main

import (
	"fmt"
	"strings"
	"os"
        "io"
	"io/ioutil"
	"path/filepath"
        "errors"
	"container/list"
	"sort"
	"hash/crc32"
	"crypto/sha1"
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

func throwError() error {
	err := errors.New("error message")
	return err
}

func testerrors() {
	err := throwError()
	if err != nil {
		fmt.Println(".... error caught : ", err)
	}
}

func beyondListsAndMaps() {
	fmt.Println("   beyond lists and maps ")
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)
	x.PushFront(4)
	x.PushFront(5)
	x.PushFront(6)

	for e := x.Front(); e != nil;  e = e.Next() {
		fmt.Println(e.Value.(int))
	}

	fmt.Println("....  now...backwards!")
	for e := x.Back(); e != nil;  e = e.Prev() {
		fmt.Println(e.Value.(int))
	}
}

type Person struct {
	Name string
	Age int
}

type ByName []Person

func (ps ByName) Len() int {
	return len(ps)
}

func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}

func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

type ByAge []Person
func (this ByAge) Len() int {
	return len(this)
}

func (this ByAge) Less (i, j int) bool {
	return this[i].Age < this[j].Age
}

func (this ByAge) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func beyondSorting() {
	kids := []Person {
		{ "Alpha", 3000},
		{ "Timmy", 25},
		{ "Jill", 9},
		{ "Jack", 10},
		{ "Zack", 2},
	}

	fmt.Println("...  Sorting kids alphabetically!")
	fmt.Println("......   remember, the sort.Interface requires three methods - Len, Less and Swap")
	fmt.Println("......        For Len (returning int) should return the length of the thing we are sorting.  For a slice simple return len(ps) ")
	fmt.Println("......        Less (returning bool) is used to determine whether the item at position i is strictly lesss than the item at position j. ")
	fmt.Println(".....         Swap (no return) swaps these items. ")

	sort.Sort(ByName(kids))
	fmt.Println(kids)

	fmt.Println("...  Sorting kids by Age!")
	sort.Sort(ByAge(kids))
	fmt.Println(kids)
}

func getHash(filename string) (uint32, error) {
	// Open the file
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}

	// close the file at the end
	defer f.Close()

	h := crc32.NewIEEE()
	_, err1 := io.Copy(h, f)  // Go complains that err1 shouldn't be err.  And example in book is wrong...io not io/ioutil

	// We don't care about how many bytes were written but we do want to handle the error
	if err1 != nil {
		return 0, err1
	}

	return h.Sum32(), nil
}

func beyondhashesAndCryptography() {
	fmt.Println("... A hash function takes a set of data and reduces it to a smaller fixed size.  In go: 2 categories: ")
	fmt.Println("...    cryptographic ")
	fmt.Println("...    non-cryptographic ")
	fmt.Println("... ")
	fmt.Println("...   - non-cryptographic - adler32,, crc32, crc64, fnv.   Example of crc32 : ")
	h := crc32.NewIEEE()
	h.Write([]byte("test"))
	v := h.Sum32()
	fmt.Println("...   crc32 of 'test' => ", v)

	fmt.Println("...   non-cryptographic hashes are sometimes used to compare files ")
	h1, err := getHash("test.txt")
	if err != nil {
		fmt.Println("Wait - we're supposed to have a test.txt file!")
		return
	}
	h2, err := getHash("test.txt")  // Hashes had better be the same
	if err != nil {
		fmt.Println("Wait - we're supposed to have a test.txt file!")
		return
	}
	h3, err := getHash("packages.go") // Hashes had better be different
	if err != nil {
		fmt.Println("Wait - we're supposed to have a packages.go file!")
		return
	}
	fmt.Println("h1 and h2 should match    : ", h1, h2, h1 == h2)
	fmt.Println("h1 and h3 shouldn't match : ", h1, h3, h1 == h3)

	fmt.Println("...   - cryptographic are similar to non-cryptographic - but are difficult to reverse.  Used in security")
	fmt.Println("...   Example: SHA-1 ")
	sh1 := sha1.New()
	sh1.Write([]byte("test"))
	bs := sh1.Sum([]byte{})
	fmt.Println("... sha1 of 'test' => ", bs)
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

	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	fmt.Println(" Errors ")
	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	testerrors()

	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	fmt.Println(" Containers and Sort ")
	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	beyondListsAndMaps()
	beyondSorting()

        fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	fmt.Println(" Hashes and Cryptography ")
	fmt.Println("-----------------------------------")
	fmt.Println("-----------------------------------")
	beyondhashesAndCryptography()
}
