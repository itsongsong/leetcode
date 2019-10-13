/**

给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。



在杨辉三角中，每个数是它左上方和右上方的数的和。

示例:

输入: 5
输出:
[
[1],
[1,1],
[1,2,1],
[1,3,3,1],
[1,4,6,4,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/pascals-triangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main
import "fmt"

func main() {
	res := generate(5)
	fmt.Println(res)
}
func generate(numRows int) [][]int {
	var res = make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		res[i] = make([]int, i+1)
		res[i][0], res[i][i] = 1, 1
		if i < 1 {
			continue
		}
		for j:=1;j<i;j++{
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}
	return res
}

