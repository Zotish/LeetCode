func minOperations(nums []int, k int) int {
    values := make(map[int]struct{})
    moves := 0

    for i := len(nums) - 1; i >= 0; i-- {
        moves++
        if nums[i] <= k {
            values[nums[i]] = struct{}{}
        }
        if len(values) == k {
            return moves
        }
    }
    return 0
}
