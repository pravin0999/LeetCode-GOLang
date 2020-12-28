// https://leetcode.com/problems/search-insert-position

func searchInsert(nums []int, target int) int {
    n,m := 0, len(nums) - 1
    
    for n <= m {
        mid := (n+m)/2
        if nums[mid] == target {
            return mid
        } 
        if nums[mid] < target {
            n = mid + 1
        } else {
            m = mid -1
        }
    }
    return n
}

