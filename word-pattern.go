/**
给定一种规律 pattern 和一个字符串 str ，判断 str 是否遵循相同的规律。

这里的 遵循 指完全匹配，例如， pattern 里的每个字母和字符串 str 中的每个非空单词之间存在着双向连接的对应规律。

示例1:

输入: pattern = "abba", str = "dog cat cat dog"
输出: true
示例 2:

输入:pattern = "abba", str = "dog cat cat fish"
输出: false
示例 3:

输入: pattern = "aaaa", str = "dog cat cat dog"
输出: false
示例 4:

输入: pattern = "abba", str = "dog dog dog dog"
输出: false
说明:
你可以假设 pattern 只包含小写字母， str 包含了由单个空格分隔的小写字母

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-pattern
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main
import "fmt"
func main(){
    fmt.Println(wordPattern("abb", "dog dog  dog")) 
}
/**
解题思路
1. 利用双循环
2. 循环str得到按照空格分组的数组
3. 循环pattern得到对应的键值
*/
func wordPattern(pattern string, str string) bool {
    m:=make(map[byte]string, len(pattern))
    m2:=make(map[string]byte, len(pattern))
    s := explode(' ', str)
    if len(s) != len(pattern) {
        return false
    }
    for i:=0; i<len(pattern); i++{
        if v,ok:=m[pattern[i]]; !ok {
            m[pattern[i]] = s[i]
        }else{
            if v != s[i] {
                return false
            }
        }
    }
    for i:=0; i<len(s); i++{
        if v,ok:=m2[s[i]]; !ok {
            m2[s[i]] = pattern[i]
        }else{
            if v != pattern[i] {
                return false
            }
        }
    }
    return true
}

func explode(delimeter byte, str string) []string {
    s := make([]string, 0, len(str)/2+1)
    var (
        index int
    )
    for i:=0; i<len(str); i++{
        if str[i] == delimeter {
            if i > 0 && str[i-1] != delimeter && len(s[index]) > 0{
                index = index+1
            }
            continue
        }else{
            if len(s) == 0 || len(s) == index {
                s = append(s, string(str[i]))
            }else{
                s[index] = s[index] + string(str[i])
            }
        }
    }
    return s
}
