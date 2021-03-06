/**
69. x 的平方根

实现 int sqrt(int x) 函数。

计算并返回 x 的平方根，其中 x 是非负整数。

由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

示例 1:

输入: 4
输出: 2
示例 2:

输入: 8
输出: 2
说明: 8 的平方根是 2.82842..., 
     由于返回类型是整数，小数部分将被舍去。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sqrtx
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main
import "fmt"
func main(){
    fmt.Println(mySqrt(121))
}

func mySqrt(x int) int {
    if x == 0 {
        return 0
    }
    var res int = 1
    var step int
    if x < 10 {
        step = 1
    }else if x < 100{
        step = 2
    }else{
        step = 10
    }
    for {
        if res * res / step == x / step {
            return res
        }else if res * res / step > x / step {
            res /= step
        }else{
            if (res+1)*(res+1)/step > x/step {
                return res
            }else if (res+1)*(res+1) == x {
                return res+1
            }
            res *= step
        }
        res++
    }
    return res
}
