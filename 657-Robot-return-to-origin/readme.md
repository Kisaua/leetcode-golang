# 657. Robot Return to Origin

Level: easy

[657. Robot Return to Origin](https://leetcode.com/problems/robot-return-to-origin/description/)

## Task

There is a robot starting at the position (0, 0), the origin, on a 2D plane.
Given a sequence of its moves, judge if this robot ends up at (0, 0) after it
completes its moves.

You are given a string moves that represents the move sequence of the robot
where moves[i] represents its ith move. Valid moves are 'R' (right), 'L'
(left), 'U' (up), and 'D' (down).

## Solution

Number of letters `L` minus number of `R`s is equal to zero and the same for
the letters `U` and `D`. Otherwise `false`.
