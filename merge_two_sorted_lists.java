/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    /**
     * 将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 
     * 示例：
     * 输入：1->2->4, 1->3->4
     * 输出：1->1->2->3->4->4
     * 来源：力扣（LeetCode）
     * 链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
     * 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
    */
    /**
     * 执行用时 :2 ms, 在所有 Java 提交中击败了89.91%的用户
     * 内存消耗 :35.9 MB, 在所有 Java 提交中击败了87.54%的用户
    */
    public ListNode mergeTwoLists(ListNode l1, ListNode l2) {
        ListNode header = new ListNode(0);
        merge(header, l1, l2);
        return header.next;
    }
    private void merge(ListNode header, ListNode l1, ListNode l2) {
        if(l1 != null && l2 != null){
            if(l1.val > l2.val) {
                header.next = l2;
                merge(header.next, l1, l2.next);
            }else{
                header.next = l1;
                merge(header.next, l1.next, l2);
            }
        }
        if(l1 == null && l2 == null) {
            header.next = null;
        }

        if(l1 == null && l2 != null) {
            header.next = l2;       
        }
        if(l1 != null && l2 == null) {
            header.next = l1;
        }
    }

    public static void main(String args[]) {
        ListNode l1 = new ListNode(1);
        l1.next = new ListNode(2);
        l1.next.next = new ListNode(4);
        ListNode l2 = new ListNode(1);
        l2.next = new ListNode(3);
        l2.next.next = new ListNode(4);

        Solution solution = new Solution();
        ListNode result = solution.mergeTwoLists(l1, l2);
        while(result != null) {
            System.out.println(result.val);
            result = result.next;
        }
    }
}
class ListNode{ //类
    int val;
    ListNode next; //下一个节点
    ListNode(int x){ //构造函数的初始化
        val = x;
        next = null;
    }
}
