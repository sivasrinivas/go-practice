package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type ByName []Person

func (b ByName) Len() int {
	return len(b)
}

func (b ByName) Less(i, j int) bool {
	return b[i].Name < b[j].Name
}

func (b ByName) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func main() {
	kids := []Person{
		{"siva", 28},
		{"gita", 26},
	}
	sort.Sort(ByName(kids))
	fmt.Println(kids)

	numbers := []int{1, 2, 3, 4, 6, 8, 9, 5, 7}
	sort.Ints(numbers)
	fmt.Println(numbers)
	fmt.Println(sort.IntsAreSorted(numbers))

	names := []string{"siva", "gita"}
	sort.Strings(names)
	fmt.Println(names)
	fmt.Println(sort.StringsAreSorted(names))
}
