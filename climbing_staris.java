import java.util.*;
class Solution {
    /**
     * 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
     * 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
     * 注意：给定 n 是一个正整数。
     * 示例 1：
     * 输入： 2
     * 输出： 2
     * 解释： 有两种方法可以爬到楼顶。
     * 1.  1 阶 + 1 阶
     * 2.  2 阶
     * 示例 2：
     * 输入： 3
     * 输出： 3
     * 解释： 有三种方法可以爬到楼顶。
     * 1.  1 阶 + 1 阶 + 1 阶
     * 2.  1 阶 + 2 阶
     * 3.  2 阶 + 1 阶
     * 来源：力扣（LeetCode）
     * 链接：https://leetcode-cn.com/problems/climbing-stairs
     * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
    */
    /**
     * 执行用时 :0 ms, 在所有 Java 提交中击败了100.00%的用户
     * 内存消耗 :33.7 MB, 在所有 Java 提交中击败了53.51%的用户
     */
    public int climbStairs(int n) {
        if(n == 1)return 1;
        int[] result = new int[n+1];
        result[1] = 1;
        result[2] = 2;
        for(int i = 3; i <= n; i++) {
            result[i] = result[i-1] + result[i-2];
        }
        return result[n];
    }

    public static void main(String args[]) {
        int n = 44;
        Solution solution = new Solution();
        int result = solution.climbStairs(n);
        System.out.println(result); 
    }
}
