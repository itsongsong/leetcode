/**
168. Excel表列名称

给定一个正整数，返回它在 Excel 表中相对应的列名称。

例如，

1 -> A
2 -> B
3 -> C
...
26 -> Z
27 -> AA
28 -> AB 
...
示例 1:

输入: 1
输出: "A"
示例 2:

输入: 28
输出: "AB"
示例 3:

输入: 701
输出: "ZY"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/excel-sheet-column-title
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main
import "fmt"
func main(){
    fmt.Println(convertToTitle(52))
    fmt.Println(convertToTitle(701))
    fmt.Println(convertToTitle(702))
    fmt.Println(convertToTitle(703))
    fmt.Println(convertToTitle(27))
    fmt.Println(convertToTitle(26))
    fmt.Println(convertToTitle(18278))
    fmt.Println(convertToTitle(18279))
}

func convertToTitle(n int) string { 
    var s string
    for n > 0{
        n--
        s = string(n%26+'A') + s
        n/=26
    }
    return s
}
