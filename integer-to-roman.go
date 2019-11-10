/**
罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

    I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
    X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。 
    C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。

给定一个整数，将其转为罗马数字。输入确保在 1 到 3999 的范围内。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/integer-to-roman
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main
import "fmt"
func main(){
    fmt.Println(intToRoman(149))
}
func intToRoman(num int) string {
    m := map[int]string{1:"I",5:"V",10:"X",50:"L",100:"C",500:"D",1000:"M"}
    var s string
    var degree int = 1
    for num > 0 {
        temp := num % 10
        shang := temp / 5
        yushu := temp % 5
        if yushu != 4 {
            s = repeat(m[degree], yushu) + s
            if shang == 1 {
                s = m[5*degree] + s 
            }
        }else{
            if shang == 1 {
                s = m[degree] + m[degree*10] + s
            }else{
                s = m[degree] + m[degree*5] + s
            }
        }
        num /= 10
        degree *= 10
    }
    return s
}
func repeat(s string, n int) string{
    var res string
    for i:=0; i<n; i++{
       res += s 
    }
    return res
}
