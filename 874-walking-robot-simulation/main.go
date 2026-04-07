package main

func main() {
}

type point struct {
	x, y int
}

func robotSim(commands []int, obstacles [][]int) int {
	direction := point{0, 1}
	position := point{0, 0}
	obsacleMap := map[point]struct{}{}
	for i := range obstacles {
		p := point{
			x: obstacles[i][0],
			y: obstacles[i][1],
		}
		obsacleMap[p] = struct{}{}
	}
	distance := 0

	for _, comand := range commands {
		switch comand {
		case -1:
			direction.x, direction.y = direction.y, -direction.x
		case -2:
			direction.x, direction.y = -direction.y, direction.x
		default:
			position = makeMove(position, direction, comand, obsacleMap)
			if position.x*position.x+position.y*position.y > distance {
				distance = position.x*position.x + position.y*position.y
			}
		}
	}

	return distance
}

func makeMove(start, direction point, count int, obstacles map[point]struct{}) point {
	for i := 1; i <= count; i++ {
		p := point{
			x: start.x + direction.x,
			y: start.y + direction.y,
		}
		if _, ok := obstacles[p]; ok {
			return start
		}
		start = p
	}
	return start
}
