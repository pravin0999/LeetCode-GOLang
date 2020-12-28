// https://leetcode.com/problems/remove-duplicates-from-sorted-array

func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    trackingIndex := 0
    for i := 1; i < len(nums); i++ {
        if nums[trackingIndex] != nums[i] {
            trackingIndex += 1
            nums[trackingIndex] = nums[i]
        }
    }
    return trackingIndex+1
}

