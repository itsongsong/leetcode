/**
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。

示例 1:

输入: [1,2,3,1]
输出: 4
解释: 偷窃 1 号房屋 (金额 = 1) ，然后偷窃 3 号房屋 (金额 = 3)。
     偷窃到的最高金额 = 1 + 3 = 4 。
示例 2:

输入: [2,7,9,3,1]
输出: 12
解释: 偷窃 1 号房屋 (金额 = 2), 偷窃 3 号房屋 (金额 = 9)，接着偷窃 5 号房屋 (金额 = 1)。
     偷窃到的最高金额 = 2 + 9 + 1 = 12 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/house-robber
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main
import "fmt"
func main(){
    fmt.Println(rob([]int{114,117,207,117,235,82,90,67,143,146,53,108,200,91,80,223,58,170,110,236,81,90,222,160,165,195,187,199,114,235,197,187,69,129,64,214,228,78,188,67,206,94,205,169,241,202,144,240}))
    fmt.Println(rob([]int{2,7,9,3,1}))
    fmt.Println(rob([]int{1,2,3,1}))
}
func rob(nums []int) int {
    m := make(map[int]int)

    return xx(nums, m)
}

func xx(nums []int, m map[int]int) int {
    if len(nums) == 1 {
        m[len(nums)] = nums[0]
        return nums[0]
    }
    if len(nums) == 0 {
        m[len(nums)] = 0
        return 0
    }

    var r1,r2 int
    if v,ok := m[len(nums)-1];ok {
        r1=v
    }else{
        r1=xx(nums[:len(nums)-1], m)
    }
    
    if v,ok := m[len(nums)-2]; ok{
        r2=v+nums[len(nums)-1]
    }else{
        r2=xx(nums[:len(nums)-2], m) + nums[len(nums)-1]
    }
    //r1:=xx(nums[:len(nums)-1], m)
    //r2:=xx(nums[:len(nums)-2], m) + nums[len(nums)-1]

    if r1 > r2 {
        m[len(nums)] = r1
        return r1
    }else{
        m[len(nums)] = r2
        return r2
    }

}
