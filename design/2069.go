package design

import "github.com/gbochora/leetcode/graphs"

func AbsInt(k int) int {
	if k < 0 {
		return -k
	}
	return k
}

var directions [][]int = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
var directionsNames []string = []string{"East", "North", "West", "South"}

type Robot struct {
	x, y, dir, w, h int
}

func ConstructorRobot(width int, height int) Robot {
	var r Robot
	r.w = width
	r.h = height
	return r
}

func (this *Robot) Move(num int) {
	num = num%(2*(this.w-1)+2*(this.h-1)) + (2*(this.w-1) + 2*(this.h-1))
	for num > 0 {
		dx := directions[this.dir][0] * num
		dy := directions[this.dir][1] * num
		if !graphs.IsInBounds(this.w, this.h, this.x+dx, this.y+dy) {
			this.dir = (this.dir + 1) % 4
		}
		if this.x+dx < 0 {
			dx = -this.x
		} else if this.x+dx >= this.w {
			dx = this.w - this.x - 1
		} else if this.y+dy < 0 {
			dy = -this.y
		} else if this.y+dy >= this.h {
			dy = this.h - this.y - 1
		}
		this.x += dx
		this.y += dy
		num -= AbsInt(dx) + AbsInt(dy)
	}
}

func (this *Robot) GetPos() []int {
	return []int{this.x, this.y}
}

func (this *Robot) GetDir() string {
	return directionsNames[this.dir]
}
