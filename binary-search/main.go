package main

import "fmt"

type Person struct {
	name string
	age  int64
}

func main() {
	people := []Person{
		{name: "Alice", age: 12},
		{name: "Bob", age: 15},
		{name: "Charlie", age: 45},
		{name: "Diana", age: 14},
		{name: "Edward", age: 8},
		{name: "Fiona", age: 32},
		{name: "George", age: 30},
		{name: "Hannah", age: 19},
		{name: "Jane", age: 51},
		{name: "Mariah", age: 90},
	}

	search := BinarySearch(people, "Diana")
	fmt.Println(search)
}

func BinarySearch(arr []Person, name string) string {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]

		// in each iteration, we will compare the exact middle
		// item with the item we are performing the binary search on
		if guess.name == name {
			return fmt.Sprintf("found %s, age %d", guess.name, guess.age)
		}
		// if the `mid` value is less than the guess(name)
		// we continue the search in the upper half
		if guess.name < name {
			// where `low` takes the value of mid + 1
			low = mid + 1
		} else {
			// here the opposite
			high = mid - 1
		}
	}

	return fmt.Sprintf("not found a person with name %s", name)
}
