// This code is the solution to the LeetCode problem - https://leetcode.com/problems/reverse-integer 
//
// Problem Statement:
//
// Given a 32-bit signed integer, reverse digits of an integer.
//
// Note:
// Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: [−231,  231 − 1].
// For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
//
// Input/Output:
//
// Input: x = 123
// Output: 321


package reverse  
 
import "fmt"

func reverse(x int) int {
	reverse := 0
	for x != 0 {
		a := x % 10
		reverse = reverse*10 + a
		x = x / 10
	}
    max_int := int(math.Pow(2, 31) - 1)
    min_int := int(math.Pow(-2, 31))
    if reverse < min_int || reverse > max_int {
        return 0 }
	return reverse
}

func main() {
	x := 123
	output := reverse()
	fmt.Println(output)
}

