# Count Submatrices with Top-Left Element and Sum Less Than k

You are given a 0-indexed integer matrix `grid`` and an integer`k`.

Return the number of submatrices that contain the top-left element of the `grid`
,and have a sum less than or equal to`k.

## Solution

Traverse the element in the matrix and update each element with the sum of the
top and left element. In this way your element is matrix become the sum of the
submatrice that contains top-left element. If this element is lower than `k`
increase counter, otherwise no need to continue counting and continue with the
next row.
