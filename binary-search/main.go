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
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		target := arr[mid]

		if target.name == name {
			return fmt.Sprintf("found %s, age %d", target.name, target.age)
		}

		if target.name < name {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return fmt.Sprintf("not found a person with name %s", name)
}
