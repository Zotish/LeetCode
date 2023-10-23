func removeDuplicates(nums []int) int {
    if len(nums)==0{
        return 0
    }
     writePos,readPos := 1,1
    checkNum := nums[0]
    count := 1
     
    for readPos<len(nums) {
        if checkNum == nums[readPos] {
            count++
            if count <= 2 {
                nums[writePos] = nums[readPos]
                writePos++
            }
        } else {
            checkNum = nums[readPos]
            count = 1
            nums[writePos] = nums[readPos]
            writePos++
        }
        readPos++
    }

    return writePos
}
