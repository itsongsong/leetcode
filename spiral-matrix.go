/*
给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

示例 1:

输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
输出: [1,2,3,6,9,8,7,4,5]
示例 2:

输入:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
输出: [1,2,3,4,8,12,11,10,9,5,6,7]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/spiral-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main
import "fmt"
func main(){
    matrix := [][]int{
    }
    fmt.Println(spiralOrder(matrix))
}

func spiralOrder(matrix [][]int) []int {
    n := len(matrix)    // 4
    if(n == 0) {
        return []int{}
    }
    m := len(matrix[0]) // 3
    res := make([]int,0, m*n)
    start_m := 0
    start_n := 0
    var max int = m * n
    for {
        res = append(res, m1(start_m, m, start_n, matrix, res)...)
        start_n++
        m--
        if len(res) >=max {break}
        res = append(res, n1(start_n, n, m, matrix, res)...)
        n--
        if len(res) >=max {break}
        res = append(res, m2(m-1, start_m, n, matrix, res)...)
        
        if len(res) >=max {break}
        res = append(res, n2(n-1, start_n, start_m, matrix, res)...)
        start_m++
        if len(res) >=max {break}
    }
    return res
}
func m1(start_m int, end_m int, start_n int, matrix [][]int, res []int) []int{
    xx := []int{}
    for i:=start_m; i<end_m; i++ {
        xx = append(xx,matrix[start_n][i])
    }
    return xx
}
func n1(start_n int, end_n int, end_m int, matrix [][]int ,res []int) []int{
    xx := []int{}
    for i:=start_n; i<end_n; i++ {
        xx = append(xx,matrix[i][end_m])
    }
    return xx
}
func m2(end_m int, start_m int, n int, matrix [][]int, res []int) []int{
    xx := []int{}
    for i:=end_m; i>=start_m; i-- {
        xx = append(xx,matrix[n][i])
    }
    return xx
}
func n2(end_n int, start_n int, m int, matrix [][]int, res []int) []int{
    xx := []int{}
    for i:=end_n; i>=start_n; i-- {
        xx = append(xx,matrix[i][m])
    }
    return xx
}
