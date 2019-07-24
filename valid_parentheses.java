import java.util.*;
class Solution {
    /**
     * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
     * 有效字符串需满足：
     * 左括号必须用相同类型的右括号闭合。
     * 左括号必须以正确的顺序闭合。
     * 注意空字符串可被认为是有效字符串。
     * 示例 1:
     * 输入: "()"
     * 输出: true
     * 示例 2:
     * 输入: "()[]{}"
     * 输出: true
     * 示例 3:
     * 输入: "(]"
     * 输出: false
     * 示例 4:
     * 输入: "([)]"
     * 输出: false
     * 示例 5:
     * 输入: "{[]}"
     * 输出: true
     * 来源：力扣（LeetCode）
     * 链接：https://leetcode-cn.com/problems/valid-parentheses
     * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
    */
    /**
     * 执行用时 :7 ms, 在所有 Java 提交中击败了48.34%的用户
     * 内存消耗 :35.2 MB, 在所有 Java 提交中击败了82.05%的用户
    */
    public boolean isValid2(String s) {
        Map<String, String> map= new HashMap<>();
        map.put("(", ")");       
        map.put("[", "]");       
        map.put("{", "}");       
        Stack<String> st = new Stack<>();
        for(int i = 0;i < s.length(); i++) {
            String ss = s.substring(i, i+1);
            if(!st.empty() && map.getOrDefault(st.peek(), "").equals(ss)) {
                st.pop();
            }else{
                st.push(ss);
            }
        }
        if(st.empty()) {
            return true;
        }
        return false;
    }
   /**
     * 执行用时 :6 ms, 在所有 Java 提交中击败了66.23%的用户
     * 内存消耗 :34.6 MB, 在所有 Java 提交中击败了83.04%的用户
    */
    public boolean isValid(String s) {
        Map<Character, Character> map= new HashMap<>();
        map.put('(', ')');
        map.put('[', ']');
        map.put('{', '}');
        Stack<Character> st = new Stack<>();
        char[] sa = s.toCharArray();
        for(int i = 0;i < sa.length; i++) {
            Character ss = sa[i];
            if(!st.empty() && map.getOrDefault(st.peek(), 'a').equals(ss)) {
                st.pop();
            }else{
                st.push(ss);
            }
        }
        if(st.empty()) {
            return true;
        }
        return false;
    }


    public static void main(String[] args) {
        String s = "([)]";
        Solution solution = new Solution();
        boolean result = solution.isValid(s);
        System.out.println(result);
    }
}
