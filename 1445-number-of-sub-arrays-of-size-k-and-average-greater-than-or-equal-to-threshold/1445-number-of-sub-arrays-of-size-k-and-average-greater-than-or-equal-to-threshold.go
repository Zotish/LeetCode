func numOfSubarrays(arr []int, k int, threshold int) int {
    l, ans, sum := 0, 0, 0
    target := k * threshold  // total sum required to meet or exceed the threshold

    for r := 0; r < len(arr); r++ {
        sum += arr[r]

        if r - l + 1 == k {
            if sum >= target {
                ans++
            }
            sum -= arr[l]
            l++
        }
    }

    return ans
}
