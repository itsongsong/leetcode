/**
给定一个字符串，你需要反转字符串中每个单词的字符顺序，同时仍保留空格和单词的初始顺序。

示例 1:

输入: "Let's take LeetCode contest"
输出: "s'teL ekat edoCteeL tsetnoc" 
注意：在字符串中，每个单词由单个空格分隔，并且字符串中不会有任何额外的空格。



来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-words-in-a-string-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main
import(
    "fmt"
)
func main(){
    fmt.Println(reverseWords("Let's take LeetCode contest"))
}
func reverse(s []byte, start, end int) {
    for start < end{
        s[start], s[end] = s[end], s[start]
        start++
        end--
    }
}

func reverseWords(s string) string {
    words := []byte(s)
    start := 0
    for end, word := range words {
        if word == ' ' {
            reverse(words, start, end-1)
            start = end + 1
        }
    }
    reverse(words, start, len(words)-1)
    return string(words)
}
