package main

import (
	"fmt"
	"sort"
)

type byLength []string

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {

	s := []float64{5.2, -1.3, 0.7, -3.8, 2.6}
	fmt.Println("Not Sorted Float Array:", s)
	fmt.Println("Float Array are sorted:", sort.Float64sAreSorted(s))
	sort.Float64s(s)
	fmt.Println("Sorted Float Array:", s)
	fmt.Println("Float Array are sorted:", sort.Float64sAreSorted(s))
	sort.Sort(sort.Reverse(sort.Float64Slice(s)))
	fmt.Println("Reverse Float Array:", s)
	fmt.Println("Reversed Float Array are sorted:", sort.Float64sAreSorted(s))

	//s = []int{5, 2, 6, 3, 1, 4}
	//cannot use []int literal (type []int) as type []float64 in assignment

	p := []int{5, 2, 6, 3, 1, 4}
	fmt.Println("Not Sorted Int Array:", p)
	fmt.Println("Int Array are sorted:", sort.IntsAreSorted(p))
	sort.Ints(p)
	fmt.Println("Sorted Int Array:", p)
	fmt.Println("Int Array are sorted:", sort.IntsAreSorted(p))
	sort.Sort(sort.Reverse(sort.IntSlice(p)))
	fmt.Println("Reverse Int Array:", p)
	fmt.Println("Reversed Int Array are sorted:", sort.IntsAreSorted(p))

	r := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	fmt.Println("Not Sorted String Array:", r)
	fmt.Println("String Array are sorted:", sort.StringsAreSorted(r))
	sort.Strings(r)
	fmt.Println("Sorted String Array:", r)
	fmt.Println("String Array are sorted:", sort.StringsAreSorted(r))
	sort.Sort(sort.Reverse(sort.StringSlice(r)))
	fmt.Println("Reverse String Array:", r)
	fmt.Println("Reversed String Array are sorted:", sort.StringsAreSorted(r))

	//We implement sort.Interface - Len, Less, and Swap - on our type so we can use the sort package’s generic Sort function.
	//Len and Swap will usually be similar across types and Less will hold the actual custom sorting logic.
	//In our case we want to sort in order of increasing string length, so we use len(s[i]) and len(s[j]) here.
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println("Sort by function:", fruits)

	a := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55}
	x := 6

	i := sort.Search(len(a), func(i int) bool { return a[i] >= x })
	if i < len(a) && a[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}

	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}

	sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("By name:", people)

	sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	fmt.Println("By age:", people)
}
