// https://leetcode.com/problems/remove-element

func removeElement(nums []int, val int) int {
    trackIndex := 0
    
    for i := 0; i < len(nums); i++ {
        if nums[i] != val {
            if nums[trackIndex] == nums[i] {
                trackIndex += 1
            } else {
                nums[trackIndex] = nums[i]
                nums[i] = val
                trackIndex += 1
            }
        }
    }
    return trackIndex
}

