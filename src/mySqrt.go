// https://leetcode.com/problems/sqrtx

func mySqrt(x int) int {
    if x <= 1 {
        return x
    }

    n := 0
    m := x

    for n < m {
        mid := n + (m-n) / 2
        square := mid*mid
        if square == x {
            return mid
        } else if square < x {
            n = mid + 1
        } else {
            m = mid - 1
        }
    }

    if m*m > x {
        return m-1
    }
    return m

}

