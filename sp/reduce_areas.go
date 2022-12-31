package sp

import "sort"

const maxWords = 256

func ReduceAreas(areas []Area) []Area {
	sort.Slice(areas, func(i int, j int) bool {
		return areas[i].address < areas[j].address
	})
	return reduceSortedAreas(areas)
}

func reduceSortedAreas(areas []Area) []Area {
	if len(areas) <= 1 {
		return areas
	}
	a1 := areas[0]
	a2 := areas[1]
	words := Words(a2.Address()-a1.Address()) + a2.Words()
	if words > maxWords {
		return append([]Area{a1}, reduceSortedAreas(areas[1:])...)
	}
	return reduceSortedAreas(append([]Area{NewArea(a1.address, words)}, areas[2:]...))
}
