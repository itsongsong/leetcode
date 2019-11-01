/**
实现一个基本的计算器来计算一个简单的字符串表达式的值。

字符串表达式仅包含非负整数，+， - ，*，/ 四种运算符和空格  。 整数除法仅保留整数部分。

示例 1:

输入: "3+2*2"
输出: 7
示例 2:

输入: " 3/2 "
输出: 1
示例 3:

输入: " 3+5 / 2 "
输出: 5
说明：

你可以假设所给定的表达式都是有效的。
请不要使用内置的库函数 eval。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/basic-calculator-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main
import "fmt"
import "strconv"
func main(){
    fmt.Println(calculate("1+23/4 + 5 - 6*78/9 "))
}

func calculate(s string) int {
    nums := make([]string,0, len(s)) 
    chars := make([]byte,0, len(s))
    var flag = true
    for i:=0; i<len(s); i++{
        if s[i] == ' '{
            continue
        }else if s[i] == '+' || s[i] == '-' || s[i] == '*' || s[i] == '/' {
            chars = append(chars, s[i])
            flag = false
        }else{
            if flag {
                if len(nums) == 0 {
                    nums = append(nums, string(s[i]))
                } else {
                    nums[len(nums)-1] = nums[len(nums)-1] + string(s[i])
                }
            }else{
                nums = append(nums, string(s[i]))
                flag = true
            }
        }
    }
    for i:=0; i<len(chars); i++ {
        if chars[i] == '*' || chars[i] == '/' {
            nums[i+1], nums[i] = strconv.Itoa(cal(nums[i], nums[i+1], chars[i])), "0"
            if i == 0 {
                chars[i] = '+'
            }else{
                chars[i] = chars[i-1]
            }
        }
    }
    for i:=0; i<len(chars); i++ {
        nums[i+1] = strconv.Itoa(cal(nums[i], nums[i+1], chars[i])) 
    }
    res,_ := strconv.Atoi(nums[len(nums)-1])
    return res
}
func cal(a1 string,b1 string,c byte)int {
    a,_ := strconv.Atoi(a1)
    b,_ := strconv.Atoi(b1)
    switch c {
        case '+':
            return a+b
        case '-':
            return a-b
        case '*':
            return a*b
        case '/':
            return a/b
    }
    return 0
}


