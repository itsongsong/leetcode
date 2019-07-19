import java.util.*;
class Solution {
    // https://leetcode-cn.com/problems/two-sum/
    /**
     * 执行用时 :4 ms, 在所有 Java 提交中击败了96.20%的用户
     * 内存消耗 :37.1 MB, 在所有 Java 提交中击败了92.07%的用户
     */
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i <= nums.length - 1; i++) {
            map.put(nums[i], i); 
        }     
        for (int i = 0; i <= nums.length - 1; i++) {
            Integer num = map.get(target - nums[i]);
            if(null != num && num != i) {
                int[] result = {i, num};
                return result;
            }
        }        
        throw new IllegalArgumentException("No two sum solution");
    }
    
    /**
     * 此方法执行时间是9ms，占用内存39.1M
     * 超越了69.5%
     */
    public int[] twoSum2(int[] nums, int target) {
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i <= nums.length - 1; i++) {
            map.put(nums[i], i);
        }
        for (int i = 0; i <= nums.length - 1; i++) {
            if(map.containsKey(target - nums[i]) && map.get(target - nums[i]) != i) {
                int[] result = {i, map.get(target - nums[i])};
                return result;
            }
        }
        throw new IllegalArgumentException("No two sum solution");


    }
    
     public static void main(String[] args) {
        int[] nums = {1,2,3,4,5};
        int target = 6;
        Solution solution = new Solution();
        int[] result = solution.twoSum(nums, target);
        System.out.println(Arrays.toString(result));
     }
}
