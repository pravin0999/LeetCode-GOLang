// This code is the solution to the LeetCode problem - https://leetcode.com/problems/two-sum/
//
// Problem Statement:
//
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.
// 
// Input/Output:
//
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Output: Because nums[0] + nums[1] == 9, we return [0, 1].

package main  
 
import "fmt"

func twoSum(nums []int, target int) []int {
	checker := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := checker[target - nums[i]]; ok {
			return []int{i, checker[target - nums[i]]}
		}
		checker[nums[i]] = i
	}
	return nil
}

func main() {
	var nums = []int{2,7,11,15}
	var target = 9
	
	var result = twoSum(nums, target)
	fmt.Println(result)
}

