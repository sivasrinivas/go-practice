package main

import (
	"sort"
	"fmt"
)

type Person struct {
	Name string
	Age int
}

type ByName []Person

func (this ByName) Len() int {
	return len(this)
}

func (this ByName) Less(i, j int) bool {
	return this[i].Name < this[j].Name
}

func (this ByName) Swap(i, j int)  {
	this[i], this[j] = this[j], this[i]
}

func main() {
	kids := []Person{
		{"siva", 28},
		{"gita", 26},
	}
	sort.Sort(ByName(kids))
	fmt.Println(kids)

	numbers := []int{1,2,3,4,6,8,9,5,7}
	sort.Ints(numbers)
	fmt.Println(numbers)
	fmt.Println(sort.IntsAreSorted(numbers))

	names := []string{"siva", "gita"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.StringsAreSorted(names))
}
