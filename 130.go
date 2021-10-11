package main

import "errors"

const VISITED = 'V'

func Solve(board [][]byte) {
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[x]); y++ {
			if board[x][y] == 'O' {
				findConnectedComponentBFS(board, x, y)
			}
		}
	}
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[x]); y++ {
			if board[x][y] == VISITED {
				board[x][y] = 'O'
			}
		}
	}
}

func findConnectedComponentBFS(board [][]byte, x, y int) {
	visitedCells := make(map[int]bool)
	q := NewQueue()
	q.Enqueue(pointToNumber(board, x, y))
	var connectedToBorder bool
	for !q.isEmpty() {
		num, _ := q.Dequeue()
		x, y = numberToPoint(board, num)
		if visitedCells[num] {
			continue
		}
		visitedCells[num] = true
		board[x][y] = VISITED
		if isBorderCell(board, x, y) {
			connectedToBorder = true
		}
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				if isAdjacentCell(dx, dy) && isInBounds(board, x+dx, y+dy) && board[x+dx][y+dy] == 'O' {
					q.Enqueue(pointToNumber(board, x+dx, y+dy))
				}
			}
		}
	}
	if !connectedToBorder {
		markAllVisitedCells(board, visitedCells)
	}
}

func markAllVisitedCells(board [][]byte, visitedCells map[int]bool) {
	for num := range visitedCells {
		x, y := numberToPoint(board, num)
		board[x][y] = 'X'
	}
}

func pointToNumber(board [][]byte, x, y int) int {
	return x*len(board[0]) + y
}

func numberToPoint(board [][]byte, num int) (int, int) {
	return num / len(board[0]), num % len(board[0])
}

func isBorderCell(board [][]byte, x, y int) bool {
	return x == 0 || y == 0 || x == len(board)-1 || y == len(board[0])-1
}

type Queue struct {
	in  []int
	out []int
}

func NewQueue() *Queue {
	q := new(Queue)
	q.in = make([]int, 0)
	q.out = make([]int, 0)
	return q
}

func (q *Queue) Dequeue() (int, error) {
	if q.isEmpty() {
		return 0, errors.New("Queue is empty")
	}

	if len(q.out) == 0 {
		for i := len(q.in) - 1; i >= 0; i-- {
			q.out = append(q.out, q.in[i])
		}
		q.in = q.in[:0]
	}
	b := q.out[len(q.out)-1]
	q.out = q.out[:len(q.out)-1]
	return b, nil
}

func (q *Queue) Enqueue(b int) {
	q.in = append(q.in, b)
}

func (q *Queue) Size() int {
	return len(q.in) + len(q.out)
}

func (q *Queue) isEmpty() bool {
	return q.Size() == 0
}
