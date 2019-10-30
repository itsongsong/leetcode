/**
67. 二进制求和

给定两个二进制字符串，返回他们的和（用二进制表示）。

输入为非空字符串且只包含数字 1 和 0。

示例 1:

输入: a = "11", b = "1"
输出: "100"
示例 2:

输入: a = "1010", b = "1011"
输出: "10101"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-binary
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main
import "fmt"
func main() {
    fmt.Println(addBinary("10", ""))
    fmt.Println(addBinary("1010", "1011"))
}

func addBinary(a string, b string) string {
    var mins string
    var maxs string
    var s string
    var index bool

    if len(a) > len(b) {
        mins = b
        maxs = a
    }else{
        mins = a
        maxs = b
    }

    var minus = len(maxs)-len(mins)
    for i:=0; i<minus; i++ {
        mins = "0"+mins
    }

    for i:=0; i<len(mins); i++{
        if mins[len(mins) - 1 - i] == '0' && maxs[len(maxs) - 1 - i] == '0' {
            if index{
                s = "1" + s
            }else{
                s = "0" + s
            }
            index = false
        }else if (mins[len(mins) - 1 - i] == '1' && maxs[len(maxs) - 1 - i] == '1') {
            if index {
                s = "1" + s
            }else{
                s = "0" + s
            }
            index = true
        }else{
            if index {
                s = "0" + s
                index = true
            }else{
                s = "1" + s
            }
        }
    }
    if index {
        s = "1" + s
    }
    return s
}
