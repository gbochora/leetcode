package graphs

import "sort"

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	informed := make([]bool, n)
	sort.SliceStable(meetings, func(i, j int) bool {
		return meetings[i][2] < meetings[j][2]
	})

	informed[0] = true
	informed[firstPerson] = true
	i := 0
	for i < len(meetings) {
		start := i
		disjointSet := make(map[int]map[int]bool)
		for i < len(meetings) && meetings[i][2] == meetings[start][2] {
			if disjointSet[meetings[i][0]] == nil {
				disjointSet[meetings[i][0]] = make(map[int]bool)
				disjointSet[meetings[i][0]][meetings[i][0]] = true
			}
			if disjointSet[meetings[i][1]] == nil {
				disjointSet[meetings[i][1]] = make(map[int]bool)
				disjointSet[meetings[i][1]][meetings[i][1]] = true
			}

			if len(disjointSet[meetings[i][0]]) > len(disjointSet[meetings[i][1]]) {
				unionMap(disjointSet[meetings[i][0]], disjointSet[meetings[i][1]])
				for key := range disjointSet[meetings[i][1]] {
					disjointSet[key] = disjointSet[meetings[i][0]]
				}
			} else {
				unionMap(disjointSet[meetings[i][1]], disjointSet[meetings[i][0]])
				for key := range disjointSet[meetings[i][0]] {
					disjointSet[key] = disjointSet[meetings[i][1]]
				}
			}
			i++
		}
		start = i

		for key := range disjointSet {
			if informed[key] {
				for k := range disjointSet[key] {
					informed[k] = true
					delete(disjointSet[key], k)
				}
			}
		}
	}

	informedList := make([]int, 0)
	for i := range informed {
		if informed[i] {
			informedList = append(informedList, i)
		}
	}

	return informedList
}

func unionMap(m1, m2 map[int]bool) {
	for key := range m2 {
		m1[key] = true
	}
}
