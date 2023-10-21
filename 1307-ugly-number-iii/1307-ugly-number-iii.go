func nthUglyNumber(n int, a int, b int, c int) int {
    left, right := 1, 2*int(1e9)
    lcmAB, lcmBC, lcmAC, lcmABC := lcm(a, b), lcm(b, c), lcm(a, c), lcm(a, lcm(b, c))
    for left < right {
        mid := (left + right) / 2
        count := mid/a + mid/b + mid/c - mid/lcmAB - mid/lcmBC - mid/lcmAC + mid/lcmABC
        if count < n {
            left = mid + 1
        } else {
            right = mid
        }
    }
    return left
}


func gcd(x, y int) int {
    if y == 0 {
        return x
    }
    return gcd(y, x%y)
}

func lcm(x, y int) int {
    return x * y / gcd(x, y)
}
