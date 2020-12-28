// https://leetcode.com/problems/implement-strstr

func strStr(haystack string, needle string) int {
    needle_r := []rune(needle)
    haystack_r := []rune(haystack)
    
    if len(needle_r) == 0 {
        return 0
    }
    if haystack == needle {
        return 0
    }
    
    match := false
    var i int
    for i = 0; i < len(haystack_r); i++ {
        if !match {
            if len(haystack_r) - i < len(needle_r) {
                break
            }
            for j := 0; j < len(needle_r); j++ {
                fmt.Println(i, j)
                if haystack_r[i+j] != needle_r[j] {
                    break
                }
                if j == len(needle_r)-1 {
                    match = true
                }
            }
        } else {break}
    }
    if match {
        return i -1
    } else {
        return -1
    }
}

