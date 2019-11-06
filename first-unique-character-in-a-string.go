/**
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

案例:

s = "leetcode"
返回 0.

s = "loveleetcode",
返回 2.
 

注意事项：您可以假定该字符串只包含小写字母。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/first-unique-character-in-a-string
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main
import "fmt"
func main(){
    fmt.Println(firstUniqChar("cc")) 
    fmt.Println(firstUniqChar("leetcode")) 
    fmt.Println(firstUniqChar("loveleetcode")) 
    fmt.Println(firstUniqChar("aadadaad")) 
}
func firstUniqChar(s string) int {
    m1 := make(map[byte]int)
    m2 := make(map[byte]int)
    for i:=len(s)-1;i>=0;i--{
        m1[s[i]] = i
    }
    for i:=0; i<len(s); i++{
        if _,ok:=m2[s[i]]; ok {
            m2[s[i]] = m2[s[i]] + 1
        }else{
            m2[s[i]] = 1
        }
    }
    var min int = -1
    for i,v := range m2 {
        if v == 1 {
            if min == -1 || min > m1[i]{
                min = m1[i]
            }
        }
    }
    return min
}
