# 874. Walking Robot Simulation

## Task

[874. Walking Robot Simulation](https://leetcode.com/problems/walking-robot-simulation/description/)

A robot on an infinite XY-plane starts at point `(0, 0)` facing north. The robot
receives an array of integers commands, which represents a sequence of moves
that it needs to execute. There are only three possible types of instructions
the robot can receive:

`-2`: Turn left 90 degrees.
`-1`: Turn right 90 degrees.
`1 <= k <= 9`: Move forward k units, one unit at a time.
Some of the grid squares are `obstacles`. The `ith` obstacle is at grid point
`obstacles[i] = (xi, yi)`. If the robot runs into an obstacle, it will stay in
its current location (on the block adjacent to the obstacle) and move onto the
next command.

Return the _maximum squared Euclidean distance_ that the robot reaches at any
point in its path (i.e. if the distance is `5`, return `25`).

## Solution

For the position and the robot, direction to move and obstacles are using `coordinates`.
Transform obstacles to the map of corresponding `coordinates` for fast lookup obstacle.

Starting direction is (0, 1) coordinate or facing north. Rotation clockwise `x,
y = y, -x` and counter clockwise is `x, y = -y, x`

Each move done one steep, checking the obstacles. After each move check the
distance and update to the maximum value. Return max distance in the end.
