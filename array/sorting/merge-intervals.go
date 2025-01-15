/*
Given an array of intervals where intervals[i] = [starti, endi], merge all overlapping intervals, 
and return an array of the non-overlapping intervals that cover all the intervals in the input.

Example 1:

Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
Example 2:

Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.

*/

func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i,j int) bool { return intervals[i][0] < intervals[j][0] })

    var result [][]int
    current := intervals[0]
    
    for j := range intervals {
        if intervals[j][0] > current[1] {
            result = append(result, current)
            current = intervals[j]
        }
        current[1] = max(current[1], intervals[j][1])
    }
    result = append(result, current)

    return result
}