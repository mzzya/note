package Sort

import "github.com/hellojqk/Study/Algorithm"

//选择排序
//思想 每次遍历集合 将最小值的与首位互换
//时间复杂度 O(n²)
// func main() {
// 	sort()
// }
func Sort_Selection() {
	var pos = 0

	for i := 0; i < Algorithm.ArrayALength-1; i++ {
		pos = i
		for j := i + 1; j < Algorithm.ArrayALength; j++ {
			//找出未排序中最小值下标
			if Algorithm.ArrayA[j] < Algorithm.ArrayA[pos] {
				pos = j
			}
		}
		//将最小值与本轮查询中的首位互换
		Algorithm.ArrayA[i], Algorithm.ArrayA[pos] = Algorithm.ArrayA[pos], Algorithm.ArrayA[i]
	}
	Algorithm.ShowArrayA(Algorithm.ArrayA)
}
