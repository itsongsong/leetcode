/*
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum-closest
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main

import "fmt"

func main() {
    fmt.Println(threeSumClosest([]int{1,1,1,0,0},-100))
    fmt.Println(threeSumClosest([]int{1,2,4,8,16,32,64,128},82))
    fmt.Println(threeSumClosest([]int{1,-4,-1,2},1))
    fmt.Println(threeSumClosest([]int{1,1,-1,-1,3},-1))
}

func threeSumClosest(nums []int, target int) int{
    var (
        res int
        start int
        end int
        temp int
    )
    for i:=0;i<len(nums);i++{
        for j:=i;j<len(nums);j++{
            if nums[j] < nums[i] {
                nums[j], nums[i] = nums[i], nums[j]
            }
        }
    }

    res = nums[0] + nums[1] + nums[2]
    for i:=0;i<len(nums);i++{
        start = i+1
        end = len(nums) -1
        for start < end {
            temp = nums[i] + nums[start] + nums[end]
            if (target-temp) * (target-temp) < (target-res) * (target-res) {
                res = temp            
            }else {
                if temp > target {
                    end--
                }else if temp < target{
                    start++
                }else {
                    return res
                }
            }
        }
    }
    return res
}
