/**
给定两个数组，编写一个函数来计算它们的交集。

示例 1:

输入: nums1 = [1,2,2,1], nums2 = [2,2]
输出: [2]
示例 2:

输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
输出: [9,4]
说明:

输出结果中的每个元素一定是唯一的。
我们可以不考虑输出结果的顺序。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/intersection-of-two-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main
import "fmt"
func main(){
    fmt.Println(intersection([]int{1,2,3,3}, []int{2,3,3,4}))
}
/**
解题思路：
1. 将其中的一个数组遍历成map，同时也做了去重操作
2. 遍历另外一个数组，找到map中一致的元素存入结果集
时间复杂度：O(n); 空间复杂度：O(n)
*/
func intersection(nums1 []int, nums2 []int) []int {
    m:=make(map[int]int, len(nums1))
    res := make([]int, 0, len(nums1))
    for i:=0; i<len(nums1); i++{
        m[nums1[i]] = nums1[i]
    }
    for i:=0; i<len(nums2); i++{
        if v,ok:=m[nums2[i]]; ok{
            res = append(res, v)
            delete(m, v)
        }
    }
    return res
}
