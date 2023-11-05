func countNumbersWithUniqueDigits(n int) int {
    if n == 0 {
        return 1
    }
    
    current := 9
    for i := 1; i<n; i++ {
        current *= (10-i)
    }
    return current + countNumbersWithUniqueDigits(n-1)
}