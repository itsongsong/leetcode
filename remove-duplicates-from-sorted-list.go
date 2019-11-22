/**
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

示例 1:

输入: 1->1->2
输出: 1->2
示例 2:

输入: 1->1->2->3->3
输出: 1->2->3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list
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
    head.Next.Next.Val = 2
    fmt.Println(deleteDuplicates(head))
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
解题思路：
1. 判断当前值和下一个节点的值是否相等
2. 如果相等，则当前节点的下一个节点为下下个节点
3. 如果不相等，则直接到下一个节点
*/
func deleteDuplicates(head *ListNode) *ListNode {
    x := head
    for x != nil && x.Next != nil{
        if x.Val == x.Next.Val {
            x.Next = x.Next.Next
        }else{
            x = x.Next
        }
    }
    return head
}

type ListNode struct {
    Val int
    Next *ListNode
}
