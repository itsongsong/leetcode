/**
统计所有小于非负整数 n 的质数的数量。

示例:

输入: 10
输出: 4
解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
https://leetcode-cn.com/problems/count-primes/
*/

package main
import "fmt"
import "math"
func main(){
    fmt.Println("1:",countPrimes(1))
    fmt.Println("2:",countPrimes(2))
    fmt.Println("3:",countPrimes(3))
    fmt.Println("4:",countPrimes(4))
    fmt.Println("5:",countPrimes(5))
    fmt.Println("6:",countPrimes(6))
    fmt.Println("10:",countPrimes(10))
}
/**
解题思路：
1. 从1-n遍历每个数是否都是质数
2. 判断数字n是否是质数的条件：n能否被 2~n-1之间所有的质数所整除（不用遍历所有的数字）
3. 将质数存起来，直接返回质数集合的数量
*/
func countPrimes(n int) int {
    if n < 2 {
        return 0
    }
    var count int
    for x:=2; x < n; x++{
        var isPrise bool = true
        for i:=2; i<=int(math.Sqrt(float64(n))) && i <x; i++{
            if x % i == 0 {
                isPrise = false
                break
            }
        }
        if isPrise {
            count++
        }
    }
    return count
}
