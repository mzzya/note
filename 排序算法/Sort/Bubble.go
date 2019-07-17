package Sort

import (
	"github.com/smgqk/Study/Algorithm"
)

var Count int

//冒泡排序算法
//遍历集合 如果后面一个数比前面的大 则交换两者位置
//时间复杂度 O(n²)
// func main() {
// 	sort()
// }
func Sort_Bubble() {
	//前一个数跟后一个数比 大的放后边
	//外层循环代表需要执行多少轮比较 最少1轮 最多len-1轮
	//var startTime = time.Now().UnixNano()
	for i := 0; i < Algorithm.ArrayALength-1; i++ {
		//内层为比较，比较相邻两个数，将大的放后边，所以每循环完成 1 轮即可确定 1 个数字
		// 1轮   比较前 i=0 确定数字个数 0 个 终点 <len-1       最后为ary[len-2]与ary[len-1]比较
		// 2轮   比较前 i=1 确定数字个数 1 个 终点 <len-1-1     最后为ary[len-2-1]与ary[len-1-1]比较
		// 3轮   比较前 i=2 确定数字个数 2 个 终点 <len-1-2     最后为ary[len-2-2]与ary[len-1-2]比较
		// ...
		// i+1轮 比较前 i=i 确定数字个数 i 个 终点 <len-1-i     最后为ary[len-2-i]与ary[len-1-i] 比较

		//该轮比较是否发生了交换
		swapped := false
		for j := 0; j < Algorithm.ArrayALength-1-i; j++ {
			if Algorithm.ArrayA[j] > Algorithm.ArrayA[j+1] {
				Algorithm.ArrayA[j], Algorithm.ArrayA[j+1] = Algorithm.ArrayA[j+1], Algorithm.ArrayA[j]
				swapped = true
			}
			Count++
		}
		//如果没有发生交换证明排序已完成
		if !swapped {
			break
		}
	}
	//var endTime = time.Now().UnixNano()
	// log.Printf("用时%20f", float64(endTime-startTime)/1000000)
	// //Algorithm.ShowArrayA(Algorithm.ArrayA)
	// log.Printf("循环次数%d\n", Count)
}
