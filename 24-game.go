/**
你有 4 张写有 1 到 9 数字的牌。你需要判断是否能通过 *，/，+，-，(，) 的运算得到 24。

示例 1:

输入: [4, 1, 8, 7]
输出: True
解释: (8-4) * (7-1) = 24

示例 2:

输入: [1, 2, 1, 2]
输出: False

注意:

    除法运算符 / 表示实数除法，而不是整数除法。例如 4 / (1 - 2/3) = 12 。
    每个运算符对两个数进行运算。特别是我们不能用 - 作为一元运算符。例如，[1, 1, 1, 1] 作为输入时，表达式 -1 - 1 - 1 - 1 是不允许的。
    你不能将数字连接在一起。例如，输入为 [1, 2, 1, 2] 时，不能写成 12 + 12 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/24-game
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main
import "fmt"
//import "time"
func main() {
    fmt.Println(judgePoint24([]int{5, 1, 1, 9}))
    fmt.Println(judgePoint24([]int{4, 1, 8, 7}))
    fmt.Println(judgePoint24([]int{3, 3, 8, 8}))

}
func judgePoint24(nums []int) bool{
    var xx []float64
    for i:=0;i<len(nums);i++{
        xx = append(xx,float64(nums[i]))
    }
    return judgePoint242(xx)
}
func judgePoint242(nums []float64) bool {
    if len(nums) == 1 {
        if nums[0] - 24 > -0.0001 && nums[0]-24 < 0.0001 {
            return true
        }else{
            return false
        }
    }
    for i:=0; i<len(nums); i++ {
        for j:=0; j<len(nums); j++ {
            if i==j {
                continue
            }
            var collect []float64
            for m:=0; m<len(nums); m++{
                if m!=i && m != j {
                    collect = append(collect, float64(nums[m]))
                }
            }
            
            for n:=0;n<4;n++{
                var temp = collect
                if n == 0 {
                    temp = append(temp, float64(nums[i]) + float64(nums[j]))
                }else if n == 1 {
                    temp = append(temp, float64(nums[i]) - float64(nums[j]))
                }else if n == 2 {
                    temp = append(temp, float64(nums[i]) * float64(nums[j]))
                }else if n ==3 {
                    if nums[j] == 0 {
                        if len(temp) == 2 {
                            return false
                        }else{
                            continue
                        }
                    }else {
                        temp = append(temp, float64(nums[i]) / float64(nums[j]))
                    }
                }
                if judgePoint242(temp){
                    return true
                }
            }

        }
    }
    return false
}
