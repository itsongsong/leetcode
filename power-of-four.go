/**
给定一个整数 (32 位有符号整数)，请编写一个函数来判断它是否是 4 的幂次方。

示例 1:

输入: 16
输出: true
示例 2:

输入: 5
输出: false
进阶：
你能不使用循环或者递归来完成本题吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/power-of-four
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main
import (
    "fmt";
    "math";
)
func main(){
    fmt.Println(isPowerOfFour(8))
    fmt.Println(isPowerOfFour(4))
    fmt.Println(isPowerOfFour(2))
}
/**
解题思路：
1. 2的幂次方首位为1，其余位都为1，所以num和num-1求与结果是否为9
2. 由于2的幂次方中只有4的幂次方才能求平方根，由此过滤掉只为2的幂次方
3. 处理特殊数字
*/
func isPowerOfFour(num int) bool {
    if num <= 0 {
        return false
    }
    if num == 1 {
        return true
    }
    return num & (num-1) == 0 && x(num)
}
func x(n int) bool{
    temp := int(math.Sqrt(float64(n)))
    return temp * temp == n
}
