func minOperations(nums []int) int {
    freq := make(map[int]int)
    for _, num := range nums {
        freq[num]++
    }

    ans := 0
    for _, val := range freq {
        if val == 1 {
            return -1
        }
        ans += val / 3
        if val % 3 != 0 {
            ans++
        }
    }
    return ans
}