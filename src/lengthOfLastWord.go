// https://leetcode.com/problems/length-of-last-word

func lengthOfLastWord(s string) int {
    runeS := []rune(s)
    
    if len(runeS) == 0 { return 0 }
    if len(runeS) == 1 && runeS[0] == rune(' ') { return 0 }
    globalCount := 0
    wordCount := 0
    for i := 0; i < len(runeS); i++ {
        if runeS[i] != rune(' ') {
            wordCount++
        } else {
            if globalCount < wordCount {
                globalCount = wordCount
            }
            wordCount = 0
        }
    }
    if globalCount < wordCount { return wordCount}
    return globalCount
}

