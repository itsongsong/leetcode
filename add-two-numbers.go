/**
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main
import "fmt"
func main(){
    var l1 = new(ListNode)
    l1.Val = 9
    
    var l1_1 = new(ListNode)
    l1_1.Val = 9

    var l1_2 = new(ListNode)
    l1_2.Val = 9
    l1_1.Next = l1_2
    l1.Next = l1_1

    var l2 = new(ListNode)
    l2.Val = 5

//    var l2_1 = new(ListNode)
 //   l2_1.Val = 6

 //   var l2_2 = new(ListNode)
 //   l2_2.Val = 4
  //  l2_1.Next = l2_2
  //  l2.Next = l2_1

    res := addTwoNumbers(l1, l2)
    for res!=nil {
        fmt.Println(res.Val)
        res=res.Next
    }
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var xx = new(ListNode)
    res :=xx
    var sum,pre = 0,0
    for l1 != nil || l2 != nil {
        var v1 int
        var v2 int
        if l1 != nil {
            v1 = l1.Val
        }
        if l2 != nil{
            v2 = l2.Val
        }
        sum = (v1 + v2 + pre) % 10
        pre = (v1 + v2+pre) / 10
        res.Next = &ListNode{Val:sum%10,Next:nil}
        res = res.Next
        
        if l1 != nil {
            l1 = l1.Next
        }
        if l2 != nil{
            l2 = l2.Next
        }
        if l1 == nil && l2 == nil && pre == 1 {
            res.Next = &ListNode{Val:pre,Next:nil}
            res = res.Next
        }
    }
    return xx.Next
}

type ListNode struct {
    Val int
    Next *ListNode
}
