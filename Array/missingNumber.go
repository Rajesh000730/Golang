func missingNumber(nums []int) int {
    sum := 0
    n := len(nums)
    for i:=0; i<n;i++{
        sum+=nums[i]
    }
    nsum := n*(n+1)/2
    return nsum - sum
}
