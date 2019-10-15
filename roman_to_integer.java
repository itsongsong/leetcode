import java.util.*;
class Solution {
    /**
     * 罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
     * 字符          数值
     * I             1
     * V             5
     * X             10
     * L             50
     * C             100
     * D             500
     * M             1000
     * 例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。
     * 通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：
     * I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
     * X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。 
     * C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
     * 给定一个罗马数字，将其转换成整数。输入确保在 1 到 3999 的范围内。
     * 示例 1:
     * 输入: "III"
     * 输出: 3
     * 示例 2:
     * 输入: "IV"
     * 输出: 4
     * 示例 3:
     * 输入: "IX"
     * 输出: 9
     * 示例 4:
     * 输入: "LVIII"
     * 输出: 58
     * 解释: L = 50, V= 5, III = 3.
     * 示例 5:
     * 输入: "MCMXCIV"
     * 输出: 1994
     * 解释: M = 1000, CM = 900, XC = 90, IV = 4.
     * 
     * 来源：力扣（LeetCode）
     * 链接：https://leetcode-cn.com/problems/roman-to-integer
     * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
     */
    /**
     * 执行用时 :43 ms, 在所有 Java 提交中击败了9.93%的用户
     * 内存消耗 :47.7 MB, 在所有 Java 提交中击败了64.14%的用户
     */
    public int romanToInt2(String s) {
        Map<String, Integer> m = new HashMap<>();
        m.put("I", 1);
        m.put("V", 5);
        m.put("X", 10);
        m.put("L", 50);
        m.put("C", 100);
        m.put("D", 500);
        m.put("M", 1000);
        int result = 0;
        int prewValue = 0;
        for(int i = 0; i < s.length(); i++) {
            int nowNumber = m.get(s.charAt(i) + "");
            if( i != s.length() - 1) {
                int nextNumber = m.get(s.charAt(i + 1) + "");
                if(nextNumber / nowNumber == 5 || nextNumber / nowNumber == 10) {
                    prewValue = nowNumber;
                    continue;
                }
            }
            if(prewValue > 0){
                result += nowNumber - prewValue;
                prewValue = 0; 
            } else {
                result += nowNumber;
            }
        }
        return result;
    }

    /**
     * 执行用时 :18 ms, 在所有 Java 提交中击败了95.45%的用户
     * 内存消耗 :40.4 MB, 在所有 Java 提交中击败了72.17%的用户
     * 将map中的键的数据类型更换一下，就不用类型转换了
     */
    public int romanToInt(String s) {
        Map<Character, Integer> m = new HashMap<>();
        m.put('I', 1);
        m.put('V', 5);
        m.put('X', 10);
        m.put('L', 50);
        m.put('C', 100);
        m.put('D', 500);
        m.put('M', 1000);
        int result = 0;
        char[] cs = s.toCharArray();
        for(int i = 0; i < cs.length; i++) {
            if(i < cs.length - 1 && m.get(cs[i]) < m.get(cs[i+1])) {
                result -= m.get(cs[i]);
            } else {
                result += m.get(cs[i]);
            }
        }
        return result;

    }
    public static void main(String args[]) {
        String s = "MCMXCVI";
        Solution solution = new Solution();
        int result = solution.romanToInt(s);
        System.out.println(result);
    }
}
