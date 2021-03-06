import java.util.*;
class Solution {
    /**
     * 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
     * 示例 1:
     * 输入: 123
     * 输出: 321
     * 示例 2:
     * 输入: -123
     * 输出: -321
     * 示例 3:
     * 输入: 120
     * 输出: 21
     * 注意:
     * 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
     * 来源：力扣（LeetCode）
     * 链接：https://leetcode-cn.com/problems/reverse-integer
     * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
     */
    /**
     * 执行用时 :2 ms, 在所有 Java 提交中击败了99.61%的用户
     * 内存消耗 :34.8 MB, 在所有 Java 提交中击败了77.78%的用户
     */
    public int reverse(int x) {
        int result = 0;
        while(x != 0){
            int num = x % 10;
            if(result > Integer.MAX_VALUE / 10 || (result == Integer.MAX_VALUE && num > 7) || result < Integer.MIN_VALUE / 10 || (result == Integer.MIN_VALUE && num < -8)) {
                return 0;
            }
            result = result * 10 + num;
            x /= 10;
        }
        return result;
    }
    
     public static void main(String[] args) {
        int x = -120;
        Solution solution = new Solution();
        int result = solution.reverse(x);
        System.out.println(result);
     }
}
