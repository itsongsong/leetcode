/**
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main
import "fmt"

func main(){
    nums := []int{10,20,30,40,50}
    res := twoSum(nums, 80)
    fmt.Println(res)
}
func twoSum(nums []int, target int) []int {
    m := make(map[int]int, len(nums))
    for index, value := range nums {
        if j, ok := m[target-value]; ok {
            return []int{j, index}
        }else{
            m[value] = index
        }
    }
    return nil
}
