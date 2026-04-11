# 3740/3741 Minimum distance between three Elements

[Leetcode - 3471 Minimum distance](https://leetcode.com/problems/minimum-distance-between-three-equal-elements-ii/description)

You are given an integer array `nums`.

A tuple `(i, j, k)` of 3 distinct indices is good if `nums[i] == nums[j] == nums[k]`.

The distance of a good tuple is `abs(i - j) + abs(j - k) + abs(k - i)`, where `abs(x)`
denotes the absolute value of x.

Return an integer denoting the minimum possible distance of a good tuple. If no good tuples exist, return -1.
