/**
爱丽丝和鲍勃一起玩游戏，他们轮流行动。爱丽丝先手开局。

最初，黑板上有一个数字 N 。在每个玩家的回合，玩家需要执行以下操作：

选出任一 x，满足 0 < x < N 且 N % x == 0 。
用 N - x 替换黑板上的数字 N 。
如果玩家无法执行这些操作，就会输掉游戏。

只有在爱丽丝在游戏中取得胜利时才返回 True，否则返回 false。假设两个玩家都以最佳状态参与游戏。

 

示例 1：

输入：2
输出：true
解释：爱丽丝选择 1，鲍勃无法进行操作。
示例 2：

输入：3
输出：false
解释：爱丽丝选择 1，鲍勃也选择 1，然后爱丽丝无法进行操作。
 

提示：

1 <= N <= 1000
*/

package main

import "fmt"

func main(){
    fmt.Println(divisorGame(5))
}

func divisorGame(N int) bool {
    m:=make(map[int]bool)
    return xx(N,m)
}

func xx(n int,m map[int]bool)bool{
    //fmt.Println(n,m)
    if  n == 1 {
        m[1] = false
        return m[2]
    }
    if n == 2 {
        m[2] = true
        return m[2]
    }
    for i:=1;i<n;i++{
        if n%i != 0 {
            continue
        }
        //fmt.Println(n,"******",i,"*******",m)
        var res bool
        if v,ok := m[i];ok {
            res = v 
        }else{
            res := !xx(n-i, m)
            m[i] = res
        }
        if res {
            return true
        }
    }
    return false
}
