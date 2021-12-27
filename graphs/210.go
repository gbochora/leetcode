package graphs

import "container/list"

func findOrder(numCourses int, prerequisites [][]int) []int {
	tree := buildPrereqTree(numCourses, prerequisites)
	rootCourses := getRootCourses(numCourses, prerequisites)
	prerequisites = buildPrereqMapping(numCourses, prerequisites)

	return getCourseOrderBFS(numCourses, tree, rootCourses, prerequisites)
}

func getCourseOrderBFS(numCourses int, tree [][]int, rootCourses []int, prerequisites [][]int) []int {
	visited := make([]bool, numCourses)
	queue := list.New()
	for _, course := range rootCourses {
		queue.PushBack(course)
	}
	order := make([]int, 0)

	for queue.Len() > 0 {
		e := queue.Front()
		queue.Remove(e)
		course := e.Value.(int)
		if visited[course] {
			continue
		}
		visited[course] = true
		order = append(order, course)

		for _, next := range tree[course] {
			allPrereqAreFinished := true
			for _, p := range prerequisites[next] {
				if !visited[p] {
					allPrereqAreFinished = false
					break
				}
			}
			if allPrereqAreFinished {
				queue.PushBack(next)
			}
		}
	}
	if len(order) != numCourses {
		order = []int{}
	}
	return order
}

func buildPrereqTree(numCourses int, prerequisites [][]int) [][]int {
	graph := make([][]int, numCourses)
	for _, prereq := range prerequisites {
		if graph[prereq[1]] == nil {
			graph[prereq[1]] = make([]int, 0)
		}
		graph[prereq[1]] = append(graph[prereq[1]], prereq[0])
	}
	return graph
}

func getRootCourses(numCourses int, prerequisites [][]int) []int {
	roots := make([]int, 0)
	notRoots := make([]bool, numCourses)
	for _, prereq := range prerequisites {
		notRoots[prereq[0]] = true
	}
	for course := range notRoots {
		if notRoots[course] {
			continue
		}
		roots = append(roots, course)
	}
	return roots
}

func buildPrereqMapping(numCourses int, prerequisites [][]int) [][]int {
	mapping := make([][]int, numCourses)
	for _, prereq := range prerequisites {
		if mapping[prereq[0]] == nil {
			mapping[prereq[0]] = make([]int, 0)
		}
		mapping[prereq[0]] = append(mapping[prereq[0]], prereq[1])
	}
	return mapping
}
