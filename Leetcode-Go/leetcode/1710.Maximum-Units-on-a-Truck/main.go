package main

import (
	"fmt"
	"sort"
)

func main() {
	box := [][]int{{1, 3}, {2, 2}, {3, 1}}
	res := maximumUnits(box, 4)
	fmt.Println("ans is ", res)
}

func maximumUnits(boxTypes [][]int, truckSize int) (res int) {
    sort.Slice(boxTypes,func(i,j int) bool{return boxTypes[i][1]>boxTypes[j][1]})
    // range BoxTypes中的所有子数组，所以自动遍历b[i][...]
    // 只需关注第二位的index
        for _,p :=range boxTypes{
            // 刚好满足或大于所剩存箱容量
            if p[0]>= truckSize{
               res+= truckSize * p[1]
               break
            }
            // 减去当前箱子
            truckSize-=p[0]
            // p[0]当前箱子 * p[1]箱内单元
            res+=p[0]*p[1]
        }
    return res
}
