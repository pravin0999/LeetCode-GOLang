// https://leetcode.com/problems/add-binary


func addBinary(a string, b string) string {
    if len(a) < len(b) {
        a, b = b, a
    }
    solution := make([]byte, len(a)+1)
    var carry byte
    for i, j := (len(a)-1), (len(b)-1); i >= 0 || j >= 0; {
        var aByte, bByte byte
        if i >= 0 { aByte = a[i] - '0' }
        if j >= 0 { bByte = b[j] - '0' }

        sum := aByte ^ bByte ^ carry
        carry = aByte & bByte | carry & (aByte ^ bByte)

        solution[i+1] = '0' + sum

        i, j = i - 1, j - 1
    }
    if carry == 1 {
        solution[0] = '1'
        return string(solution)
    } else {
        return string(solution[1:])
    }
}

