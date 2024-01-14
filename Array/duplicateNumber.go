func findDuplicate(nums []int) int {
    slow:=0
    fast := 0
    for{
        slow = nums[slow]
        fast = nums[nums[fast]]

        if nums[slow] == nums[fast]{
            break;
        }
    }
    slow = 0
    for slow != fast{
        slow = nums[slow]
        fast = nums[fast]
    }
    return slow
}
