func isHappy(n int) bool {
    sVisit := make(map[int]bool)

    for !sVisit[n] {
        sVisit[n] = true
        n = sumOfSquare(n)
        
        if n == 1 {
            return true
        }
    }
    return false
}

func sumOfSquare(n int) int {
    output := 0

    for n>0 {
        digit := n % 10
        digit  = digit * digit
        output += digit
        n = n / 10
    }
    return output
}