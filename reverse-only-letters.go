/**
给定一个字符串 S，返回 “反转后的” 字符串，其中不是字母的字符都保留在原地，而所有字母的位置发生反转。

 

示例 1：

输入："ab-cd"
输出："dc-ba"
示例 2：

输入："a-bC-dEf-ghIj"
输出："j-Ih-gfE-dCba"
示例 3：

输入："Test1ng-Leet=code-Q!"
输出："Qedo1ct-eeLg=ntse-T!"
 

提示：

S.length <= 100
33 <= S[i].ASCIIcode <= 122 
S 中不包含 \ or "

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-only-letters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main
import "fmt"
func main(){
    fmt.Println(reverseOnlyLetters("a-bC-----------dEf-ghIj"))
    //fmt.Println(reverseOnlyLetters("aTj-Ih-gfE-dest1ng-Leet=code-Q!b-cTestCba1ng-Leet=code-Q!d"))
    //    fmt.Println(reverseOnlyLetters("Qedo1ct-eeLg=ntse-T!"))
    //    fmt.Println(reverseOnlyLetters("Qa-bC-dEf-ghIjedo1ct-eeLg=ntse-T!"))
}
func reverseOnlyLetters(S string) string {
    var leftIndex int = 0
    var rightIndex int = len(S) - 1
    for leftIndex < rightIndex {
        if check(S[leftIndex]) && check(S[rightIndex]){
            S = swap(S, leftIndex, rightIndex) 
            leftIndex++
            rightIndex--    
        }else if !check(S[rightIndex]) && check(S[leftIndex]){ 
            rightIndex--
        }else if check(S[rightIndex]) && !check(S[leftIndex]){ 
            leftIndex++
        }else if !check(S[leftIndex]) && !check(S[rightIndex]){
            leftIndex++
            rightIndex--
        }
    }
    return S
}
func check(b byte) bool {
    return (b>='a' && b<='z') || (b>='A' && b<='Z')
}
func swap(s string, i, j int) string {
    if i > j {
        i,j = j,i
    }
    if j == len(s) - 1 {
        return s[:i] + string(s[j]) + s[i+1:j] + string(s[i])
    }else{
        return s[:i] + string(s[j]) + s[i+1:j] + string(s[i]) + s[j+1:]
    }
}
