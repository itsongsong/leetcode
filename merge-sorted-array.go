/**
给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。

说明:

初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
示例:

输入:
nums1 = [1,2,3,0,0], m = 3
nums2 = [2,5,6]       n = 3

输出: [1,2,2,3,5,6]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。    
*/
package main
import "fmt"
func main(){
    a := []int{2,0}
    b := []int{1}
    merge(a, 1, b, 1)
    fmt.Println(a)
}


func merge(nums1 []int, m int, nums2 []int, n int)  {
    if len(nums1) == 0 || len(nums2) == 0 {
        return
    }
    nums3 := make([]int, len(nums1))
    copy(nums3, nums1)
    i,j,index := 0,0,0
    for i < m && j < n {
        if  nums3[i] < nums2[j]{
            nums1[index] = nums3[i]
            i++
        }else{
            nums1[index] = nums2[j]
            j++
        }
        index++
    }
    for i < m {
        nums1[index] = nums3[i]
        i++
        index++
    }
    for j < n {
        nums1[index] = nums2[j]
        j++
        index++
    }
}
