class Solution {
    /**
     * 报数序列是一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：
     * 1.     1
     * 2.     11
     * 3.     21
     * 4.     1211
     * 5.     111221
     * 1 被读作  "one 1"  ("一个一") , 即 11。
     * 11 被读作 "two 1s" ("两个一"）, 即 21。
     * 21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。
     * 给定一个正整数 n（1 ≤ n ≤ 30），输出报数序列的第 n 项。
     * 注意：整数顺序将表示为一个字符串。
     * 示例 1:
     * 输入: 1
     * 输出: "1"
     * 示例 2:
     * 输入: 4
     * 输出: "1211"
     * 来源：力扣（LeetCode）
     * 链接：https://leetcode-cn.com/problems/count-and-say
     * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
     */
    /**
     * 执行用时 :39 ms, 在所有 Java 提交中击败了14.58%的用户
     * 内存消耗 :63.9 MB, 在所有 Java 提交中击败了19.59%的用户
    */
    public String countAndSay2(int n) {
       String s = "";
       if(n < 1 || n > 30) {
            return "";
       }
       for(int i = 1; i <= n; i++) {
            if(i != 1) {
                char[] cs = s.toCharArray();
                String tempS = "";
                int count = 1;
                for(int j = 0; j < cs.length; j++) {
                   if(j < cs.length - 1 && cs[j] == cs[j+1]){
                        count++;
                        continue;
                   } 
                   tempS += count + "" + cs[j];
                   count = 1;
                }
                s = tempS;
            } else {
                s = "1";
            }
       }
       return s; 
    }

    // 用递归的方法的来解决 53ms ...
    public String countAndSay(int n) {
        if(n == 1) {
            return "1";
        } else {
            return getNext(countAndSay(n -1));
        }
    }
    private String getNext(String s) {
        char[] cs = s.toCharArray();
        StringBuilder sb = new StringBuilder();
        int i = 0;
        int count = 1;
        while(i < cs.length) {
            while(i < cs.length - 1 && cs[i] == cs[i+1]) {
                count++;
                i++;
            }
            System.out.println(count);
            sb.append(count);
            sb.append(cs[i]);
            count = 1;
            i++;
        } 
        return sb.toString();
    }

    public static void main(String args[]) {
        int n = 5;
        Solution solution = new Solution();
        String result = solution.countAndSay(n);
        System.out.println(result);
    }
}
