package main

type coord struct {
	x, y int
}

type Robot struct {
	field     coord
	position  coord
	direction coord
}

func Constructor(width int, height int) Robot {
	return Robot{
		field: coord{
			x: width,
			y: height,
		},
		position: coord{
			x: 0,
			y: 0,
		},
		direction: coord{
			x: 1,
			y: 0,
		},
	}
}

func (r *Robot) Step(num int) {
	l := 2 * (r.field.x + r.field.y - 2)
	if l == 0 || num == 0 {
		return
	}
	num = num % l
	if num == 0 {
		num = l
	}

	for i := 0; i < num; i++ {
		nextPosition := coord{
			x: r.position.x + r.direction.x,
			y: r.position.y + r.direction.y,
		}
		if nextPosition.x > r.field.x-1 || nextPosition.x < 0 || nextPosition.y > r.field.y-1 || nextPosition.y < 0 {
			r.direction.x, r.direction.y = -r.direction.y, r.direction.x
			i -= 1
			continue
		}
		r.position = nextPosition
	}
}

func (r *Robot) GetPos() []int {
	return []int{r.position.x, r.position.y}
}

func (r *Robot) GetDir() string {
	switch r.direction {
	case coord{1, 0}:
		return "East"
	case coord{0, 1}:
		return "North"
	case coord{0, -1}:
		return "South"
	case coord{-1, 0}:
		return "West"
	}
	return "unknown"
}

/**
 * Your Robot object will be instantiated and called as such:
 * obj := Constructor(width, height);
 * obj.Step(num);
 * param_2 := obj.GetPos();
 * param_3 := obj.GetDir();
 */
