import java.util.*;
class Solution {
    /**
     * 判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
     * 示例 1:
     * 输入: 121
     * 输出: true
     * 示例 2:
     * 输入: -121
     * 输出: false
     * 解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
     * 示例 3:
     * 输入: 10
     * 输出: false
     * 解释: 从右向左读, 为 01 。因此它不是一个回文数。
     * 进阶:
     * 你能不将整数转为字符串来解决这个问题吗？
     * 来源：力扣（LeetCode）
     * 链接：https://leetcode-cn.com/problems/palindrome-number
     * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
     */
    /**
     * 用了反转数据的逻辑
     * 执行用时 :48 ms, 在所有 Java 提交中击败了87.62%的用户
     * 内存消耗 :37.7 MB, 在所有 Java 提交中击败了86.53%的用户
     */
     public boolean isPalindrome2(int x) {
        int result = 0;
        int originData = x;
        while(x > 0 && originData % 10 != 0){
            int num = x % 10;
            if(result > Integer.MAX_VALUE / 10 || (result == Integer.MAX_VALUE && num > 7) || result < Integer.MIN_VALUE / 10 || (result == Integer.MIN_VALUE && num < -8)) {
                return false;
            }
            result = result * 10 + num;
            x /= 10;
        }
        return result == originData;
     }

     /**
      * 网上得到的思路：只取一半的值进行比较
      * 执行用时 :49 ms, 在所有 Java 提交中击败了81.44%的用户
      * 内存消耗 :37.8 MB, 在所有 Java 提交中击败了86.13%的用户
      */
     public boolean isPalindrome(int x) {
        if(x < 0 || (x % 10 == 0 && x != 0)) {
            return false;
        }
        int result = 0;
        while(x > result) {
            result = result * 10 + x % 10;
            x /= 10;
        }
        return x == result || result / 10 == x;
     }
    
     public static void main(String[] args) {
        int x = 10201;
        Solution solution = new Solution();
        boolean result = solution.isPalindrome(x);
        System.out.println(result);
     }
}
