// https://leetcode.com/problems/maximum-subarray

func maxSubArrayBruteForce(nums []int) int {
    sum := nums[0]
    size := 1
    for i := 0; i < len(nums); i++ {
        fmt.Println(size)
        for j := 0; j + size <= len(nums); j++ {
            localSum := 0
            for k := j; k < j+size; k++ {
                localSum += nums[k]
            }
            if sum < localSum {
                sum = localSum
            }
        }
        size++
    }
    return sum
}

func maxSubArrayDynamicMemoryAlloc(nums []int) int {
    n := len(nums)
    sumStore := make([][]int, n)
    for i := 0; i < n; i++ {
        sumStore[i] = make([]int, n)
    }
    
    sum := nums[0]
    
    for i := 0; i < n; i++ {
        sumStore[i][i] = nums[i]
        if sum < nums[i] {
            sum = nums[i]
        }
    }
    size := 1
    for size < n {
        j := 0
        for i := 0+size; i < n; i++ { 
            localSum := sumStore[j][i-1] + nums[i]
            sumStore[j][i] = localSum
            if sum < localSum {
                sum = localSum
            }
            j++
        }
        size++
    }
    return sum
}

