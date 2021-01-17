// https://leetcode.com/problems/climbing-stairs

func climbStairs(n int) int {
    count := 0
    getCombinations(n, []int{1, 2}, &count)
    return count
}

func getCombinations(n int, steps []int, count *int) {
    if n == 0 {
        *count++
        return
    }
    if n < 0 {
        return
    }
    for i := 0; i < len(steps); i++ {
        tempN := n - steps[i]
        getCombinations(tempN, steps, count)
    }
}

