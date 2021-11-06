package graphs

import "container/list"

func canFinish(numCourses int, prerequisites [][]int) bool {
	isPrereqFor, prereqCourses := buildGraphMaps(prerequisites)

	done := make(map[int]bool)
	queue := getCoursesWithNoPrerequisites(prereqCourses, numCourses)

	for queue.Len() > 0 {
		e := queue.Front()
		queue.Remove(e)
		course := e.Value.(int)
		done[course] = true

		for c := range isPrereqFor[course] {
			if done[c] {
				continue
			}
			if isAllCoursesFinished(prereqCourses[c], done) {
				queue.PushBack(c)
			}
		}
	}
	return len(done) == numCourses
}

func isAllCoursesFinished(prerequisites, done map[int]bool) bool {
	for c := range prerequisites {
		if !done[c] {
			return false
		}
	}
	return true
}

func buildGraphMaps(prerequisites [][]int) (isPrereqFor, prereqCourses map[int]map[int]bool) {
	isPrereqFor = make(map[int]map[int]bool)
	prereqCourses = make(map[int]map[int]bool)
	for i := range prerequisites {
		if _, ok := isPrereqFor[prerequisites[i][1]]; !ok {
			isPrereqFor[prerequisites[i][1]] = make(map[int]bool)
		}
		isPrereqFor[prerequisites[i][1]][prerequisites[i][0]] = true

		if _, ok := prereqCourses[prerequisites[i][0]]; !ok {
			prereqCourses[prerequisites[i][0]] = make(map[int]bool)
		}
		prereqCourses[prerequisites[i][0]][prerequisites[i][1]] = true
	}
	return
}

func getCoursesWithNoPrerequisites(prereqCourses map[int]map[int]bool, numCourses int) *list.List {
	queue := list.New()
	for i := 0; i < numCourses; i++ {
		if prereqCourses[i] == nil {
			queue.PushBack(i)
		}
	}
	return queue
}
