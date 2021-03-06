/**
206. 反转链表

反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main
import (
    "fmt"
)
func main(){
    head := new(ListNode)
    head.Val = 1
    head.Next = new(ListNode)
    head.Next.Val = 2
    head.Next.Next = new(ListNode)
    head.Next.Next.Val = 3
    fmt.Println(reverseList(head))
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    x:=new(ListNode)
    for head != nil {
        temp := new(ListNode)
        temp.Val = head.Val
        temp.Next = x.Next
        x.Next = temp
        head=head.Next
    }
    return x.Next
}
type ListNode struct {
    Val int
    Next *ListNode
}
