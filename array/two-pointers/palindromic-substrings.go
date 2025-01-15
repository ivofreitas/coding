package two_pointers

/*
Given a string s, return the number of palindromic substrings in it.

A string is a palindrome when it reads the same backward as forward.

A substring is a contiguous sequence of characters within the string.

 

Example 1:

Input: s = "abc"
Output: 3
Explanation: Three palindromic strings: "a", "b", "c".
Example 2:

Input: s = "aaa"
Output: 6
Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
*/
func countSubstrings(s string) int {
	count := 0
	for i := range s {
		count += isPalindromeSubstring(s, i, i) + isPalindromeSubstring(s, i, i+1)
	}
	return count
}

func isPalindromeSubstring(s string, l, r int) int {
	res := 0
	for l >= 0 && r < len(s) && s[l] == s[r] {
		res++
		l--
		r++
	}
	return res
}
