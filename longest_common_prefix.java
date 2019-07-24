import java.util.*;
class Solution {
    /**
     * 编写一个函数来查找字符串数组中的最长公共前缀。
     * 如果不存在公共前缀，返回空字符串 ""。
     * 示例 1:
     * 输入: ["flower","flow","flight"]
     * 输出: "fl"
     * 示例 2:
     * 输入: ["dog","racecar","car"]
     * 输出: ""
     * 解释: 输入不存在公共前缀。
     * 说明:
     * 所有输入只包含小写字母 a-z 。
     * 来源：力扣（LeetCode）
     * 链接：https://leetcode-cn.com/problems/longest-common-prefix
     * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
     */
    /**
     * 执行用时 :11 ms, 在所有 Java 提交中击败了7.85%的用户
     * 内存消耗 :38 MB, 在所有 Java 提交中击败了64.35%的用户
     */
    public String longestCommonPrefix(String[] strs) {
        if(strs.length == 0){
            return "";
        }
        // 两遍遍历
        String result = strs[0];
        char[] standTemp = result.toCharArray();
        for(int i = 1; i < strs.length; i++){
            result = "";
            char[] cs = strs[i].toCharArray();
            int minLength = cs.length > standTemp.length ? standTemp.length : cs.length;
            for(int j = 0; j < minLength; j++){
               if(cs[j] != standTemp[j]) {
                    standTemp = result.toCharArray();
                    break;
               } 
               result += cs[j];
            }
            standTemp = result.toCharArray();
        }
        return result;
    }
    

    public static void main(String args[]) {
        String[] strs = {"flower","flow","flight"};
        // System.out.println(strs[0].toCharArray().length);

        Solution solution = new Solution();
        String result = solution.longestCommonPrefix(strs);
        System.out.println(result);
    }
}
