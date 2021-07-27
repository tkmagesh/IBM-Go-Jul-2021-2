package main

import "fmt"

func main() {
	//Array
	//var nos [5]int
	//var nos [5]int = [5]int{3, 1, 4, 2, 5}
	//var nos [5]int = [...]int{3, 1, 4, 2, 5}
	nos := [...]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	//iterate
	for i := 0; i < len(nos); i++ {
		fmt.Println(nos[i])
	}

	fmt.Println("Iterating using 'range'")
	for idx, no := range nos {
		fmt.Println(idx, no)
	}

	//creating a copy of an array
	newNos := nos
	newNos[0] = 200
	fmt.Println(nos, newNos)

	//Slice
	//var names []string
	var names = []string{"Gaurav", "Sanchit", "Rohit", "Kunal", "Saurabh"}
	fmt.Println(names)

	//adding new values
	//names = append(names, "Magesh")
	//names = append(names, "Magesh", "Suresh")

	newNames := []string{"Magesh", "Suresh"}
	names = append(names, newNames...)
	fmt.Println(names)

	//slicing
	/*
		[lo:hi] => from lo to hi-1
		[lo:] => from lo to end
		[:hi] => from 0 to hi-1
		[lo : lo] => empty slice
		[lo : lo+1] => item at 'lo'
		[:] => creates a copy of the slice
	*/
	fmt.Println(len(names), cap(names))
	fmt.Println(names[0:3])
	fmt.Println(names[3:])
	fmt.Println(names[:3])
	fmt.Println(names[3:3])
	fmt.Println(names[3:4])
	namesCopy := names[:]
	namesCopy[0] = "John"
	fmt.Println(names, namesCopy)

	namesCopy = append(namesCopy, "Faisal")
	fmt.Println(names, namesCopy)

	names = append(names, "Lokesh", "Kapil", "Srikar", "Akash")
	fmt.Println(names, namesCopy)
	fmt.Println(len(names), cap(names))

	list := make([]int, 10, 100)
	fmt.Println(list, len(list), cap(list))
}
