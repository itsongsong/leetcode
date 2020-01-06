/**
给定一个非负整数数组 A，返回一个数组，在该数组中， A 的所有偶数元素之后跟着所有奇数元素。

你可以返回满足此条件的任何数组作为答案。

 

示例：

输入：[3,1,2,4]
输出：[2,4,3,1]
输出 [4,2,3,1]，[2,4,1,3] 和 [4,2,1,3] 也会被接受。
 

提示：

1 <= A.length <= 5000
0 <= A[i] <= 5000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-array-by-parity
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main
import "fmt"
func main(){
    fmt.Println(sortArrayByParity([]int{3,2,1,4}))
    fmt.Println(sortArrayByParity([]int{1,2,2,6,6,6,7,3,2,1,4}))
}
// 双指针
func sortArrayByParity(A []int) []int {
    leftIndex := 0
    rightIndex := len(A) - 1
    for leftIndex < rightIndex {
        if A[leftIndex] & 1 == 0 {
            leftIndex++
            continue
        } 
        if A[rightIndex] & 1 == 1 {
            rightIndex--
            continue
        }
        A[leftIndex], A[rightIndex] = A[rightIndex], A[leftIndex]
    }
    return A
}
/*
func sortArrayByParity(A []int) []int {
    for i:=0; i<len(A); i++{
        if A[i] & 1 == 0 {
            continue
        }
        for j:=len(A)-1; j>i; j--{
            if A[j] & 1 == 0 {
                A[i],A[j] = A[j],A[i] 
            }
        } 
    }
    return A
}
*/
