/**
125. 验证回文串

给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:

输入: "race a car"
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-palindrome
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main
import "fmt"
import "strings"
func main(){
    fmt.Println(isPalindrome("A man, a plan, a 1c1anal: Panama"))
    fmt.Println(isPalindrome(""))
}
func isPalindrome(s string) bool {
    s = strings.ToUpper(s)
    var i int = 0
    var j int = len(s) - 1
    for i < j {
        if !valid(s[i]){
            i++
            continue
        }
        if !valid(s[j]) {
            j--
            continue
        }
        if s[i] != s[j] {
            return false
        }else{
            i++
            j--
        }
    }
    return true
}

func valid(c byte) bool {
    if (c>='A' && c<='Z') || (c>='a' && c<='z') || (c>='0' && c<='9') {
        return true
    }else{
        return false
    }
}
