package Sort

import "github.com/smgqk/Study/Algorithm"

// func main() {
// 	Straight(Algorithm.ArrayA)
// }

//直接插入排序
func Straight(array []int) {
	var arrayLen = len(array)
	for i := 1; i < arrayLen; i++ {
		//如果前一个值比后一个值大
		//发现了一个介于排序好的数组的值
		if array[i] < array[i-1] {
			var temp = array[i]
			var j = 0
			//循环 将大的值往后移动
			for j = i - 1; j >= 0 && temp < array[j]; j-- {
				array[j+1] = array[j]
			}
			//跳出循环的位置的数比 temp 小 则空位为j+1
			array[j+1] = temp
		}
	}
	Algorithm.ShowArrayA(array)
}
