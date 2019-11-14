/**
给定正整数 n，找到若干个完全平方数（比如 1, 4, 9, 16, ...）使得它们的和等于 n。你需要让组成和的完全平方数的个数最少。

示例 1:

输入: n = 12
输出: 3 
解释: 12 = 4 + 4 + 4.
示例 2:

输入: n = 13
输出: 2
解释: 13 = 4 + 9.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/perfect-squares
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main
import (
    "fmt";
)
func main(){
    fmt.Println(numSquares(13))
}
/**
解题思路：
1. 对应n来说，最坏的情况是n个1相加
2. 从1开始到n遍历出每个数字最小平方和的个数
3. 1~n的遍历中，需要比较1~j的每个组合最小平方和的个数（j为n-j*j>0的最大值）
*/
func numSquares(n int) int {
    m := map[int]int{}
    for i:=1; i<=n; i++{
        m[i] = i
        for j:=1;i-j*j>=0;j++{
            m[i] = min(m[i], m[i-j*j]+1)
        }
    }   
    return m[n]
}
func min(a int, b int) int{
    if a>b{
        return b
    }
    return a
}
