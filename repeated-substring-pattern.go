/**
给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。

示例 1:

输入: "abab"

输出: True

解释: 可由子字符串 "ab" 重复两次构成。
示例 2:

输入: "aba"

输出: False
示例 3:

输入: "abcabcabcabc"

输出: True

解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)

*/
package main 
import "fmt"
func main() {
    fmt.Println(repeatedSubstringPattern("abcabcabcabcabc"))
    fmt.Println(repeatedSubstringPattern("abaaba"))
    fmt.Println(repeatedSubstringPattern("aa"))
}
func repeatedSubstringPattern(s string) bool {
    var repeat string
    for i:=0; i<len(s)/2; i++ {
        repeat += string(s[i])
        if len(s) % len(repeat) != 0 {
            continue
        }
        if srepeat(repeat, len(s)/len(repeat)) != s {
            continue
        }else{
            return true
        }
    }
    return false
}
func srepeat(s string, n int) string {
    var res string
    for i:=0; i<n; i++{
        res+=s
    }
    return res
}
