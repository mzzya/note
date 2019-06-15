package Sort

import (
	"log"

	"github.com/smgqk/Study/Algorithm"
)

// //堆排序
// func main() {
// 	simpleSort(Algorithm.ArrayA)
// }

//这个方法的作用主要是排序 三个一组 的树形位置
//parent 父节点位置
//length 数组长度
func HeapAdjust(array []int, parent int, length int) {
	//代表当前节点父级
	var temp = array[parent]

	//左孩子位置
	var child = 2*parent + 1

	// if child+1 < length {
	// 	log.Printf("父[%d]=%d,左孩子[%d]=%d,右孩子[%d]=%d,parent=%d,length=%d,array=%+v\n",
	// 		parent, temp, child, array[child], child+1, array[child+1], parent, length, array)
	// } else {
	// 	log.Printf("父[%d]=%d,左孩子[%d]=%d,右孩子[%d]=%d,parent=%d,length=%d,array=%+v\n",
	// 		parent, temp, child, array[child], "-1", "-1", parent, length, array)
	// }

	for child < length {

		//这句话就是判断左孩子大还是右孩子大
		// 然后用大的跟 上级父节点比
		// 如果最大的孩子 都小于父节点 证明这组（三个值）已经排序好了(排成大根堆)
		if child+1 < length && array[child] < array[child+1] {
			child++
		}

		//孩子小于父节点，证明此节点已经排好 就可推出这三个的循环
		if temp >= array[child] {
			break
		}

		//父节点 小于 最大孩子节点
		//将最大孩子节点交给父节点
		array[parent] = array[child]
		//复制后需要判断当前孩子节点的值是否还满足 大根堆
		parent = child
		child = 2*parent + 1
	}
	//此时array[parent]即为交换后空余的子节点
	//需要将之前的父节点给它
	//如果没有交换 则原地复制 此处可以判断一下 然后就不需要复制了 if array[parent]=temp {return}
	array[parent] = temp
}

func simpleSort(array []int) {
	newAry := make([]int, len(array))
	//1.构建大顶堆
	for i := len(array)/2 - 1; i >= 0; i-- {
		HeapAdjust(array, i, len(array))
	}

	//log.Println(array)
	//倒着循环 将堆顶值与最末端的值交换 并重新调整堆
	for i := len(array) - 1; i >= 0; i-- {
		array[0], array[i] = array[i], array[0]
		HeapAdjust(array, 0, i)
		//此时最后一个值即为当前存在的最大值
		newAry[i] = array[i]
	}
	Algorithm.ShowArrayA(newAry)
}

//不完善的输出。。
func show(array []int) {
	var parent = len(array)/2 - 1
	//循环深度
	for i := 0; i <= parent; i++ {
		//开始位置为2*parent 结束位置为2*(parent+1)-1 因为下层深度开始为2*(parent+1)
		var startIndex = 2<<uint(i-1) - 1
		var endIndex = 2<<uint(i) - 2
		if i == 0 {
			endIndex = 0
			startIndex = endIndex
		}
		//log.Printf("第{%d}层,startIndex=%d,endIndex=%d\n", i, startIndex, endIndex)
		for j := startIndex; j <= endIndex && j < len(array); j++ {
			log.Printf("%d ", array[j])
		}
		//log.Print("\n")
	}
}
