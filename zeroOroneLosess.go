import "sort"
func findWinners(matches [][]int) [][]int {
    var winners []int
    var losers []int
    var ans [][]int
    temp := [100001]int{}
    n := len(matches)
    for i:=0;i<n;i++{
        temp[matches[i][1]] +=1
    }
    for j:=0;j<100001;j++{
        if temp[j]==1{
            losers = append(losers, j)
        }
    }
    for k:=0;k<n;k++{
        if temp[matches[k][0]] == 0{
            winners = append(winners, matches[k][0])
            temp[matches[k][0]] = 1
        }
    }
    sort.Ints(winners)
    sort.Ints(losers)
    ans = append(ans, winners)
    ans = append(ans, losers)
    return ans
}
