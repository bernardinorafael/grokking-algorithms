package main

import "fmt"

type Song struct {
	Name       string
	Author     string
	PlaysCount int32
}

func main() {
	songs := []Song{
		{Name: "Aproveita", Author: "Clayton & Romário", PlaysCount: 1209},
		{Name: "Amor Na Praia", Author: "NATTAN", PlaysCount: 109},
		{Name: "Anestesiado", Author: "Murilo Huff", PlaysCount: 187},
		{Name: "Alibi", Author: "Sevdaliza, Pabllo Vittar", PlaysCount: 876},
		{Name: "Vazou Na Braquiara", Author: "Hugo & Guilherme", PlaysCount: 1000},
		{Name: "Tumbátum", Author: "Tropa do Bruxo", PlaysCount: 98},
		{Name: "Morena de goiânia", Author: "Hugo & Guilherme, Maiara 7 Maraisa", PlaysCount: 18},
		{Name: "Como Um Beijo", Author: "Marcos & Belutti, Nadson O Feirinha", PlaysCount: 998},
		{Name: "Casca de Bala", Author: "Thullio Milionário", PlaysCount: 716},
		{Name: "Dois Tristes", Author: "Simone Mendes", PlaysCount: 91},
	}

	ordered := SelectionSort(songs)
	fmt.Println(ordered)
}

func SelectionSort(songs []Song) []Song {
	var ordered []Song

	// makes a copy to preserve the original array
	// the built-in go append fn always return a copy of the slice src
	copy := append(ordered, songs...)

	for len(copy) > 0 {
		// finds the index of the song with the lowest count
		minIndex := FindMinIndex(copy)
		// adds the song with the lowest plays count
		// at the beginning of the original array(ordered)
		ordered = append(ordered, copy[minIndex])
		// remove from copy slice the minIndex song
		copy = append(copy[:minIndex], copy[minIndex+1:]...)
	}
	return ordered
}

func FindMinIndex(songs []Song) (minIndex int) {
	minIndex = 0
	for i := range songs {
		// if the observable item of the array(songs[i]) is smaller
		// than my initial item(min), the smallest item
		// in the array has been found, we end loop
		if songs[i].PlaysCount < songs[minIndex].PlaysCount {
			// we assign the index and continue
			minIndex = i
		}
	}
	return
}
