// This code is the solution to the LeetCode problem - https://leetcode.com/problems/palindrome-number 
//
// Problem Statement:
//
// Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.
// Follow up: Could you solve it without converting the integer to a string?
//
// Input/Output:
//
// Input: x = 121
// Output: true



package main  
 
import "fmt"

func reverse(x int) int {
	reverse := 0

	for x != 0 {
        t := x % 10
		reverse = reverse * 10 + t
		x = x / 10
	}
	return reverse
}

func isPalindrome(x int) bool {
    if x < 0 { return false }
	reverseX := reverse(x)
	return x == reverseX
}

func main() {
	x := 121

	output = isPalindrome(x)
	fmt.Println(output)
}

