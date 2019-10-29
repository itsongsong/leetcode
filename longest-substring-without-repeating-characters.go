/**
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3 
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package main
import "fmt"
//import "time"
func main(){
    fmt.Println(lengthOfLongestSubstring("abac"))
    fmt.Println(lengthOfLongestSubstring("abcabcde"))
    fmt.Println(lengthOfLongestSubstring("pwwkew"))
    fmt.Println(lengthOfLongestSubstring("erterdertf"))
}

func lengthOfLongestSubstring(s string) int{
    m := make(map[string]int)
    var maxCount int
    for i:=0; i<len(s); i++{
        //time.Sleep(time.Second)
        //fmt.Println(m,i)
        if  v:=m[string(s[i])];v>0 {
            //fmt.Println("充值点：", v)
            if maxCount < len(m) {
                maxCount = len(m)
            }
            i = v-1
            m = make(map[string]int)
        }else{
            m[string(s[i])] = i + 1
        }
    }
    if maxCount > len(m) {
        return maxCount
    }else{
        return len(m)
    }
}
