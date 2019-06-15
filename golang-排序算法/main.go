package main

import (
	"fmt"
	"math/rand"
	"time"
)

//ArrayLen 测试数组长度
var ArrayLen = 10

//Array 测试数组
var Array []int

const IntMapX = 10

const IntMapY = 10

//IntMap 二位数组
var IntMap [IntMapX][IntMapY]int

func main() {
	fmt.Println("####################")
	var result = merge(Array)
	showArray2("结果", result)

	validateArray(result)
}

//merge 归并排序
func merge(array []int) []int {
	var aryLen = len(array)
	if aryLen <= 1 {
		return array
	}
	//首先拆成最小交换单元数据
	var partCount = aryLen / 2
	left := merge(array[0:partCount])
	right := merge(array[partCount:])
	return mergeCore(left, right)
}

//mergeCore 归并排序合并方法
func mergeCore(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	//循环结束时 必定有一个集合为空 另一个集合有1+以上值最坏全部值，
	//切改值必定大于已排序完成的集合
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return
}

//Select 选择排序
func Select(array []int) {
	for i := 0; i < len(array)-1; i++ {
		//假设下标未0位最小项
		var minIndex = i
		//排除已排序好的位置
		for j := i + 1; j < len(array); j++ {
			//在剩余数字中寻找最小值 即为已排序完成的最大值
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}
		array[i], array[minIndex] = array[minIndex], array[i]
	}
}

//Insert 插入排序
func Insert(array []int) {
	//234 75788
	for i := 1; i < len(array); i++ {
		//如果当前的数大于已排序好的最后一个数则跳过
		if array[i] > array[i-1] {
			continue
		}
		//简版判断 当前值比前一个值小则交换两者数据
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			array[j-1], array[j] = array[j], array[j-1]
		}

		// //保存当前值
		// var temp = array[i]
		// var j int
		// //从已经排好的队尾取值与当前数字比较，如果取出的值大于当前值，则将队尾值向后移动
		// for j = i - 1; j >= 0 && array[j] > temp; j-- {
		// 	array[j+1] = array[j]
		// }
		// //最坏j=-1,排好的队列全部右移
		// array[j+1] = temp
	}
}

//Bubble 冒泡排序
func Bubble(array []int) {
	for i := 0; i < len(array)-1; i++ {
		//内循环从下一个数开始，如果比第一个数小则交换位置
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[i] {
				array[i], array[j] = array[j], array[i]
			}
		}
		//这样一圈下来，最小的数一定在最左侧
	}
}

//Quick 快速排序算法
func Quick(array []int) {
	if len(array) <= 1 {
		return
	}
	//将第一个值作为比较项目
	key := array[0]
	//定义左侧和右侧边间
	i, j := 0, len(array)-1
	//当循环内还存在数据时
	for i < j {
		//从右侧开始找大于key的值，存在则从右侧缩小范围
		for i < j && array[j] >= key {
			j--
		}
		//从右侧找到小于key的值时，将值给换到最左侧
		array[i] = array[j]
		showArray("右侧交换")
		//从左侧开始找小于key的值，存在则从左侧缩小范围
		for i < j && array[i] <= key {
			i++
		}
		array[j] = array[i]
		showArray("左侧交换")
	}
	//循环退出时必然存在i=j 所以下方 i 或 j 无所谓
	array[i] = key
	// showArray("一圈")
	Quick(array[:i])
	Quick(array[i+1:])
}

func init() {
	initAry()
	// initMap()
}
func initMap() {
	for i := 0; i < IntMapX; i++ {
		var ary [IntMapY]int
		for j := 0; j < IntMapY; j++ {
			ary[j] = j
		}
		IntMap[i] = ary
	}
	showMap()
}
func showMap() {
	for i := 0; i < len(IntMap); i++ {
		for j := 0; j < len(IntMap[i]); j++ {
			fmt.Printf("[%d,%d]  ", j, len(IntMap)-1-i)
		}
		fmt.Printf("\n\n\n")
	}
}

//Init 初始化
func initAry() {
	Array = make([]int, ArrayLen)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < ArrayLen; i++ {
		Array[i] = r.Intn(ArrayLen * 10)
	}
	showArray("初始")
}

//showArray 展示数组
func showArray(label string) {
	showArray2(label, Array)
}

//showArray 展示数组
func showArray2(label string, arr []int) {
	fmt.Printf("%s\tArrayLen:%d\tArray:%#v\n", label, len(arr), arr)
}

//validateArray 验证数组
func validateArray(array []int) {
	result := true
	for i := 0; i < len(array)-1; i++ {
		if array[i] > array[i+1] {
			result = false
			break
		}
	}
	fmt.Printf("数组排序验证结果:%t\n", result)
}
