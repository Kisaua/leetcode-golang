# 1914. Cyclically Rotating a Grid

You are given an m x n integer matrix grid, where `m` and `n` are both even
integers, and an integer `k`.

The matrix is composed of several layers.

to rotate a layer we have to create one dimension array
given array have two layers

1 2 3 4
16 1 2 5
15 8 3 6
14 7 4 7
13 6 5 8
12 11 10 9

algorithm will create two arrays

1...16, as for easy understanding numbers is from 1 to 16 is outer perimeter.
1...8 inner layer.

shift numbers in array, may be need to do with initial array and not to create a new one. didn't optimized.
then reverse assigning of the one dimensional array to the initial grid.

a bit wired calculation of the indexes, lack of time to make it look nicer.
