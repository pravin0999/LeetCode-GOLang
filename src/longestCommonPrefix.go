// This code is the solution to the LeetCode problem -  https://leetcode.com/problems/longest-common-prefix 
//
// Problem Statement:
// Write a function to find the longest common prefix string amongst an array of strings.
//
// If there is no common prefix, return an empty string "".
//
// Input/Output:
//
// Input: strs = ["flower","flow","flight"]
// Output: "fl"

package main  
 
import "fmt"


func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }
    if len(strs) == 1 {
        return strs[0]
    }

    prefix_len := len(strs[0])
    for i := 1; i < len(strs); i++ {

        if prefix_len == 0 {
            return ""
        }

        if prefix_len > len(strs[i]) {
            prefix_len = len(strs[i])
        }

        next_prefix_len := 0
        for j := 0; j < prefix_len; j++ {
            if strs[i-1][j] == strs[i][j] {
                next_prefix_len = next_prefix_len + 1
            } else {
                prefix_len = next_prefix_len
                next_prefix_len = 0
                break
            }
        }
    }

    return strs[0][0:prefix_len]

}

func main() {
	strs := make([]string, 3)
	strs[0] = "flower"
	strs[1] = "flower"
	strs[2] = "flight"

	result = longestCommonPrefix(strs)

	fmt.Println(result)

}

