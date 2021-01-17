// This code is the solution to the LeetCode problem - https://leetcode.com/problems/valid-parentheses 
//
// Problem Statement:
// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
// An input string is valid if:
//
    // Open brackets must be closed by the same type of brackets.
    // Open brackets must be closed in the correct order.
//
// Input/Output:
// Input: s = "()"
// Output: true

package main  
 
import "fmt"

func isValid(s string) bool {
    sRune := []rune(s)
    openBrackets := make(map[rune]rune)
    openBrackets['('] = ')'
    openBrackets['{'] = '}'
    openBrackets['['] = ']'
    
    if len(sRune) == 0 {
        return true
    }
    
    if len(sRune) == 1 {
        return false
    }
    
    mockStack := []rune{}

    for _, i := range sRune {
        
        if _, ok := openBrackets[i]; ok {
            mockStack = append(mockStack, i)
        } else {
            if len(mockStack) > 0 && openBrackets[mockStack[len(mockStack)-1]] == i { 
                mockStack = mockStack[:len(mockStack)-1] 
            } else {
                return false
            }
        }
    }
    if len(mockStack) == 0 {
        return true
    } else {
        return false
    }
}


func main() {
	
	s = "()[]{}"
	result = isValid(s)
	fmt.Println(result)
}

