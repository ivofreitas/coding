package practice

/*
There are n people standing in a queue, and they numbered from 0 to n - 1 in left to right order. 
You are given an array heights of distinct integers where heights[i] represents the height of the ith person.

A person can see another person to their right in the queue if everybody in between is shorter than both of them.
 More formally, the ith person can see the jth person if i < j and min(heights[i], heights[j]) > max(heights[i+1], heights[i+2], ..., heights[j-1]).

Return an array answer of length n where answer[i] is the number of people the ith person can see to their right in the queue.
*/
func canSeePersonsCount(heights []int) []int {
    // Initialize the result array with the same length as heights
    res := make([]int, len(heights))
    
    // Monotonic stack to keep track of heights
    monoStack := []int{}
    
    // Iterate from the end of the heights array to the beginning
    for i := len(heights) - 1; i >= 0; i-- {
        // If the stack is empty, no one can be seen, initialize the result
        if len(monoStack) == 0 {
            monoStack = append(monoStack, heights[i]) // Push current height
            res[i] = 0 // No one can be seen
            continue
        }

        tmp := 0 // Temporary counter for visible persons
        // While there are people in the stack and they are shorter than the current person
        for len(monoStack) > 0 && monoStack[len(monoStack)-1] < heights[i] {
            monoStack = monoStack[:len(monoStack)-1] // Pop from the stack
            tmp++ // Count this person as visible
        }

        // If the stack is not empty, the person on top is taller
        if len(monoStack) > 0 {
            tmp++ // Current person can see this taller person
        }

        // Push the current person's height onto the stack
        monoStack = append(monoStack, heights[i])
        // Store the count of visible persons in the result array
        res[i] = tmp
    }
    return res // Return the result array
}