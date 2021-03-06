/**
8 间牢房排成一排，每间牢房不是有人住就是空着。

每天，无论牢房是被占用或空置，都会根据以下规则进行更改：

如果一间牢房的两个相邻的房间都被占用或都是空的，那么该牢房就会被占用。
否则，它就会被空置。
（请注意，由于监狱中的牢房排成一行，所以行中的第一个和最后一个房间无法有两个相邻的房间。）

我们用以下方式描述监狱的当前状态：如果第 i 间牢房被占用，则 cell[i]==1，否则 cell[i]==0。

根据监狱的初始状态，在 N 天后返回监狱的状况（和上述 N 种变化）。

示例 1：

输入：cells = [0,1,0,1,1,0,0,1], N = 7
输出：[0,0,1,1,0,0,0,0]
解释：
下表概述了监狱每天的状况：
Day 0: [0, 1, 0, 1, 1, 0, 0, 1]
Day 1: [0, 1, 1, 0, 0, 0, 0, 0]
Day 2: [0, 0, 0, 0, 1, 1, 1, 0]
Day 3: [0, 1, 1, 0, 0, 1, 0, 0]
Day 4: [0, 0, 0, 0, 0, 1, 0, 0]
Day 5: [0, 1, 1, 1, 0, 1, 0, 0]
Day 6: [0, 0, 1, 0, 1, 1, 0, 0]
Day 7: [0, 0, 1, 1, 0, 0, 0, 0]

示例 2：

输入：cells = [1,0,0,1,0,0,1,0], N = 1000000000
输出：[0,0,1,1,1,1,1,0]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/prison-cells-after-n-days
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main
import "fmt"
func main(){
    //fmt.Println(prisonAfterNDays([]int{0,1,0,1,1,0,0,1}, 1000000000))
    //fmt.Println("-------------------")
    fmt.Println(prisonAfterNDays([]int{0,1,0,1,1,0,0,1}, 64))
}
func prisonAfterNDays(cells []int, N int) []int {
    if N % 14 == 0 {
        N = 14
    }else{
        N = N%14
    }
    for i:=0; i<N; i++{
        prev := cells[0]
        for j:=1; j<len(cells)-1; j++{
            if prev == cells[j+1] {
                prev = cells[j]
                cells[j] = 1
            }else{
                prev = cells[j]
                cells[j] = 0
            }
        }
        cells[0] = 0
        cells[len(cells)-1] = 0
    }
    return cells
}
