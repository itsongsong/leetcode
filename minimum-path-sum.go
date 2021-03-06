/**
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例:

输入:
[
  [1,3,1],
[1,5,1],
[4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main
import "fmt"
func main(){
    grid0 := [][]int{{1,3,11},{9,5,1}}
    grid1 := [][]int{{1,3,1},{1,5,1},{4,2,1}}
    grid2 := [][]int{{1},{5},{2}}
    grid3 := [][]int{{1,3,14,2,1}}
    grid4 := [][]int{{3,8,6,0,5,9,9,6,3,4,0,5,7,3,9,3},{0,9,2,5,5,4,9,1,4,6,9,5,6,7,3,2},{8,2,2,3,3,3,1,6,9,1,1,6,6,2,1,9},{1,3,6,9,9,5,0,3,4,9,1,0,9,6,2,7},{8,6,2,2,1,3,0,0,7,2,7,5,4,8,4,8},{4,1,9,5,8,9,9,2,0,2,5,1,8,7,0,9},{6,2,1,7,8,1,8,5,5,7,0,2,5,7,2,1},{8,1,7,6,2,8,1,2,2,6,4,0,5,4,1,3},{9,2,1,7,6,1,4,3,8,6,5,5,3,9,7,3},{0,6,0,2,4,3,7,6,1,3,8,6,9,0,0,8},{4,3,7,2,4,3,6,4,0,3,9,5,3,6,9,3},{2,1,8,8,4,5,6,5,8,7,3,7,7,5,8,3},{0,7,6,6,1,2,0,3,5,0,8,0,8,7,4,3},{0,4,3,4,9,0,1,9,7,7,8,6,4,6,9,5},{6,5,1,9,9,2,2,7,4,2,7,2,2,3,7,2},{7,1,9,6,1,2,7,0,9,6,6,4,4,5,1,0},{3,4,9,2,8,3,1,2,6,9,7,0,2,4,2,0},{5,1,8,8,4,6,8,5,2,4,1,6,2,2,9,7}}
    fmt.Println(minPathSum(grid0))
    fmt.Println(minPathSum(grid1))
    fmt.Println(minPathSum(grid2))
    fmt.Println(minPathSum(grid3))
    fmt.Println(minPathSum(grid4))
}
func minPathSum(grid [][]int) int {
    for i:=0; i<len(grid); i++{
        for j:=0; j<len(grid[0]); j++{
            if i==0 && j>0 {
                grid[i][j] = grid[i][j-1] + grid[i][j]
            }else if i>0 && j==0{
                grid[i][j] = grid[i-1][j] + grid[i][j]
            }else if i>0 && j>0 {
                grid[i][j] = min(grid[i-1][j], grid[i][j-1]) + grid[i][j]
            }
        }
    }
    return grid[len(grid)-1][len(grid[0])-1]
}
func min(a int, b int) int {
    if a>b{
        return b
    }else{
        return a
    }
}
