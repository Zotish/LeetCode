func sumOddLengthSubarrays(arr []int) int {
    n := len(arr)
    sum := 0

    for i := 0; i < n; i++ {
        for j := i; j < n; j++ {
            if (j-i+1)%2 == 1 {
                for k := i; k <= j; k++ {
                    sum += arr[k]
                }
            }
        }
    }

    return sum
}
