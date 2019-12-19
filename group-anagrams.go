/**
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

示例:

输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
说明：

所有输入均为小写字母。
不考虑答案输出的顺序。
在真实的面试中遇到过这道题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/group-anagrams
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import "fmt"

func main() {
    strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
    fmt.Println(groupAnagrams(strs))
}

func groupAnagrams(strs []string) [][]string {
    m := make(map[string][]string)
    for i:=0; i<len(strs); i++ {
        key := sort(strs[i])
        m[key] = append(m[key], strs[i])
    }

    res := make([][]string, 0, len(m))
    for _,v := range m {
        res = append(res, v)
    }
    return res
}

func sort(str string) string{
    for i:=0; i<len(str); i++ {
        for j:=i+1; j<len(str); j++ {
            if str[i] > str[j] {
                str = str[:i] + string(str[j]) + str[i+1:j] + string(str[i]) + str[j+1:]
            }
        }    
    } 
    return str
}
